package user

import (
	"context"

	"github.com/vnnyx/gas-api/model"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user model.User) (model.User, error)
	FindUserById(ctx context.Context, userID string) (user model.User, err error)
	FindUserByUsername(ctx context.Context, username string) (user model.User, err error)
	FindAllUser(ctx context.Context) (users []model.User, err error)
	UpdateUser(ctx context.Context, user model.User) (model.User, error)
	DeleteUser(ctx context.Context, userID string) error
}
