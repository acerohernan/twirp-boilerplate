docker-dev:
	docker compose -f docker-compose.dev.yaml up

dev:
	CompileDaemon --build="go build main.go" --command=./main

gen:
	go generate ./...

wire:
	cd pkg/server && wire

protoc:
	buf generate proto

test:
	go test -v ./...