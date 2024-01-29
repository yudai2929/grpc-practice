package usecase

import "github.com/yudai2929/grpc-practice/pkg/user/domain/entity"

type SignUpInput struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type SignUpOutput struct {
	Token  string
	UserID string
}

type SignInInput struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type SignInOutput struct {
	Token  string
	UserID string
}

type GetUserByIDInput struct {
	ID string `validate:"required"`
}

type GetUserByIDOutput struct {
	User *entity.User
}
