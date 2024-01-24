package usecase

import "github.com/yudai2929/grpc-practice/pkg/todo/domain/entity"

type CreateTodoInput struct {
	Title       string `validate:"required"`
	Description string `validate:"required"`
}

type CreateTodoOutput struct {
	Todo *entity.Todo
}

type GetTodoByIDInput struct {
	ID string `validate:"required"`
}

type GetTodoByIDOutput struct {
	Todo *entity.Todo
}
