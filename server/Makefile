.PHONY: run start build

run: build start

start:
	go run -race main.go

build:
	go build -race -buildmode=plugin -o plugins/login.so plugin/login/main.go
	go build -race -buildmode=plugin -o plugins/handshake.so plugin/handshake/main.go