package routes

import (
	"github.com/DevisArya/car-rental/handler"
	"github.com/labstack/echo/v4"
)

func RegisterCustomerRoutes(e *echo.Echo, customerHandler handler.CustomerHandler) {

	// route for customer
	v1 := e.Group("/v1")

	v1.POST("/customer", customerHandler.Create)
	v1.PATCH("/customer/:id", customerHandler.Update)
	v1.DELETE("/customer/:id", customerHandler.Delete)
	v1.GET("/customer/:id", customerHandler.FindById)
	v1.GET("/customers", customerHandler.FindAll)
}
