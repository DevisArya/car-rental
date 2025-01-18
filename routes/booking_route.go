package routes

import (
	"github.com/DevisArya/car-rental/handler"
	"github.com/labstack/echo/v4"
)

func RegisterBookingRoutes(e *echo.Echo, bookingHandler handler.BookingHandler) {

	// route for booking
	v2 := e.Group("/v2")

	v2.POST("/booking", bookingHandler.Create)
	v2.PATCH("/booking/:id", bookingHandler.Update)
	v2.PATCH("/booking/status/:id", bookingHandler.UpdateStatus)
	v2.DELETE("/booking/:id", bookingHandler.Delete)
	v2.GET("/booking/:id", bookingHandler.FindById)
	v2.GET("/bookings", bookingHandler.FindAll)
}
