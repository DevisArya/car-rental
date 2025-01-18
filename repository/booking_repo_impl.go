package repository

import (
	"context"

	"github.com/DevisArya/car-rental/models"
	"gorm.io/gorm"
)

type BookingRepositoryImpl struct {
}

func NewBookingRepository() BookingRepository {
	return &BookingRepositoryImpl{}
}

// Save implements BookingRepository
func (*BookingRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, dataBooking *models.Booking) (*models.Booking, error) {

	if err := tx.WithContext(ctx).
		Create(dataBooking).
		Error; err != nil {
		return nil, err
	}

	return dataBooking, nil
}

// Update implements BookingRepository
func (*BookingRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, dataBooking *models.Booking) error {

	if err := tx.WithContext(ctx).
		Model(&models.Booking{}).
		Where("booking_id = ?", dataBooking.BookingID).
		Updates(dataBooking).
		Error; err != nil {
		return err
	}

	return nil
}

// Delete implements BookingRepository
func (*BookingRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, bookingId uint) error {

	if err := tx.WithContext(ctx).
		Delete(&models.Booking{}, bookingId).
		Error; err != nil {
		return err
	}

	return nil
}

// FindById implements BookingRepository
func (*BookingRepositoryImpl) FindById(ctx context.Context, db *gorm.DB, bookingId uint) (*models.Booking, error) {
	var booking models.Booking

	if err := db.WithContext(ctx).
		First(&booking, bookingId).
		Error; err != nil {
		return nil, err
	}

	return &booking, nil
}

// FindAll implements BookingRepository
func (*BookingRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB, limit int, offset int) (*[]models.Booking, int64, error) {
	var bookings []models.Booking
	var totalCount int64

	if err := db.WithContext(ctx).
		Model(&models.Booking{}).
		Count(&totalCount).
		Error; err != nil {
		return nil, 0, err
	}

	if err := db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&bookings).
		Error; err != nil {
		return nil, 0, err
	}

	if len(bookings) == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}

	return &bookings, totalCount, nil
}
