package usecase

type TodoUsecase interface {
	CreateTodo(CreateTodoInput) (*CreateTodoOutput, error)
	GetTodoByID(GetTodoByIDInput) (*GetTodoByIDOutput, error)
}
