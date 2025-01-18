package routes

import (
	"github.com/DevisArya/car-rental/handler"
	"github.com/labstack/echo/v4"
)

func RegisterCarRoutes(e *echo.Echo, carHandler handler.CarHandler) {

	// route for car
	v2 := e.Group("/v2")

	v2.POST("/car", carHandler.Create)
	v2.PATCH("/car/:id", carHandler.Update)
	v2.DELETE("/car/:id", carHandler.Delete)
	v2.GET("/car/:id", carHandler.FindById)
	v2.GET("/cars", carHandler.FindAll)
}
