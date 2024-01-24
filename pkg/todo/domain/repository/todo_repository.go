package repository

import "github.com/yudai2929/grpc-practice/pkg/todo/domain/entity"

type TodoRepository interface {
	CreateTodo(*entity.Todo) (*entity.Todo, error)
	GetTodoByID(string) (*entity.Todo, error)
}
