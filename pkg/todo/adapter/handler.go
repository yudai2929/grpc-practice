package adapter

import (
	"context"

	"github.com/yudai2929/grpc-practice/pkg/todo/usecase"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/yudai2929/grpc-practice/proto/gen/go/todo"
)

type TodoHandler struct {
	usecase usecase.TodoUsecase
}

func NewTodoHandler(usecase usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{
		usecase: usecase,
	}
}

func (th *TodoHandler) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	input := &usecase.CreateTodoInput{
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
	}

	output, err := th.usecase.CreateTodo(*input)

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

func (th *TodoHandler) GetTodo(ctx context.Context, req *pb.GetTodoRequest) (*pb.GetTodoResponse, error) {
	input := &usecase.GetTodoByIDInput{
		ID: req.GetId(),
	}

	output, err := th.usecase.GetTodoByID(*input)

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
