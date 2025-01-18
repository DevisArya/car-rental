package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/DevisArya/car-rental/dto"
	"github.com/DevisArya/car-rental/helper"
	"github.com/DevisArya/car-rental/service"
	"github.com/labstack/echo/v4"
)

type DriverHandlerImpl struct {
	DriverService service.DriverService
}

func NewDriverHandler(driverService service.DriverService) DriverHandler {
	return &DriverHandlerImpl{
		DriverService: driverService,
	}
}

// Create implements DriverHandler
func (handler *DriverHandlerImpl) Create(c echo.Context) error {
	var req dto.DriverRequest

	//bind ke struct req
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.NewErrorResponse(http.StatusBadRequest, []string{"invalid request payload"}))
	}

	//buat timeout dari parent context yang dibuat echo
	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Minute)
	defer cancel()

	response, err := handler.DriverService.Create(ctx, &req)

	if err != nil {
		return helper.HandleErrorResponse(c, err)
	}

	return c.JSON(http.StatusCreated,
		helper.NewResponseWithData(http.StatusCreated, "driver created sucessfully", response))
}

// Update implements DriverHandler
func (handler *DriverHandlerImpl) Update(c echo.Context) error {

	var req dto.DriverRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.NewErrorResponse(http.StatusBadRequest, []string{"invalid request payload"}))
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.NewErrorResponse(http.StatusBadRequest, []string{"invalid driver id"}))
	}

	//buat timeout dari parent context yang dibuat echo
	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Minute)
	defer cancel()

	if err := handler.DriverService.Update(ctx, &req, uint(id)); err != nil {
		return helper.HandleErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK,
		helper.NewResponse(http.StatusOK, "driver update sucessfully"))
}

// Delete implements DriverHandler
func (handler *DriverHandlerImpl) Delete(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.NewErrorResponse(http.StatusBadRequest, []string{"invalid driver id"}))
	}

	//buat timeout dari parent context yang dibuat echo
	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Minute)
	defer cancel()

	if err := handler.DriverService.Delete(ctx, uint(id)); err != nil {
		return helper.HandleErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK,
		helper.NewResponse(http.StatusOK, "driver deleted sucessfully"))
}

// FindById implements DriverHandler
func (handler *DriverHandlerImpl) FindById(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.NewErrorResponse(http.StatusBadRequest, []string{"invalid driver id"}))
	}

	//buat timeout dari parent context yang dibuat echo
	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Minute)
	defer cancel()

	response, err := handler.DriverService.FindById(ctx, uint(id))

	if err != nil {
		return helper.HandleErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK,
		helper.NewResponseWithData(http.StatusOK, "find driver succesfully", response))
}

// FindAll implements DriverHandler
func (handler *DriverHandlerImpl) FindAll(c echo.Context) error {

	//buat timeout dari parent context yang dibuat echo
	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Minute)
	defer cancel()

	limit := 10
	offset := 0

	// Parse query parameters
	if l := c.QueryParam("limit"); l != "" {
		if val, err := strconv.Atoi(l); err == nil {
			limit = val
		} else {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid limit"})
		}
	}
	if o := c.QueryParam("offset"); o != "" {
		if val, err := strconv.Atoi(o); err == nil {
			offset = val
		} else {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid offset"})
		}
	}

	responses, totalCount, totalPages, err := handler.DriverService.FindAll(ctx, limit, offset)

	if err != nil {
		return helper.HandleErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK,
		helper.NewResponseWithDatas(http.StatusOK, "find drivers succesfully", responses, totalPages, totalCount, limit, offset))
}
