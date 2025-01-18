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

type MembershipHandlerImpl struct {
	MembershipService service.MembershipService
}

func NewMembershipHandler(membershipService service.MembershipService) MembershipHandler {
	return &MembershipHandlerImpl{
		MembershipService: membershipService,
	}
}

// Create implements MembershipHandler
func (handler *MembershipHandlerImpl) Create(c echo.Context) error {
	var req dto.MembershipRequest

	//bind ke struct req
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.NewErrorResponse(http.StatusBadRequest, []string{"invalid request payload"}))
	}

	//buat timeout dari parent context yang dibuat echo
	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Minute)
	defer cancel()

	response, err := handler.MembershipService.Create(ctx, &req)

	if err != nil {
		return helper.HandleErrorResponse(c, err)
	}

	return c.JSON(http.StatusCreated,
		helper.NewResponseWithData(http.StatusCreated, "membership created sucessfully", response))
}

// Update implements MembershipHandler
func (handler *MembershipHandlerImpl) Update(c echo.Context) error {

	var req dto.MembershipRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.NewErrorResponse(http.StatusBadRequest, []string{"invalid request payload"}))
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.NewErrorResponse(http.StatusBadRequest, []string{"invalid membership id"}))
	}

	//buat timeout dari parent context yang dibuat echo
	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Minute)
	defer cancel()

	if err := handler.MembershipService.Update(ctx, &req, uint(id)); err != nil {
		return helper.HandleErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK,
		helper.NewResponse(http.StatusOK, "membership update sucessfully"))
}

// Delete implements MembershipHandler
func (handler *MembershipHandlerImpl) Delete(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.NewErrorResponse(http.StatusBadRequest, []string{"invalid membership id"}))
	}

	//buat timeout dari parent context yang dibuat echo
	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Minute)
	defer cancel()

	if err := handler.MembershipService.Delete(ctx, uint(id)); err != nil {
		return helper.HandleErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK,
		helper.NewResponse(http.StatusOK, "membership deleted sucessfully"))
}

// FindById implements MembershipHandler
func (handler *MembershipHandlerImpl) FindById(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.NewErrorResponse(http.StatusBadRequest, []string{"invalid membership id"}))
	}

	//buat timeout dari parent context yang dibuat echo
	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Minute)
	defer cancel()

	response, err := handler.MembershipService.FindById(ctx, uint(id))

	if err != nil {
		return helper.HandleErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK,
		helper.NewResponseWithData(http.StatusOK, "find membership succesfully", response))
}

// FindAll implements MembershipHandler
func (handler *MembershipHandlerImpl) FindAll(c echo.Context) error {

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

	responses, totalCount, totalPages, err := handler.MembershipService.FindAll(ctx, limit, offset)

	if err != nil {
		return helper.HandleErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK,
		helper.NewResponseWithDatas(http.StatusOK, "find memberships succesfully", responses, totalPages, totalCount, limit, offset))
}
