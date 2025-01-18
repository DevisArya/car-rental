package repository

import (
	"context"

	"github.com/DevisArya/car-rental/models"
	"gorm.io/gorm"
)

type DriverRepository interface {
	Create(ctx context.Context, tx *gorm.DB, driver *models.Driver) (*models.Driver, error)
	Update(ctx context.Context, tx *gorm.DB, driver *models.Driver) error
	Delete(ctx context.Context, tx *gorm.DB, driverId uint) error
	FindById(ctx context.Context, db *gorm.DB, driverId uint) (*models.Driver, error)
	FindByNikAndPhoneNumber(ctx context.Context, db *gorm.DB, PhoneNumber string, Nik string) (bool, error)
	FindAll(ctx context.Context, db *gorm.DB, limit int, offset int) (*[]models.Driver, int64, error)
}
