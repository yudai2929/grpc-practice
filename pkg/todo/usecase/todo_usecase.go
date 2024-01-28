package usecase

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/yudai2929/grpc-practice/pkg/todo/domain/entity"
	"github.com/yudai2929/grpc-practice/pkg/todo/domain/repository"
)

type todoUsecase struct {
	validate       *validator.Validate
	todoRepository repository.TodoRepository
}

func NewTodoUsecase(todoRepository repository.TodoRepository) TodoUsecase {
	return &todoUsecase{
		validate:       validator.New(),
		todoRepository: todoRepository,
	}
}

func (tu *todoUsecase) CreateTodo(input CreateTodoInput) (*CreateTodoOutput, error) {
	if err := tu.validate.Struct(input); err != nil {
		return nil, err
	}

	todo := &entity.Todo{
		ID:          uuid.New().String(),
		Title:       input.Title,
		Description: input.Description,
		Done:        false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	todo, err := tu.todoRepository.CreateTodo(todo)

	if err != nil {
		return nil, err
	}

	output := &CreateTodoOutput{
		todo,
	}

	return output, nil
}

func (tu *todoUsecase) GetTodoByID(input GetTodoByIDInput) (*GetTodoByIDOutput, error) {
	if err := tu.validate.Struct(input); err != nil {
		return nil, err
	}

	todo, err := tu.todoRepository.GetTodoByID(input.ID)

	if err != nil {
		return nil, err
	}

	output := &GetTodoByIDOutput{
		todo,
	}

	return output, nil
}
