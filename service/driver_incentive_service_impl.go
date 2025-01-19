package service

import (
	"context"
	"errors"
	"math"

	"github.com/DevisArya/car-rental/dto"
	"github.com/DevisArya/car-rental/helper"
	"github.com/DevisArya/car-rental/repository"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type DriverIncentiveServiceImpl struct {
	DriverIncentiveRepository repository.DriverIncentiveRepository
	DB                        *gorm.DB
	validate                  *validator.Validate
}

func NewDriverIncentiveService(driverIncentiveRepository repository.DriverIncentiveRepository, DB *gorm.DB, validate *validator.Validate) DriverIncentiveService {

	return &DriverIncentiveServiceImpl{
		DriverIncentiveRepository: driverIncentiveRepository,
		DB:                        DB,
		validate:                  validate,
	}
}

// FindById implements DriverIncentiveService
func (service *DriverIncentiveServiceImpl) FindById(ctx context.Context, driverIncentiveId uint) (*dto.DriverIncentiveResponse, error) {

	//validasi input
	if driverIncentiveId <= 0 {
		return nil, errors.New("invalid driverIncentive id")
	}

	driverIncentive, err := service.DriverIncentiveRepository.FindById(ctx, service.DB, driverIncentiveId)
	if err != nil {
		return nil, err
	}

	return helper.ToDriverIncentiveResponse(driverIncentive), nil
}

// FindAll implements DriverIncentiveService
func (service *DriverIncentiveServiceImpl) FindAll(ctx context.Context, limit, offset int) (*[]dto.DriverIncentiveResponse, int, int, error) {

	driverIncentives, totalCount, err := service.DriverIncentiveRepository.FindAll(ctx, service.DB, limit, offset)
	if err != nil {
		return nil, 0, 0, err
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))

	return helper.ToDriverIncentiveResponses(driverIncentives), int(totalCount), totalPages, nil
}
