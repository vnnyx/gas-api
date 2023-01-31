package user

import (
	"context"

	"github.com/vnnyx/gas-api/model"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: DB}
}

func (repository *UserRepositoryImpl) InsertUser(ctx context.Context, user model.User) (model.User, error) {
	err := repository.DB.WithContext(ctx).Create(&user).Error
	return user, err
}

func (repository *UserRepositoryImpl) FindUserById(ctx context.Context, userID string) (user model.User, err error) {
	err = repository.DB.WithContext(ctx).Where("user_id", userID).First(&user).Error
	return user, err
}

func (repository *UserRepositoryImpl) FindUserByUsername(ctx context.Context, username string) (user model.User, err error) {
	err = repository.DB.WithContext(ctx).Where("username", username).First(&user).Error
	return user, err
}

func (repository *UserRepositoryImpl) FindAllUser(ctx context.Context) (users []model.User, err error) {
	err = repository.DB.WithContext(ctx).Find(&users).Error
	return users, err
}

func (repository *UserRepositoryImpl) UpdateUser(ctx context.Context, user model.User) (model.User, error) {
	err := repository.DB.WithContext(ctx).Where("user_id", user.ID).Updates(&user).Error
	return user, err
}

func (repository *UserRepositoryImpl) DeleteUser(ctx context.Context, userID string) error {
	return repository.DB.WithContext(ctx).Where("user_id", userID).Delete(&model.User{}).Error
}
