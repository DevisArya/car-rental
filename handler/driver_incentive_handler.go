package handler

import "github.com/labstack/echo/v4"

type DriverIncentiveHandler interface {
	FindById(c echo.Context) error
	FindAll(c echo.Context) error
}
