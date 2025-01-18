package routes

import (
	"github.com/DevisArya/car-rental/handler"
	"github.com/labstack/echo/v4"
)

func RegisterBookingRoutes(e *echo.Echo, bookingHandler handler.BookingHandler) {

	// route for booking
	v1 := e.Group("/v1")

	v1.POST("/booking", bookingHandler.Create)
	v1.PATCH("/booking/:id", bookingHandler.Update)
	v1.PATCH("/booking/status/:id", bookingHandler.UpdateStatus)
	v1.DELETE("/booking/:id", bookingHandler.Delete)
	v1.GET("/booking/:id", bookingHandler.FindById)
	v1.GET("/bookings", bookingHandler.FindAll)
}
