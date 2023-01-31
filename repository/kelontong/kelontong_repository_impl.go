package kelontong

import (
	"context"

	"github.com/vnnyx/gas-api/model"
	"gorm.io/gorm"
)

type KelontongRepositoryImpl struct {
	DB *gorm.DB
}

func NewKelontongRepository(DB *gorm.DB) KelontongRepository {
	return &KelontongRepositoryImpl{DB: DB}
}

func (repository *KelontongRepositoryImpl) InsertKelontong(ctx context.Context, kelontong model.Kelontong) (model.Kelontong, error) {
	err := repository.DB.WithContext(ctx).Create(&kelontong).Error
	return kelontong, err
}

func (repository *KelontongRepositoryImpl) FindKelontongById(ctx context.Context, kelontongID string) (kelontong model.Kelontong, err error) {
	err = repository.DB.WithContext(ctx).Where("kelontong_id", kelontongID).First(&kelontong).Error
	return kelontong, err
}

func (repository *KelontongRepositoryImpl) FindKelontongByName(ctx context.Context, name string) (listKelontong []model.Kelontong, err error) {
	err = repository.DB.WithContext(ctx).Where("name", name).Find(&listKelontong).Error
	return listKelontong, err
}

func (repository *KelontongRepositoryImpl) FindKelontongByLocation(ctx context.Context, location string) (listKelontong []model.Kelontong, err error) {
	err = repository.DB.WithContext(ctx).Where("location", location).Find(&listKelontong).Error
	return listKelontong, err
}

func (repository *KelontongRepositoryImpl) FindAllKelontong(ctx context.Context) (listKelontong []model.Kelontong, err error) {
	err = repository.DB.WithContext(ctx).Find(&listKelontong).Error
	return listKelontong, err
}

func (repository *KelontongRepositoryImpl) UpdateKelontong(ctx context.Context, kelontong model.Kelontong) (model.Kelontong, error) {
	err := repository.DB.WithContext(ctx).Where("kelontong_id", kelontong.ID).Updates(&kelontong).Error
	return kelontong, err
}

func (repository *KelontongRepositoryImpl) DeleteKelontong(ctx context.Context, kelontongID string) error {
	return repository.DB.WithContext(ctx).Where("kelontong_id", kelontongID).Delete(&model.Kelontong{}).Error
}
