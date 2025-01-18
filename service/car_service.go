package service

import (
	"context"

	"github.com/DevisArya/car-rental/dto"
)

type CarService interface {
	Create(ctx context.Context, request *dto.CarRequest) (*dto.CarResponse, error)
	Update(ctx context.Context, request *dto.CarRequest, carId uint) error
	Delete(ctx context.Context, carId uint) error
	FindById(ctx context.Context, carId uint) (*dto.CarResponse, error)
	FindAll(ctx context.Context, limit, offset int) (*[]dto.CarResponse, int, int, error)
}
