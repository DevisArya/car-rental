package config

import (
	"github.com/DevisArya/car-rental/helper"
	"github.com/DevisArya/car-rental/models"
	"gorm.io/gorm"
)

func InitialMigration(db *gorm.DB) {

	err := db.AutoMigrate(
		models.Customer{},
		models.Car{},
		models.Booking{},
		models.Membership{},
		models.Driver{},
	)

	helper.PanicIfError(err)
}
