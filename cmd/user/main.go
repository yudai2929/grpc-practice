package main

import (
	"fmt"
	"net"

	"github.com/yudai2929/grpc-practice/pkg/user/server"
)

func main() {
	port := 8081
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	server := server.New()

	fmt.Println("start gRPC Server" + fmt.Sprintf(" :%d", port))

	if err := server.Serve(listener); err != nil {
		panic(err)
	}

}
