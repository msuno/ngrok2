.PHONY: default


go-bindata = $(GOPATH)/bin
BUILD_ENV := CGO_ENABLED=0

$(go-bindata)/go-bindata:
	go install -a github.com/jteeuwen/go-bindata/...@latest

default: start client server end

deps: assets
	go mod download

server: deps
	go build ./main/ngrokd

client: deps
	go build ./main/ngrok

assets: client-assets server-assets

client-assets: $(go-bindata)/go-bindata
	$(go-bindata)/go-bindata -pkg=assets -o=client/assets/assets.go assets/client/...

server-assets: $(go-bindata)/go-bindata
	$(go-bindata)/go-bindata -pkg=assets -o=server/assets/assets.go assets/server/...

all: start deps build-windows-amd64 build-windows-386 build-linux-amd64 build-linux-386 build-darwin-amd64 end

build-windows-amd64:
	${BUILD_ENV} GOARCH=amd64 GOOS=windows go build -o windows-amd64/ngrokd.exe ./main/ngrokd
	${BUILD_ENV} GOARCH=amd64 GOOS=windows go build -o windows-amd64/ngrok.exe ./main/ngrok

build-windows-386:
	${BUILD_ENV} GOARCH=386 GOOS=windows go build -o windows-386/ngrok.exe ./main/ngrok
	${BUILD_ENV} GOARCH=386 GOOS=windows go build -o windows-386/ngrokd.exe ./main/ngrokd

build-linux-amd64:
	${BUILD_ENV} GOARCH=amd64 GOOS=linux go build -o linux-amd64/ngrokd ./main/ngrokd
	${BUILD_ENV} GOARCH=amd64 GOOS=linux go build -o linux-amd64/ngrok ./main/ngrok

build-linux-386:
	${BUILD_ENV} GOARCH=386 GOOS=linux go build -o linux-386/ngrok ./main/ngrok
	${BUILD_ENV} GOARCH=386 GOOS=linux go build -o linux-386/ngrokd ./main/ngrokd

build-darwin-amd64:
	${BUILD_ENV} GOARCH=amd64 GOOS=darwin go build -o darwin-amd64/ngrokd ./main/ngrokd
	${BUILD_ENV} GOARCH=amd64 GOOS=darwin go build -o darwin-amd64/ngrok ./main/ngrok

build-darwin-386:
	${BUILD_ENV} GOARCH=386 GOOS=darwin go build -o darwin-386/ngrok ./main/ngrok
	${BUILD_ENV} GOARCH=386 GOOS=darwin go build -o darwin-386/ngrokd ./main/ngrokd

start:
	@echo '开始编译，环境变量GOPATH地址为：'$(GOPATH)

end:
	$(info 编译结束, bin目录$(GOPATH)/bin)

clean:
	rm $(go-bindata)/go-bindata ngrok ngrokd server/assets client/assets linux-* windows-* darwin-* -rf
