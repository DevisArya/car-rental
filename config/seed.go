package config

import (
	"github.com/DevisArya/car-rental/helper"
	"github.com/DevisArya/car-rental/models"
	"gorm.io/gorm"
)

func SeedBookingTypes(db *gorm.DB) {
	bookingTypes := []models.BookingType{
		{
			BookingType: "Car Only",
			Description: "Rent Car Only",
		},
		{
			BookingType: "Car & Driver",
			Description: "Rent Car and a Driver",
		},
	}

	for _, bookingType := range bookingTypes {
		//cek apakah booking type sudah ada
		var count int64
		if err := db.Model(&models.BookingType{}).Where("booking_type_id = ?", bookingType.BookingTypeID).Count(&count).Error; err != nil {
			helper.PanicIfError(err)
		}

		if count == 0 {
			// Insert booking type jika tidak ada
			if err := db.Create(&bookingType).Error; err != nil {
				helper.PanicIfError(err)
			}
		}
	}
}
