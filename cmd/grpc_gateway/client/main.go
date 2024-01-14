package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/yudai2929/grpc-practice/gen/go/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	scanner *bufio.Scanner
	client  pb.GateWayServiceClient
)

func main() {
	fmt.Println("start gRPC Client.")

	scanner = bufio.NewScanner(os.Stdin)

	address := "localhost:8080"
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)

	if err != nil {
		log.Fatal("Connection failed.")
		return
	}

	defer conn.Close()

	client := pb.NewGateWayServiceClient(conn)

	fmt.Println("Please enter your name.")

	scanner.Scan()
	name := scanner.Text()

	req := &pb.StringMessage{
		Value: name,
	}

	ctx := context.Background()

	res, err := client.Echo(ctx, req)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetValue())

}
