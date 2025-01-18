package repository

import (
	"context"

	"github.com/DevisArya/car-rental/models"
	"gorm.io/gorm"
)

type CarRepositoryImpl struct {
}

func NewCarRepository() CarRepository {
	return &CarRepositoryImpl{}
}

// Save implements CarRepository
func (*CarRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, dataCar *models.Car) (*models.Car, error) {

	if err := tx.WithContext(ctx).
		Create(dataCar).
		Error; err != nil {
		return nil, err
	}

	return dataCar, nil
}

// Update implements CarRepository
func (*CarRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, dataCar *models.Car) error {

	if err := tx.WithContext(ctx).
		Model(&models.Car{}).
		Where("car_id = ?", dataCar.CarID).
		Updates(dataCar).
		Error; err != nil {
		return err
	}

	return nil
}

// Delete implements CarRepository
func (*CarRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, carId uint) error {

	if err := tx.WithContext(ctx).
		Delete(&models.Car{}, carId).
		Error; err != nil {
		return err
	}

	return nil
}

// FindById implements CarRepository
func (*CarRepositoryImpl) FindById(ctx context.Context, db *gorm.DB, carId uint) (*models.Car, error) {
	var car models.Car

	if err := db.WithContext(ctx).
		First(&car, carId).
		Error; err != nil {
		return nil, err
	}

	return &car, nil
}

// FindAll implements CarRepository
func (*CarRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB, limit int, offset int) (*[]models.Car, int64, error) {
	var cars []models.Car
	var totalCount int64

	if err := db.WithContext(ctx).
		Model(&models.Car{}).
		Count(&totalCount).
		Error; err != nil {
		return nil, 0, err
	}

	if err := db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&cars).
		Error; err != nil {
		return nil, 0, err
	}

	if len(cars) == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}

	return &cars, totalCount, nil
}
