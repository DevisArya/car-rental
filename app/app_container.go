package app

import (
	"github.com/DevisArya/car-rental/handler"
	"github.com/DevisArya/car-rental/repository"
	"github.com/DevisArya/car-rental/service"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func NewAppContainer(db *gorm.DB, validate *validator.Validate) *handler.AppHandler {

	customerRepository := repository.NewCustomerRepository()
	customerService := service.NewCustomerService(customerRepository, db, validate)
	customerHandler := handler.NewCustomerHandler(customerService)

	carRepository := repository.NewCarRepository()
	carService := service.NewCarService(carRepository, db, validate)
	carHandler := handler.NewCarHandler(carService)

	bookingRepository := repository.NewBookingRepository()
	bookingService := service.NewBookingService(bookingRepository, carService, db, validate)
	bookingHandler := handler.NewBookingHandler(bookingService)

	return &handler.AppHandler{
		CustomerHandler: customerHandler,
		CarHandler:      carHandler,
		BookingHandler:  bookingHandler,
	}
}
