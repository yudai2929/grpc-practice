package usecase

type UserUsecase interface {
	GetUserByID(GetUserByIDInput) (*GetUserByIDOutput, error)
}
