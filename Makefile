.PHONY: default
# export GOPATH:=$(shell pwd)

ifndef $(GOAPTH)
GOPATH = $(HOME)/go
endif

go-bindata = $(GOPATH)/bin
BUILD_ENV := CGO_ENABLED=0

$(go-bindata)/go-bindata:
	go get -d github.com/jteeuwen/go-bindata/...

default: client server

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

all: build-windows-amd64 build-windows-386 build-linux-amd64 build-linux-386 build-darwin-amd64 build-darwin-386

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

ifneq ($(MAKECMDGOALS),)
	echo 'sdf'
endif

clean:
	rm $(go-bindata)/go-bindata ngrok ngrokd server/assets client/assets -rf