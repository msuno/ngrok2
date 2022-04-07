package util

import "time"

var (
	RedisUserList       string        = "Ngrok-User-List"
	RedisUserSet        string        = "Ngrok-User-Set"
	RedisAuthHeader     string        = "Admin-Authorization"
	RedisAccessToken    string        = "Admin-Access-token"
	RedisRefreshToken   string        = "Admin-refresh-token"
	RedisAccessExpired  time.Duration = 2 * time.Hour
	RedisReflashExpired time.Duration = 48 * time.Hour
)
