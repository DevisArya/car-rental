package routes

import (
	"github.com/DevisArya/car-rental/handler"
	"github.com/labstack/echo/v4"
)

func RegisterDriverRoutes(e *echo.Echo, driverHandler handler.DriverHandler) {

	// route for driver
	v2 := e.Group("/v2")

	v2.POST("/driver", driverHandler.Create)
	v2.PATCH("/driver/:id", driverHandler.Update)
	v2.DELETE("/driver/:id", driverHandler.Delete)
	v2.GET("/driver/:id", driverHandler.FindById)
	v2.GET("/drivers", driverHandler.FindAll)
}
