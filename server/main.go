package server

import (
	"crypto/tls"
	"errors"
	"math/rand"
	"ngrok/conn"
	"ngrok/log"
	"ngrok/msg"
	"ngrok/util"
	"os"
	"runtime/debug"
	"time"

	"github.com/go-redis/redis"
)

const (
	registryCacheSize   uint64        = 1024 * 1024 // 1 MB
	connReadTimeout                   = 12 * time.Second
	redisUserList       string        = "Ngrok-User-List"
	redisAuthHeader     string        = "Admin-Authorization"
	redisAccessToken    string        = "Admin-Access-token"
	redisRefreshToken   string        = "Admin-refresh-token"
	redisAccessExpired  time.Duration = 2 * time.Hour
	redisReflashExpired time.Duration = 48 * time.Hour
)

// GLOBALS
var (
	tunnelRegistry  *TunnelRegistry
	controlRegistry *ControlRegistry

	// XXX: kill these global variables - they're only used in tunnel.go for constructing forwarding URLs
	opts      *Options
	listeners map[string]*conn.Listener
	client    *redis.Client
)

func NewProxy(pxyConn conn.Conn, regPxy *msg.RegProxy) {
	// fail gracefully if the proxy connection fails to register
	defer func() {
		if r := recover(); r != nil {
			pxyConn.Warn("Failed with error: %v", r)
			pxyConn.Close()
		}
	}()

	// set logging prefix
	pxyConn.SetType("pxy")

	// look up the control connection for this proxy
	pxyConn.Info("Registering new proxy for %s", regPxy.ClientId)
	ctl := controlRegistry.Get(regPxy.ClientId)

	if ctl == nil {
		panic("No client found for identifier: " + regPxy.ClientId)
	}

	ctl.RegisterProxy(pxyConn)
}

func CheckAuth(auth string) (string, error) {
	if client == nil {
		return "", nil
	}
	res := client.HMGet(redisUserList, auth).Val()
	if len(res) != 1 {
		return "", errors.New("auth token string error")
	}
	v := res[0]
	log.Info("%s - %s", auth, v)
	if v == nil {
		return "", errors.New("auth token string error nil")
	}
	return v.(string), nil
}

// Listen for incoming control and proxy connections
// We listen for incoming control and proxy connections on the same port
// for ease of deployment. The hope is that by running on port 443, using
// TLS and running all connections over the same port, we can bust through
// restrictive firewalls.
func tunnelListener(addr string, tlsConfig *tls.Config) {
	// listen for incoming connections
	listener, err := conn.Listen(addr, "tun", tlsConfig)
	if err != nil {
		panic(err)
	}

	log.Info("Listening for control and proxy connections on %s", listener.Addr.String())
	for c := range listener.Conns {
		go func(tunnelConn conn.Conn) {
			// don't crash on panics
			defer func() {
				if r := recover(); r != nil {
					tunnelConn.Info("tunnelListener failed with error %v: %s", r, debug.Stack())
				}
			}()

			tunnelConn.SetReadDeadline(time.Now().Add(connReadTimeout))
			var rawMsg msg.Message
			if rawMsg, err = msg.ReadMsg(tunnelConn); err != nil {
				tunnelConn.Warn("Failed to read message: %v", err)
				tunnelConn.Close()
				return
			}
			// don't timeout after the initial read, tunnel heartbeating will kill
			// dead connections
			tunnelConn.SetReadDeadline(time.Time{})

			switch m := rawMsg.(type) {
			case *msg.Auth:
				NewControl(tunnelConn, m)

			case *msg.RegProxy:
				NewProxy(tunnelConn, m)

			default:
				tunnelConn.Close()
			}
		}(c)
	}
}

func Main() {
	// parse options
	opts = parseArgs()

	// init logging
	log.LogTo(opts.logto, opts.loglevel)

	// seed random number generator
	seed, err := util.RandomSeed()
	if err != nil {
		panic(err)
	}
	rand.Seed(seed)

	// init tunnel/control registry
	registryCacheFile := os.Getenv("REGISTRY_CACHE_FILE")
	tunnelRegistry = NewTunnelRegistry(registryCacheSize, registryCacheFile)
	controlRegistry = NewControlRegistry()

	// start listeners
	listeners = make(map[string]*conn.Listener)

	// load tls configuration
	tlsConfig, err := LoadTLSConfig(opts.tlsCrt, opts.tlsKey)
	if err != nil {
		panic(err)
	}

	// listen for http
	if opts.httpAddr != "" {
		listeners["http"] = startHttpListener(opts.httpAddr, nil)
	}

	// listen for https
	if opts.httpsAddr != "" {
		listeners["https"] = startHttpListener(opts.httpsAddr, tlsConfig)
	}

	if opts.redisAddr != "" {
		client = redis.NewClient(&redis.Options{Addr: opts.redisAddr, Password: opts.redisPwd})
		log.Info("Add redis %s token type, abandon proxy", opts.redisAddr)
		go start()
	}

	// ngrok clients
	tunnelListener(opts.tunnelAddr, tlsConfig)
}
