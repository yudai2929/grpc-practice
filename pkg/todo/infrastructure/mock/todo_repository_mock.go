package mock

import (
	"fmt"

	"github.com/yudai2929/grpc-practice/pkg/todo/domain/entity"
	"github.com/yudai2929/grpc-practice/pkg/todo/domain/repository"
)

type todoRepositoryMock struct {
	db []*entity.Todo
}

func NewTodoRepositoryMock() repository.TodoRepository {
	return &todoRepositoryMock{}
}

func (tr *todoRepositoryMock) CreateTodo(todo *entity.Todo) (*entity.Todo, error) {
	tr.db = append(tr.db, todo)
	return todo, nil
}

func (tr *todoRepositoryMock) GetTodoByID(id string) (*entity.Todo, error) {
	for _, todo := range tr.db {
		if todo.ID == id {
			return todo, nil
		}
	}

	return nil, fmt.Errorf("todo not found")
}
