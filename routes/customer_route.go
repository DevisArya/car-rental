package routes

import (
	"github.com/DevisArya/car-rental/handler"
	"github.com/labstack/echo/v4"
)

func RegisterCustomerRoutes(e *echo.Echo, customerHandler handler.CustomerHandler) {

	// route for customer
	v2 := e.Group("/v2")

	v2.POST("/customer", customerHandler.Create)
	v2.PATCH("/customer/:id", customerHandler.Update)
	v2.DELETE("/customer/:id", customerHandler.Delete)
	v2.GET("/customer/:id", customerHandler.FindById)
	v2.GET("/customers", customerHandler.FindAll)
}
