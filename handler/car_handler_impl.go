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

type CarHandlerImpl struct {
	CarService service.CarService
}

func NewCarHandler(carService service.CarService) CarHandler {
	return &CarHandlerImpl{
		CarService: carService,
	}
}

// Create implements CarHandler
func (handler *CarHandlerImpl) Create(c echo.Context) error {
	var req dto.CarRequest

	//bind ke struct req
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.NewErrorResponse(http.StatusBadRequest, []string{"invalid request payload"}))
	}

	//buat timeout dari parent context yang dibuat echo
	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Minute)
	defer cancel()

	response, err := handler.CarService.Create(ctx, &req)

	if err != nil {
		return helper.HandleErrorResponse(c, err)
	}

	return c.JSON(http.StatusCreated,
		helper.NewResponseWithData(http.StatusCreated, "car created sucessfully", response))
}

// Update implements CarHandler
func (handler *CarHandlerImpl) Update(c echo.Context) error {

	var req dto.CarRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.NewErrorResponse(http.StatusBadRequest, []string{"invalid request payload"}))
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.NewErrorResponse(http.StatusBadRequest, []string{"invalid car id"}))
	}

	//buat timeout dari parent context yang dibuat echo
	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Minute)
	defer cancel()

	if err := handler.CarService.Update(ctx, &req, uint(id)); err != nil {
		return helper.HandleErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK,
		helper.NewResponse(http.StatusOK, "car update sucessfully"))
}

// Delete implements CarHandler
func (handler *CarHandlerImpl) Delete(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.NewErrorResponse(http.StatusBadRequest, []string{"invalid car id"}))
	}

	//buat timeout dari parent context yang dibuat echo
	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Minute)
	defer cancel()

	if err := handler.CarService.Delete(ctx, uint(id)); err != nil {
		return helper.HandleErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK,
		helper.NewResponse(http.StatusOK, "car deleted sucessfully"))
}

// FindById implements CarHandler
func (handler *CarHandlerImpl) FindById(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.NewErrorResponse(http.StatusBadRequest, []string{"invalid car id"}))
	}

	//buat timeout dari parent context yang dibuat echo
	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Minute)
	defer cancel()

	response, err := handler.CarService.FindById(ctx, uint(id))

	if err != nil {
		return helper.HandleErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK,
		helper.NewResponseWithData(http.StatusOK, "find car succesfully", response))
}

// FindAll implements CarHandler
func (handler *CarHandlerImpl) FindAll(c echo.Context) error {

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

	responses, totalCount, totalPages, err := handler.CarService.FindAll(ctx, limit, offset)

	if err != nil {
		return helper.HandleErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK,
		helper.NewResponseWithDatas(http.StatusOK, "find cars succesfully", responses, totalPages, totalCount, limit, offset))
}
