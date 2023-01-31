package kelontong

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/vnnyx/gas-api/model"
	"github.com/vnnyx/gas-api/repository/kelontong"
)

type KelontongServiceImpl struct {
	kelontong.KelontongRepository
}

func NewKelontongService(kelontongRepository kelontong.KelontongRepository) KelontongService {
	return &KelontongServiceImpl{KelontongRepository: kelontongRepository}
}

func (service *KelontongServiceImpl) CreateKelontong(ctx context.Context, request model.KelontongCreateRequest) (response model.KelontongResponse, err error) {
	kelontong := model.Kelontong{
		ID:       uuid.New().String(),
		Name:     request.Name,
		Location: request.Location,
	}

	kelontong, err = service.KelontongRepository.InsertKelontong(ctx, kelontong)
	if err != nil {
		return response, err
	}

	response = model.KelontongResponse{
		ID:       kelontong.ID,
		Name:     kelontong.Name,
		Location: kelontong.Location,
	}

	return response, nil
}

func (service *KelontongServiceImpl) GetKelontongByID(ctx context.Context, kelontongID string) (response model.KelontongResponse, err error) {
	kelontong, err := service.KelontongRepository.FindKelontongById(ctx, kelontongID)
	if err != nil {
		return response, errors.New("KELONTONG_NOT_FOUND")
	}

	response = model.KelontongResponse{
		ID:       kelontong.ID,
		Name:     kelontong.Name,
		Location: kelontong.Location,
	}

	return response, nil
}

func (service *KelontongServiceImpl) GetKelontongByName(ctx context.Context, name string) (response []model.KelontongResponse, err error) {
	listKelontong, err := service.KelontongRepository.FindKelontongByName(ctx, name)
	if err != nil {
		return response, err
	}

	for _, kelontong := range listKelontong {
		response = append(response, model.KelontongResponse{
			ID:       kelontong.ID,
			Name:     kelontong.Name,
			Location: kelontong.Location,
		})
	}

	return response, nil
}

func (service *KelontongServiceImpl) GetKelontongByLocation(ctx context.Context, location string) (response []model.KelontongResponse, err error) {
	listKelontong, err := service.KelontongRepository.FindKelontongByLocation(ctx, location)
	if err != nil {
		return response, err
	}

	for _, kelontong := range listKelontong {
		response = append(response, model.KelontongResponse{
			ID:       kelontong.ID,
			Name:     kelontong.Name,
			Location: kelontong.Location,
		})
	}

	return response, nil
}

func (service *KelontongServiceImpl) GetAllKelontong(ctx context.Context) (response []model.KelontongResponse, err error) {
	listKelontong, err := service.KelontongRepository.FindAllKelontong(ctx)
	if err != nil {
		return response, err
	}

	for _, kelontong := range listKelontong {
		response = append(response, model.KelontongResponse{
			ID:       kelontong.ID,
			Name:     kelontong.Name,
			Location: kelontong.Location,
		})
	}

	return response, nil
}

func (service *KelontongServiceImpl) UpdateKelontong(ctx context.Context, request model.KelontongUpdateRequest) (response model.KelontongResponse, err error) {
	kelontong, err := service.KelontongRepository.FindKelontongById(ctx, request.ID)
	if err != nil {
		return response, errors.New("KELONTONG_NOT_FOUND")
	}

	kelontong, err = service.KelontongRepository.UpdateKelontong(ctx, model.Kelontong{
		ID:       kelontong.ID,
		Name:     request.Name,
		Location: request.Location,
	})

	if err != nil {
		return response, err
	}

	response = model.KelontongResponse{
		ID:       kelontong.ID,
		Name:     kelontong.Name,
		Location: kelontong.Location,
	}

	return response, nil
}

func (service *KelontongServiceImpl) RemoveKelontong(ctx context.Context, kelontongID string) error {
	return service.KelontongRepository.DeleteKelontong(ctx, kelontongID)
}
