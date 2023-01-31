package kelontong

import (
	"context"

	"github.com/vnnyx/gas-api/model"
)

type KelontongService interface {
	CreateKelontong(ctx context.Context, request model.KelontongCreateRequest) (response model.KelontongResponse, err error)
	GetKelontongByID(ctx context.Context, kelontongID string) (response model.KelontongResponse, err error)
	GetKelontongByName(ctx context.Context, name string) (response []model.KelontongResponse, err error)
	GetKelontongByLocation(ctx context.Context, location string) (response []model.KelontongResponse, err error)
	GetAllKelontong(ctx context.Context) (response []model.KelontongResponse, err error)
	UpdateKelontong(ctx context.Context, request model.KelontongUpdateRequest) (response model.KelontongResponse, err error)
	RemoveKelontong(ctx context.Context, kelontongID string) error
}
