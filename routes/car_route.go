package routes

import (
	"github.com/DevisArya/car-rental/handler"
	"github.com/labstack/echo/v4"
)

func RegisterCarRoutes(e *echo.Echo, carHandler handler.CarHandler) {

	// route for car
	v1 := e.Group("/v1")

	v1.POST("/car", carHandler.Create)
	v1.PATCH("/car/:id", carHandler.Update)
	v1.DELETE("/car/:id", carHandler.Delete)
	v1.GET("/car/:id", carHandler.FindById)
	v1.GET("/cars", carHandler.FindAll)
}
