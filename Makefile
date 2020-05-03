.PHONY: default server client deps all assets client-assets server-assets clean
# export GOPATH:=$(shell pwd)

ifndef $(GOAPTH)
GOPATH = $(HOME)/go
endif

go-bindata = $(GOPATH)/bin

$(go-bindata)/go-bindata:
	go get -u github.com/jteeuwen/go-bindata/...

default: all

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

all: client server

clean:
	rm $(go-bindata)/go-bindata ngrok ngrokd server/assets client/assets -rf