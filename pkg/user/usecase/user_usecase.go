package usecase

import (
	"github.com/go-playground/validator/v10"
	"github.com/yudai2929/grpc-practice/pkg/user/domain/repository"
)

type userUsecase struct {
	validate       *validator.Validate
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		validate:       validator.New(),
		userRepository: userRepository,
	}
}

func (uu *userUsecase) GetUserByID(input GetUserByIDInput) (*GetUserByIDOutput, error) {
	if err := uu.validate.Struct(input); err != nil {
		return nil, err
	}

	user, err := uu.userRepository.GetUserByID(input.ID)

	if err != nil {
		return nil, err
	}

	output := &GetUserByIDOutput{
		user,
	}

	return output, nil
}
