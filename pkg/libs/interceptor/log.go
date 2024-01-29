package interceptor

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func LogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("start: %v", info.FullMethod)
	log.Printf("request: %v", req)
	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf("error: %v", err)
	}
	log.Printf("response: %v", resp)
	log.Printf("end: %v", info.FullMethod)
	return resp, err
}
