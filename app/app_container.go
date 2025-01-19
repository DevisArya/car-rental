package app

import (
	"github.com/DevisArya/car-rental/handler"
	"github.com/DevisArya/car-rental/repository"
	"github.com/DevisArya/car-rental/service"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func NewAppContainer(db *gorm.DB, validate *validator.Validate) *handler.AppHandler {
	carRepository := repository.NewCarRepository()
	carService := service.NewCarService(carRepository, db, validate)
	carHandler := handler.NewCarHandler(carService)

	membershipRepository := repository.NewMembershipRepository()
	membershipService := service.NewMembershipService(membershipRepository, db, validate)
	membershipHandler := handler.NewMembershipHandler(membershipService)

	driverRepository := repository.NewDriverRepository()
	driverService := service.NewDriverService(driverRepository, db, validate)
	driverHandler := handler.NewDriverHandler(driverService)

	driverIncentiveRepository := repository.NewDriverIncentiveRepository()
	driverIncentiveService := service.NewDriverIncentiveService(driverIncentiveRepository, db, validate)
	driverIncentiveHandler := handler.NewDriverIncentiveHandler(driverIncentiveService)

	customerRepository := repository.NewCustomerRepository()
	customerService := service.NewCustomerService(customerRepository, membershipService, db, validate)
	customerHandler := handler.NewCustomerHandler(customerService)

	bookingRepository := repository.NewBookingRepository()
	bookingService := service.NewBookingService(bookingRepository, carService, customerService, driverService, db, validate)
	bookingHandler := handler.NewBookingHandler(bookingService)

	return &handler.AppHandler{
		CustomerHandler:        customerHandler,
		CarHandler:             carHandler,
		BookingHandler:         bookingHandler,
		MembershipHandler:      membershipHandler,
		DriverHandler:          driverHandler,
		DriverIncentiveHandler: driverIncentiveHandler,
	}
}
