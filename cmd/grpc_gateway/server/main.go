package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	// _ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	// _ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	"google.golang.org/grpc"
	// _ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	// _ "google.golang.org/protobuf/cmd/protoc-gen-go"

	pb "github.com/yudai2929/grpc-practice/proto/gen/go/gateway"
)

type myServer struct {
	// pb.UnimplementedGateWayServiceServer
}

func newServer() *myServer {
	return &myServer{}
}

func (s *myServer) Echo(ctx context.Context, req *pb.StringMessage) (*pb.StringMessage, error) {
	return &pb.StringMessage{
		Value: fmt.Sprintf("Hello, %s!", req.GetValue()),
	}, nil
}

func main() {
	fmt.Println("start gRPC Server.")

	port := 8080
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("failed to listen.")
		return
	}

	server := grpc.NewServer()
	pb.RegisterGateWayServiceServer(server, newServer())

	go func() {
		log.Printf("start gRPC server port: %v", port)
		server.Serve(lis)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	server.GracefulStop()

}
