package user

import (
	"context"

	"github.com/vnnyx/gas-api/model"
)

type UserService interface {
	CreateUser(ctx context.Context, request model.UserCreateRequest) (response model.UserResponse, err error)
	GetUserByID(ctx context.Context, userID string) (response model.UserResponse, err error)
	GetAllUser(ctx context.Context) (response []model.UserResponse, err error)
	UpdateUserProfile(ctx context.Context, request model.UserUpdateRequest) (response model.UserResponse, err error)
	RemoveUser(ctx context.Context, userID string) error
}
