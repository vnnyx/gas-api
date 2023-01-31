package kelontong

import (
	"context"

	"github.com/vnnyx/gas-api/model"
)

type KelontongRepository interface {
	InsertKelontong(ctx context.Context, kelontong model.Kelontong) (model.Kelontong, error)
	FindKelontongById(ctx context.Context, kelontongID string) (kelontong model.Kelontong, err error)
	FindKelontongByName(ctx context.Context, name string) (listKelontong []model.Kelontong, err error)
	FindKelontongByLocation(ctx context.Context, location string) (listKelontong []model.Kelontong, err error)
	FindAllKelontong(ctx context.Context) (listKelontong []model.Kelontong, err error)
	UpdateKelontong(ctx context.Context, kelontong model.Kelontong) (model.Kelontong, error)
	DeleteKelontong(ctx context.Context, kelontongID string) error
}
