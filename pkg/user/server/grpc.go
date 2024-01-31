package server

import (
	"github.com/yudai2929/grpc-practice/pkg/libs/interceptor"
	"github.com/yudai2929/grpc-practice/pkg/user/adapter"
	"github.com/yudai2929/grpc-practice/pkg/user/infrastructure/spanner"
	"github.com/yudai2929/grpc-practice/pkg/user/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func New() *grpc.Server {
	s := grpc.NewServer(grpc.ChainUnaryInterceptor(interceptor.LogInterceptor))

	userRepository := spanner.NewUserRepositoryImpl()

	userUsecase := usecase.NewUserUsecase(userRepository)

	authUsecase := usecase.NewAuthUsecase(userRepository)

	userServer := adapter.NewUserServer(userUsecase, authUsecase)

	userServer.Register(s)

	reflection.Register(s)
	return s
}
