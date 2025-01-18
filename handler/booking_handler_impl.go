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

type BookingHandlerImpl struct {
	BookingService service.BookingService
}

func NewBookingHandler(bookingService service.BookingService) BookingHandler {
	return &BookingHandlerImpl{
		BookingService: bookingService,
	}
}

// Create implements BookingHandler
func (handler *BookingHandlerImpl) Create(c echo.Context) error {
	var req dto.BookingRequest

	//bind ke struct req
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.NewErrorResponse(http.StatusBadRequest, []string{"invalid request payload"}))
	}

	//buat timeout dari parent context yang dibuat echo
	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Minute)
	defer cancel()

	response, err := handler.BookingService.Create(ctx, &req)

	if err != nil {
		return helper.HandleErrorResponse(c, err)
	}

	return c.JSON(http.StatusCreated,
		helper.NewResponseWithData(http.StatusCreated, "booking created sucessfully", response))
}

// Update implements BookingHandler
func (handler *BookingHandlerImpl) Update(c echo.Context) error {

	var req dto.BookingRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.NewErrorResponse(http.StatusBadRequest, []string{"invalid request payload"}))
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.NewErrorResponse(http.StatusBadRequest, []string{"invalid booking id"}))
	}

	//buat timeout dari parent context yang dibuat echo
	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Minute)
	defer cancel()

	if err := handler.BookingService.Update(ctx, &req, uint(id)); err != nil {
		return helper.HandleErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK,
		helper.NewResponse(http.StatusOK, "booking update sucessfully"))
}

// UpdateStatus implements BookingHandler
func (handler *BookingHandlerImpl) UpdateStatus(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.NewErrorResponse(http.StatusBadRequest, []string{"invalid booking id"}))
	}

	//buat timeout dari parent context yang dibuat echo
	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Minute)
	defer cancel()

	if err := handler.BookingService.UpdateStatus(ctx, uint(id)); err != nil {
		return helper.HandleErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK,
		helper.NewResponse(http.StatusOK, "booking status update sucessfully"))
}

// Delete implements BookingHandler
func (handler *BookingHandlerImpl) Delete(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.NewErrorResponse(http.StatusBadRequest, []string{"invalid booking id"}))
	}

	//buat timeout dari parent context yang dibuat echo
	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Minute)
	defer cancel()

	if err := handler.BookingService.Delete(ctx, uint(id)); err != nil {
		return helper.HandleErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK,
		helper.NewResponse(http.StatusOK, "booking deleted sucessfully"))
}

// FindById implements BookingHandler
func (handler *BookingHandlerImpl) FindById(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.NewErrorResponse(http.StatusBadRequest, []string{"invalid booking id"}))
	}

	//buat timeout dari parent context yang dibuat echo
	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Minute)
	defer cancel()

	response, err := handler.BookingService.FindById(ctx, uint(id))

	if err != nil {
		return helper.HandleErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK,
		helper.NewResponseWithData(http.StatusOK, "find booking succesfully", response))
}

// FindAll implements BookingHandler
func (handler *BookingHandlerImpl) FindAll(c echo.Context) error {

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

	responses, totalCount, totalPages, err := handler.BookingService.FindAll(ctx, limit, offset)

	if err != nil {
		return helper.HandleErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK,
		helper.NewResponseWithDatas(http.StatusOK, "find bookings succesfully", responses, totalPages, totalCount, limit, offset))
}
