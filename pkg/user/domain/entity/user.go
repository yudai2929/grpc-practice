package entity

import (
	"time"

	"github.com/yudai2929/grpc-practice/pkg/libs/password"
)

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) IsEqualPassword(notHashPassword string) bool {
	hash, err := password.Hash(notHashPassword)

	if err != nil {
		return false
	}

	return u.Password == hash
}
