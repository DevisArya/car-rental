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

type CustomerHandlerImpl struct {
	CustomerService service.CustomerService
}

func NewCustomerHandler(customerService service.CustomerService) CustomerHandler {
	return &CustomerHandlerImpl{
		CustomerService: customerService,
	}
}

// Create implements CustomerHandler
func (handler *CustomerHandlerImpl) Create(c echo.Context) error {
	var req dto.CustomerRequest

	//bind ke struct req
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.NewErrorResponse(http.StatusBadRequest, []string{"invalid request payload"}))
	}

	//buat timeout dari parent context yang dibuat echo
	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Minute)
	defer cancel()

	response, err := handler.CustomerService.Create(ctx, &req)

	if err != nil {
		return helper.HandleErrorResponse(c, err)
	}

	return c.JSON(http.StatusCreated,
		helper.NewResponseWithData(http.StatusCreated, "customer created sucessfully", response))
}

// Update implements CustomerHandler
func (handler *CustomerHandlerImpl) Update(c echo.Context) error {

	var req dto.CustomerRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.NewErrorResponse(http.StatusBadRequest, []string{"invalid request payload"}))
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.NewErrorResponse(http.StatusBadRequest, []string{"invalid customer id"}))
	}

	//buat timeout dari parent context yang dibuat echo
	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Minute)
	defer cancel()

	if err := handler.CustomerService.Update(ctx, &req, uint(id)); err != nil {
		return helper.HandleErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK,
		helper.NewResponse(http.StatusOK, "customer update sucessfully"))
}

// Delete implements CustomerHandler
func (handler *CustomerHandlerImpl) Delete(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.NewErrorResponse(http.StatusBadRequest, []string{"invalid customer id"}))
	}

	//buat timeout dari parent context yang dibuat echo
	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Minute)
	defer cancel()

	if err := handler.CustomerService.Delete(ctx, uint(id)); err != nil {
		return helper.HandleErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK,
		helper.NewResponse(http.StatusOK, "customer deleted sucessfully"))
}

// FindById implements CustomerHandler
func (handler *CustomerHandlerImpl) FindById(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.NewErrorResponse(http.StatusBadRequest, []string{"invalid customer id"}))
	}

	//buat timeout dari parent context yang dibuat echo
	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Minute)
	defer cancel()

	response, err := handler.CustomerService.FindById(ctx, uint(id))

	if err != nil {
		return helper.HandleErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK,
		helper.NewResponseWithData(http.StatusOK, "find customer succesfully", response))
}

// FindAll implements CustomerHandler
func (handler *CustomerHandlerImpl) FindAll(c echo.Context) error {

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

	responses, totalCount, totalPages, err := handler.CustomerService.FindAll(ctx, limit, offset)

	if err != nil {
		return helper.HandleErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK,
		helper.NewResponseWithDatas(http.StatusOK, "find customers succesfully", responses, totalPages, totalCount, limit, offset))
}
