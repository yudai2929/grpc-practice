package spanner

import (
	"context"

	"github.com/yudai2929/grpc-practice/db/gen/yo"
	libspanner "github.com/yudai2929/grpc-practice/pkg/libs/spanner"
	"github.com/yudai2929/grpc-practice/pkg/user/domain/entity"
	"github.com/yudai2929/grpc-practice/pkg/user/domain/repository"

	"cloud.google.com/go/spanner"
)

type userRepositoryImpl struct{}

func NewUserRepositoryImpl() repository.UserRepository {
	return &userRepositoryImpl{}
}

func (ur *userRepositoryImpl) CreateUser(user *entity.User) (*entity.User, error) {
	client, err := libspanner.NewClient()

	if err != nil {
		return nil, err
	}

	defer client.Close()

	ctx := context.Background()

	f := func(ctx context.Context, tx *spanner.ReadWriteTransaction) error {
		muts := []*spanner.Mutation{
			toYoUser(user).Insert(ctx),
		}
		return tx.BufferWrite(muts)
	}
	_, err = client.ReadWriteTransaction(ctx, f)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepositoryImpl) GetUserByID(id string) (*entity.User, error) {

	client, err := libspanner.NewClient()

	if err != nil {
		return nil, err
	}
	defer client.Close()

	ctx := context.Background()

	tx := client.ReadOnlyTransaction()

	defer tx.Close()

	user, err := yo.FindUser(ctx, tx, id)

	if err != nil {
		return nil, err
	}

	return toEntityUser(user), nil

}

func (ur *userRepositoryImpl) GetUserByEmail(email string) (*entity.User, error) {

	client, err := libspanner.NewClient()

	if err != nil {
		return nil, err
	}

	defer client.Close()

	ctx := context.Background()

	tx := client.ReadOnlyTransaction()

	defer tx.Close()

	users, err := yo.FindUsersByEmail(ctx, tx, email)

	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	user := toEntityUser(users[0])

	return user, nil

}

func toYoUser(user *entity.User) *yo.User {
	return &yo.User{
		UserID:    user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func toEntityUser(user *yo.User) *entity.User {
	return &entity.User{
		ID:        user.UserID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
