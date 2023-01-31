package fashion

import (
	"context"

	"github.com/vnnyx/gas-api/model"
	"gorm.io/gorm"
)

type FashionRepositoryImpl struct {
	DB *gorm.DB
}

func NewFashionRepository(DB *gorm.DB) FashionRepository {
	return &FashionRepositoryImpl{DB: DB}
}

func (repository *FashionRepositoryImpl) InsertFashion(ctx context.Context, fashion model.Fashion) (model.Fashion, error) {
	err := repository.DB.WithContext(ctx).Create(&fashion).Error
	return fashion, err
}

func (repository *FashionRepositoryImpl) FindFashionById(ctx context.Context, fashionID string) (fashion model.Fashion, err error) {
	err = repository.DB.WithContext(ctx).Where("fashion_id", fashionID).First(&fashion).Error
	return fashion, err
}

func (repository *FashionRepositoryImpl) FindFashionByName(ctx context.Context, name string) (listFashion []model.Fashion, err error) {
	err = repository.DB.WithContext(ctx).Where("name", name).Find(&listFashion).Error
	return listFashion, err
}

func (repository *FashionRepositoryImpl) FindAllFashion(ctx context.Context) (listFashion []model.Fashion, err error) {
	err = repository.DB.WithContext(ctx).Find(&listFashion).Error
	return listFashion, err
}

func (repository *FashionRepositoryImpl) UpdateFashion(ctx context.Context, fashion model.Fashion) (model.Fashion, error) {
	err := repository.DB.WithContext(ctx).Where("fashion_id", fashion.ID).Updates(&fashion).Error
	return fashion, err
}

func (repository *FashionRepositoryImpl) DeleteFashion(ctx context.Context, fashionID string) error {
	return repository.DB.WithContext(ctx).Where("fashion_id", fashionID).Delete(&model.Fashion{}).Error
}
