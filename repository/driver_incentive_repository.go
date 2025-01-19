package repository

import (
	"context"

	"github.com/DevisArya/car-rental/models"
	"gorm.io/gorm"
)

type DriverIncentiveRepository interface {
	FindById(ctx context.Context, db *gorm.DB, driverIncentiveId uint) (*models.DriverIncentive, error)
	FindAll(ctx context.Context, db *gorm.DB, limit int, offset int) (*[]models.DriverIncentive, int64, error)
}
