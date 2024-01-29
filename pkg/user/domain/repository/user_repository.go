package repository

import "github.com/yudai2929/grpc-practice/pkg/user/domain/entity"

type UserRepository interface {
	CreateUser(user *entity.User) (*entity.User, error)
	GetUserByID(id string) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
}
