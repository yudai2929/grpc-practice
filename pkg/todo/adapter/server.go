package adapter

import (
	"github.com/yudai2929/grpc-practice/pkg/todo/usecase"
	"google.golang.org/grpc"

	pb "github.com/yudai2929/grpc-practice/proto/gen/go/todo"
)

type TodoServer struct {
	pb.UnimplementedTodoServiceServer
	usecase usecase.TodoUsecase
}

func NewTodoServer(usecase usecase.TodoUsecase) *TodoServer {
	return &TodoServer{
		usecase: usecase,
	}
}

func (s *TodoServer) Register(server *grpc.Server) {
	pb.RegisterTodoServiceServer(server, s)
}
