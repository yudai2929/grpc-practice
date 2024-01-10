generate:
	cd api && protoc --go_out=../pkg/grpc --go_opt=paths=source_relative \
	--go-grpc_out=../pkg/grpc --go-grpc_opt=paths=source_relative \
	hello.proto

run-server:
	go run cmd/server/main.go


run-client:
	go run cmd/client/main.go


PHONY: generate run-server run-client