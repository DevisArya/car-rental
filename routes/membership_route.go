package routes

import (
	"github.com/DevisArya/car-rental/handler"
	"github.com/labstack/echo/v4"
)

func RegisterMembershipRoutes(e *echo.Echo, membershipHandler handler.MembershipHandler) {

	// route for membership
	v2 := e.Group("/v2")

	v2.POST("/membership", membershipHandler.Create)
	v2.PATCH("/membership/:id", membershipHandler.Update)
	v2.DELETE("/membership/:id", membershipHandler.Delete)
	v2.GET("/membership/:id", membershipHandler.FindById)
	v2.GET("/memberships", membershipHandler.FindAll)
}
