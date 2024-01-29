package mock

import (
	"fmt"

	"github.com/yudai2929/grpc-practice/pkg/user/domain/entity"
	"github.com/yudai2929/grpc-practice/pkg/user/domain/repository"
)

type userRepositoryMock struct {
	db []*entity.User
}

func NewUserRepositoryMock() repository.UserRepository {
	return &userRepositoryMock{}
}

func (ur *userRepositoryMock) CreateUser(user *entity.User) (*entity.User, error) {
	ur.db = append(ur.db, user)
	return user, nil
}

func (ur *userRepositoryMock) GetUserByID(id string) (*entity.User, error) {
	for _, user := range ur.db {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, fmt.Errorf("user not found")
}

func (ur *userRepositoryMock) GetUserByEmail(email string) (*entity.User, error) {
	for _, user := range ur.db {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, fmt.Errorf("user not found")
}
