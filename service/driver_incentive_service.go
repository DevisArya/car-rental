package service

import (
	"context"

	"github.com/DevisArya/car-rental/dto"
)

type DriverIncentiveService interface {
	FindById(ctx context.Context, driverIncentiveId uint) (*dto.DriverIncentiveResponse, error)
	FindAll(ctx context.Context, limit, offset int) (*[]dto.DriverIncentiveResponse, int, int, error)
}
