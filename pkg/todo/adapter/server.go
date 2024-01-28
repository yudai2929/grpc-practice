package adapter

import (
	"context"

	"github.com/yudai2929/grpc-practice/pkg/todo/usecase"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

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

func (s *TodoServer) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	input := &usecase.CreateTodoInput{
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
	}

	output, err := s.usecase.CreateTodo(*input)

	if err != nil {
		return nil, err
	}

	res := &pb.CreateTodoResponse{
		Todo: &pb.Todo{
			Id:          output.Todo.ID,
			Title:       output.Todo.Title,
			Description: output.Todo.Description,
			Done:        output.Todo.Done,
			CreatedAt:   timestamppb.New(output.Todo.CreatedAt),
			UpdatedAt:   timestamppb.New(output.Todo.UpdatedAt),
		},
	}

	return res, nil
}

func (s *TodoServer) GetTodo(ctx context.Context, req *pb.GetTodoRequest) (*pb.GetTodoResponse, error) {
	input := &usecase.GetTodoByIDInput{
		ID: req.GetId(),
	}

	output, err := s.usecase.GetTodoByID(*input)

	if err != nil {
		return nil, err
	}

	res := &pb.GetTodoResponse{
		Todo: &pb.Todo{
			Id:          output.Todo.ID,
			Title:       output.Todo.Title,
			Description: output.Todo.Description,
			Done:        output.Todo.Done,
			CreatedAt:   timestamppb.New(output.Todo.CreatedAt),
			UpdatedAt:   timestamppb.New(output.Todo.UpdatedAt),
		},
	}

	return res, nil
}
