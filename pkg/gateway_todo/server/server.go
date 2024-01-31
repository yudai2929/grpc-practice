package server

import (
	"github.com/yudai2929/grpc-practice/proto/gen/go/gateway_todo"
)

type Server struct {
	todoClient gateway_todo.TodoServiceClient
	userClient gateway_todo.UserServiceClient
}

func New(todoClient gateway_todo.TodoServiceClient, userClient gateway_todo.UserServiceClient) *Server {
	return &Server{
		todoClient: todoClient,
		userClient: userClient,
	}
}
