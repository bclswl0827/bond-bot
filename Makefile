.PHONY: build docker
BINARY=bond-bot
OS_ARCH=`uname -m`

build: armv7l aarch64 x86_64 i386

armv7l:
	@mkdir -p ./release/armv7l
	@rm -rf  ./release/armv7l/*
	@env GOOS=linux GOARCH=arm GOARM=7 go build -ldflags="-s -w" -o ./release/armv7l/${BINARY} ./*.go
	@cp ./config.json ./release/armv7l

aarch64:
	@mkdir -p ./release/aarch64
	@rm -rf  ./release/aarch64/*
	@env GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o ./release/aarch64/${BINARY} ./*.go
	@cp ./config.json ./release/aarch64

x86_64:
	@mkdir -p ./release/x86_64
	@rm -rf  ./release/x86_64/*
	@env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./release/x86_64/${BINARY} ./*.go
	@cp ./config.json ./release/x86_64

i386:
	@mkdir -p ./release/i386
	@rm -rf  ./release/i386/*
	@env GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o ./release/i386/${BINARY} ./*.go
	@cp ./config.json ./release/i386

docker:
	@docker build -t bond-bot:latest --build-arg ARCH=$(OS_ARCH) .
