package routes

import (
	"github.com/DevisArya/car-rental/handler"
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, AppHandler *handler.AppHandler) {

	// Regist All Routes
	RegisterCustomerRoutes(e, AppHandler.CustomerHandler)
	RegisterCarRoutes(e, AppHandler.CarHandler)
	RegisterBookingRoutes(e, AppHandler.BookingHandler)
	RegisterMembershipRoutes(e, AppHandler.MembershipHandler)
	RegisterDriverRoutes(e, AppHandler.DriverHandler)
}
