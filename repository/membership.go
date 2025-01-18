package repository

import (
	"context"

	"github.com/DevisArya/car-rental/models"
	"gorm.io/gorm"
)

type MembershipRepository interface {
	Create(ctx context.Context, tx *gorm.DB, membership *models.Membership) (*models.Membership, error)
	Update(ctx context.Context, tx *gorm.DB, membership *models.Membership) error
	Delete(ctx context.Context, tx *gorm.DB, membershipId uint) error
	FindById(ctx context.Context, db *gorm.DB, membershipId uint) (*models.Membership, error)
	FindAll(ctx context.Context, db *gorm.DB, limit int, offset int) (*[]models.Membership, int64, error)
}
