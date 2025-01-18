package repository

import (
	"context"

	"github.com/DevisArya/car-rental/models"
	"gorm.io/gorm"
)

type CarRepository interface {
	Create(ctx context.Context, tx *gorm.DB, car *models.Car) (*models.Car, error)
	Update(ctx context.Context, tx *gorm.DB, car *models.Car) error
	Delete(ctx context.Context, tx *gorm.DB, carId uint) error
	FindById(ctx context.Context, db *gorm.DB, carId uint) (*models.Car, error)
	FindAll(ctx context.Context, db *gorm.DB, limit int, offset int) (*[]models.Car, int64, error)
}
