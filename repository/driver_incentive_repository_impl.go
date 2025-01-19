package repository

import (
	"context"

	"github.com/DevisArya/car-rental/models"
	"gorm.io/gorm"
)

type DriverIncentiveRepositoryImpl struct {
}

func NewDriverIncentiveRepository() DriverIncentiveRepository {
	return &DriverIncentiveRepositoryImpl{}
}

// FindById implements DriverIncentiveRepository
func (*DriverIncentiveRepositoryImpl) FindById(ctx context.Context, db *gorm.DB, driver_incentive_id uint) (*models.DriverIncentive, error) {
	var driverIncentive models.DriverIncentive

	if err := db.WithContext(ctx).
		First(&driverIncentive, driver_incentive_id).
		Error; err != nil {
		return nil, err
	}

	return &driverIncentive, nil
}

// FindAll implements DriverIncentiveRepository
func (*DriverIncentiveRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB, limit int, offset int) (*[]models.DriverIncentive, int64, error) {
	var driverIncentives []models.DriverIncentive
	var totalCount int64

	if err := db.WithContext(ctx).
		Model(&models.DriverIncentive{}).
		Count(&totalCount).
		Error; err != nil {
		return nil, 0, err
	}

	if err := db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&driverIncentives).
		Error; err != nil {
		return nil, 0, err
	}

	if len(driverIncentives) == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}

	return &driverIncentives, totalCount, nil
}
