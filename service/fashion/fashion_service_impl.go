package fashion

import (
	"context"

	"github.com/vnnyx/gas-api/model"
	"github.com/vnnyx/gas-api/repository/fashion"
)

type FashionServiceImpl struct {
	fashion.FashionRepository
}

func NewFashionService(fashionRepository fashion.FashionRepository) FashionService {
	return &FashionServiceImpl{FashionRepository: fashionRepository}
}

func (service *FashionServiceImpl) CreateFashion(ctx context.Context, request model.FashionCreateRequest) (response model.FashionResponse, err error) {
	panic("implement me")
}

func (service *FashionServiceImpl) GetFashionByID(ctx context.Context, fashionID string) (response model.FashionResponse, err error) {
	panic("implement me")
}

func (service *FashionServiceImpl) GetAllFashion(ctx context.Context) (response []model.FashionResponse, err error) {
	panic("implement me")
}

func (service *FashionServiceImpl) UpdateFashion(ctx context.Context, request model.FashionUpdateRequest) (response model.FashionResponse, err error) {
	panic("implement me")
}

func (service *FashionServiceImpl) RemoveFashion(ctx context.Context, fashionID string) error {
	panic("implement me")
}
