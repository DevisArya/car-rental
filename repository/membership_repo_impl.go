package repository

import (
	"context"

	"github.com/DevisArya/car-rental/models"
	"gorm.io/gorm"
)

type MembershipRepositoryImpl struct {
}

func NewMembershipRepository() MembershipRepository {
	return &MembershipRepositoryImpl{}
}

// Save implements MembershipRepository
func (*MembershipRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, dataMembership *models.Membership) (*models.Membership, error) {

	if err := tx.WithContext(ctx).
		Create(dataMembership).
		Error; err != nil {
		return nil, err
	}

	return dataMembership, nil
}

// Update implements MembershipRepository
func (*MembershipRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, dataMembership *models.Membership) error {

	if err := tx.WithContext(ctx).
		Model(&models.Membership{}).
		Where("membership_id = ?", dataMembership.MembershipID).
		Updates(dataMembership).
		Error; err != nil {
		return err
	}

	return nil
}

// Delete implements MembershipRepository
func (*MembershipRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, membershipId uint) error {

	if err := tx.WithContext(ctx).
		Delete(&models.Membership{}, membershipId).
		Error; err != nil {
		return err
	}

	return nil
}

// FindById implements MembershipRepository
func (*MembershipRepositoryImpl) FindById(ctx context.Context, db *gorm.DB, membershipId uint) (*models.Membership, error) {
	var membership models.Membership

	if err := db.WithContext(ctx).
		First(&membership, membershipId).
		Error; err != nil {
		return nil, err
	}

	return &membership, nil
}

// FindAll implements MembershipRepository
func (*MembershipRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB, limit int, offset int) (*[]models.Membership, int64, error) {
	var memberships []models.Membership
	var totalCount int64

	if err := db.WithContext(ctx).
		Model(&models.Membership{}).
		Count(&totalCount).
		Error; err != nil {
		return nil, 0, err
	}

	if err := db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&memberships).
		Error; err != nil {
		return nil, 0, err
	}

	if len(memberships) == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}

	return &memberships, totalCount, nil
}
