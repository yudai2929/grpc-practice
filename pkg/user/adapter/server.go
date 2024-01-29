package adapter

import (
	"github.com/yudai2929/grpc-practice/pkg/user/usecase"
	pb "github.com/yudai2929/grpc-practice/proto/gen/go/user"
	"google.golang.org/grpc"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	userUsecase usecase.UserUsecase
	authUsecase usecase.AuthUsecase
}

func NewUserServer(userUsecase usecase.UserUsecase, authUsecase usecase.AuthUsecase) *UserServer {
	return &UserServer{
		userUsecase: userUsecase,
		authUsecase: authUsecase,
	}
}

func (s *UserServer) Register(server *grpc.Server) {
	pb.RegisterUserServiceServer(server, s)
}
