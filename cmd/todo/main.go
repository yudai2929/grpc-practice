package main

import (
	"fmt"
	"net"

	"github.com/yudai2929/grpc-practice/pkg/todo/server"
)

func main() {
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	server := server.New()

	fmt.Printf("start gRPC server port: %v", port)

	if err := server.Serve(listener); err != nil {
		panic(err)
	}

}
