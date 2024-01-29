package usecase

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/yudai2929/grpc-practice/pkg/user/domain/entity"
	"github.com/yudai2929/grpc-practice/pkg/user/domain/repository"
)

type authUsecase struct {
	validate       *validator.Validate
	userRepository repository.UserRepository
}

func NewAuthUsecase(userRepository repository.UserRepository) AuthUsecase {
	return &authUsecase{
		validate:       validator.New(),
		userRepository: userRepository,
	}
}

func (au *authUsecase) SignUp(input SignUpInput) (*SignUpOutput, error) {

	if err := au.validate.Struct(input); err != nil {
		return nil, err
	}

	user, err := au.userRepository.GetUserByEmail(input.Email)

	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, fmt.Errorf("user already exists")
	}

	newUser := &entity.User{
		ID:        uuid.New().String(),
		Name:      input.Name,
		Email:     input.Email,
		Password:  input.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	newUser, err = au.userRepository.CreateUser(newUser)

	if err != nil {
		return nil, err
	}

	output := &SignUpOutput{
		Token:  newUser.ID,
		UserID: newUser.ID,
	}

	return output, nil
}

func (au *authUsecase) SignIn(input SignInInput) (*SignInOutput, error) {

	if err := au.validate.Struct(input); err != nil {
		return nil, err
	}

	user, err := au.userRepository.GetUserByEmail(input.Email)

	if err != nil {
		return nil, err
	}

	if user.Password != input.Password {
		return nil, err
	}

	output := &SignInOutput{
		Token:  user.ID,
		UserID: user.ID,
	}

	return output, nil
}
