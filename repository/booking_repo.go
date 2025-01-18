package repository

import (
	"context"

	"github.com/DevisArya/car-rental/models"
	"gorm.io/gorm"
)

type BookingRepository interface {
	Create(ctx context.Context, tx *gorm.DB, booking *models.Booking) (*models.Booking, error)
	Update(ctx context.Context, tx *gorm.DB, booking *models.Booking) error
	Delete(ctx context.Context, tx *gorm.DB, bookingId uint) error
	FindById(ctx context.Context, db *gorm.DB, bookingId uint) (*models.Booking, error)
	FindAll(ctx context.Context, db *gorm.DB, limit int, offset int) (*[]models.Booking, int64, error)
}
