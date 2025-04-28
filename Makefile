dev:
	go run cmd/main.go

build:
	bash cmd/build.sh

generate_grpc:
	protoc --go_out=. --go-grpc_out=. api/*/*.proto