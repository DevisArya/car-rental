package repository

import (
	"context"

	"github.com/DevisArya/car-rental/models"
	"gorm.io/gorm"
)

type DriverRepositoryImpl struct {
}

func NewDriverRepository() DriverRepository {
	return &DriverRepositoryImpl{}
}

// Save implements DriverRepository
func (*DriverRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, dataDriver *models.Driver) (*models.Driver, error) {

	if err := tx.WithContext(ctx).
		Create(dataDriver).
		Error; err != nil {
		return nil, err
	}

	return dataDriver, nil
}

// Update implements DriverRepository
func (*DriverRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, dataDriver *models.Driver) error {

	if err := tx.WithContext(ctx).
		Model(&models.Driver{}).
		Where("driver_id = ?", dataDriver.DriverID).
		Updates(dataDriver).
		Error; err != nil {
		return err
	}

	return nil
}

// Delete implements DriverRepository
func (*DriverRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, driverId uint) error {

	if err := tx.WithContext(ctx).
		Delete(&models.Driver{}, driverId).
		Error; err != nil {
		return err
	}

	return nil
}

// FindById implements DriverRepository
func (*DriverRepositoryImpl) FindById(ctx context.Context, db *gorm.DB, driverId uint) (*models.Driver, error) {
	var driver models.Driver

	if err := db.WithContext(ctx).
		First(&driver, driverId).
		Error; err != nil {
		return nil, err
	}

	return &driver, nil
}

// FindAll implements DriverRepository
func (*DriverRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB, limit int, offset int) (*[]models.Driver, int64, error) {
	var drivers []models.Driver
	var totalCount int64

	if err := db.WithContext(ctx).
		Model(&models.Driver{}).
		Count(&totalCount).
		Error; err != nil {
		return nil, 0, err
	}

	if err := db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&drivers).
		Error; err != nil {
		return nil, 0, err
	}

	if len(drivers) == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}

	return &drivers, totalCount, nil
}

// FindByNikAndPhoneNumber implements DriverRepository
func (*DriverRepositoryImpl) FindByNikAndPhoneNumber(ctx context.Context, db *gorm.DB, PhoneNumber string, Nik string) (bool, error) {
	var driver models.Driver

	err := db.WithContext(ctx).
		Where("nik = ? OR phone_number = ?", Nik, PhoneNumber).
		First(&driver).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return true, nil
		}
		return false, err
	}

	return false, nil
}

// // FindById implements DriverRepository
// func (*DriverRepositoryImpl) FindDriverIncentive(ctx context.Context, db *gorm.DB, driverId uint, date *dto.FindDriverIncentiveRequest) (*models.Driver, error) {
// 	var driver models.Driver

// 	if err := db.WithContext(ctx).
// 		Preload("Bookings", "start_date BETWEEN ? AND ?", date.StartDate, date.EndDate).
// 		First(&driver, driverId).
// 		Error; err != nil {
// 		return nil, err
// 	}

// 	return &driver, nil
// }
