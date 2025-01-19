package config

import (
	"github.com/DevisArya/car-rental/helper"
	"github.com/DevisArya/car-rental/models"
	"gorm.io/gorm"
)

func InitialMigration(db *gorm.DB) {

	err := db.AutoMigrate(
		models.Booking{},
		models.BookingType{},
		models.Customer{},
		models.Membership{},
		models.Car{},
		models.Driver{},
		models.DriverIncentive{},
	)

	helper.PanicIfError(err)
}
