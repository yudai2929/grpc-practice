generate:
	cd api && protoc --go_out=../pkg/grpc --go_opt=paths=source_relative \
	--go-grpc_out=../pkg/grpc --go-grpc_opt=paths=source_relative \
	hello.proto

run-server:
	go run cmd/grpc_gateway/server/main.go


run-client:
	go run cmd/grpc_gateway/client/main.go

run-proxy:
	go run cmd/grpc_gateway/proxy/main.go

gen-stub:
	buf generate

PHONY: generate run-server run-client
