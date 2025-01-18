package service

import (
	"context"

	"github.com/DevisArya/car-rental/dto"
)

type DriverService interface {
	Create(ctx context.Context, request *dto.DriverRequest) (*dto.DriverResponse, error)
	Update(ctx context.Context, request *dto.DriverRequest, driverId uint) error
	Delete(ctx context.Context, driverId uint) error
	FindById(ctx context.Context, driverId uint) (*dto.DriverResponse, error)
	FindAll(ctx context.Context, limit, offset int) (*[]dto.DriverResponse, int, int, error)
}
