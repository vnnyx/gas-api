package fashion

import (
	"context"

	"github.com/vnnyx/gas-api/model"
)

type FashionService interface {
	CreateFashion(ctx context.Context, request model.FashionCreateRequest) (response model.FashionResponse, err error)
	GetFashionByID(ctx context.Context, fashionID string) (response model.FashionResponse, err error)
	GetAllFashion(ctx context.Context) (response []model.FashionResponse, err error)
	UpdateFashion(ctx context.Context, request model.FashionUpdateRequest) (response model.FashionResponse, err error)
	RemoveFashion(ctx context.Context, fashionID string) error
}
