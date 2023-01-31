package fashion

import (
	"context"

	"github.com/vnnyx/gas-api/model"
)

type FashionRepository interface {
	InsertFashion(ctx context.Context, fashion model.Fashion) (model.Fashion, error)
	FindFashionById(ctx context.Context, fashionID string) (fashion model.Fashion, err error)
	FindFashionByName(ctx context.Context, name string) (listFashion []model.Fashion, err error)
	FindAllFashion(ctx context.Context) (listFashion []model.Fashion, err error)
	UpdateFashion(ctx context.Context, fashion model.Fashion) (model.Fashion, error)
	DeleteFashion(ctx context.Context, fashionID string) error
}
