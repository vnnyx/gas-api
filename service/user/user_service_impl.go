package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/vnnyx/gas-api/model"
	"github.com/vnnyx/gas-api/repository/user"
	"github.com/vnnyx/gas-api/validation"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	user.UserRepository
	*gorm.DB
}

func NewUserService(userRepository user.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

func (service *UserServiceImpl) CreateUser(ctx context.Context, request model.UserCreateRequest) (response model.UserResponse, err error) {
	validation.CreateUSerValidation(request)

	if request.Password != request.PasswordConfirmation {
		return response, errors.New("PASSWORD_NOT_MATCH")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return response, err
	}

	user := model.User{
		ID:        uuid.New().String(),
		Username:  request.Username,
		Email:     request.Email,
		Handphone: request.Handphone,
		Password:  string(password),
	}

	user, err = service.UserRepository.InsertUser(ctx, user)
	if err != nil {
		return response, err
	}

	response = model.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Handphone: user.Handphone,
	}

	return response, nil

}

func (service *UserServiceImpl) GetUserByID(ctx context.Context, userID string) (response model.UserResponse, err error) {
	user, err := service.UserRepository.FindUserById(ctx, userID)
	if err != nil {
		return response, err
	}

	response = model.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Handphone: user.Handphone,
	}

	return response, nil
}

func (service *UserServiceImpl) GetAllUser(ctx context.Context) (response []model.UserResponse, err error) {
	users, err := service.UserRepository.FindAllUser(ctx)
	if err != nil {
		return response, err
	}

	for _, user := range users {
		response = append(response, model.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Handphone: user.Handphone,
		})
	}

	return response, nil
}

func (service *UserServiceImpl) UpdateUserProfile(ctx context.Context, request model.UserUpdateRequest) (response model.UserResponse, err error) {
	user, err := service.UserRepository.FindUserById(ctx, request.ID)

	if err != nil {
		return response, errors.New("USER_NOT_FOUND")
	}

	user, err = service.UserRepository.UpdateUser(ctx, model.User{
		ID:        user.ID,
		Username:  request.Username,
		Email:     response.Email,
		Handphone: request.Handphone,
	})

	if err != nil {
		return response, err
	}

	response = model.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Handphone: user.Handphone,
	}

	return response, nil
}

func (service *UserServiceImpl) RemoveUser(ctx context.Context, userID string) error {
	return service.UserRepository.DeleteUser(ctx, userID)
}
