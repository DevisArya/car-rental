package repository

import (
	"context"

	"github.com/DevisArya/car-rental/models"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(ctx context.Context, tx *gorm.DB, customer *models.Customer) (*models.Customer, error)
	Update(ctx context.Context, tx *gorm.DB, customer *models.Customer) error
	Delete(ctx context.Context, tx *gorm.DB, customerId uint) error
	FindById(ctx context.Context, db *gorm.DB, customerId uint) (*models.Customer, error)
	FindByNikAndPhoneNumber(ctx context.Context, db *gorm.DB, PhoneNumber string, Nik string) (bool, error)
	FindAll(ctx context.Context, db *gorm.DB, limit int, offset int) (*[]models.Customer, int64, error)
	FindByNik(ctx context.Context, db *gorm.DB, Nik string) (bool, error)
	FindByPhoneNumber(ctx context.Context, db *gorm.DB, PhoneNumber string) (bool, error)
}
