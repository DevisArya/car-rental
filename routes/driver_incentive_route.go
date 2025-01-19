package routes

import (
	"github.com/DevisArya/car-rental/handler"
	"github.com/labstack/echo/v4"
)

func RegisterDriverIncentiveRoutes(e *echo.Echo, driverIncentiveHandler handler.DriverIncentiveHandler) {

	// route for driverIncentive
	v2 := e.Group("/v2")

	v2.GET("/driver-incentive/:id", driverIncentiveHandler.FindById)
	v2.GET("/driver-incentives", driverIncentiveHandler.FindAll)
}
