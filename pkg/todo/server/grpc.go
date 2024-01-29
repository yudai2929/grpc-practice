package server

import (
	"github.com/yudai2929/grpc-practice/pkg/libs/interceptor"
	"github.com/yudai2929/grpc-practice/pkg/todo/adapter"
	"github.com/yudai2929/grpc-practice/pkg/todo/infrastructure/mock"
	"github.com/yudai2929/grpc-practice/pkg/todo/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func New() *grpc.Server {
	s := grpc.NewServer(grpc.ChainUnaryInterceptor(interceptor.LogInterceptor))

	todoRepository := mock.NewTodoRepositoryMock()

	todoUsecase := usecase.NewTodoUsecase(todoRepository)

	todoServer := adapter.NewTodoServer(todoUsecase)

	todoServer.Register(s)

	reflection.Register(s)

	return s
}
