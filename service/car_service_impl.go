package service

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net/http"

	"github.com/DevisArya/car-rental/dto"
	"github.com/DevisArya/car-rental/helper"
	"github.com/DevisArya/car-rental/models"
	"github.com/DevisArya/car-rental/repository"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CarServiceImpl struct {
	CarRepository repository.CarRepository
	DB            *gorm.DB
	validate      *validator.Validate
}

func NewCarService(carRepository repository.CarRepository, DB *gorm.DB, validate *validator.Validate) CarService {

	return &CarServiceImpl{
		CarRepository: carRepository,
		DB:            DB,
		validate:      validate,
	}
}

// Create implements CarService
func (service *CarServiceImpl) Create(ctx context.Context, request *dto.CarRequest) (*dto.CarResponse, error) {

	//validasi input
	if err := service.validate.Struct(request); err != nil {

		var validationMessages []string
		for _, e := range err.(validator.ValidationErrors) {
			validationMessages = append(validationMessages, fmt.Sprintf("%s: %s", e.Field(), e.Tag()))
		}

		return nil, helper.NewValidationError(http.StatusBadRequest, validationMessages)
	}

	//melakukan database transaksional
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	//menyiapkan data untuk disimpan
	custData := models.Car{
		Name:      request.Name,
		Stock:     request.Stock,
		DailyRent: request.DailyRent,
	}

	car, err := service.CarRepository.Create(ctx, tx, &custData)
	if err != nil {
		return nil, err
	}

	return helper.ToCarResponse(car), nil
}

// Update implements CarService
func (service *CarServiceImpl) Update(ctx context.Context, request *dto.CarRequest, carId uint) error {
	//validasi input
	if err := service.validate.Struct(request); err != nil {

		var validationMessages []string
		for _, e := range err.(validator.ValidationErrors) {
			validationMessages = append(validationMessages, fmt.Sprintf("%s: %s", e.Field(), e.Tag()))
		}

		return helper.NewValidationError(http.StatusBadRequest, validationMessages)
	}

	//melakukan database transaksional
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	//menyiapkan data untuk update
	custData := models.Car{
		CarID:     carId,
		Name:      request.Name,
		Stock:     request.Stock,
		DailyRent: request.DailyRent,
	}

	// cek apakah car dengan id ini ada
	_, err := service.CarRepository.FindById(ctx, tx, custData.CarID)
	if err != nil {
		return err
	}

	if err := service.CarRepository.Update(ctx, tx, &custData); err != nil {
		return err
	}

	return nil
}

// FindById implements CarService
func (service *CarServiceImpl) FindById(ctx context.Context, carId uint) (*dto.CarResponse, error) {

	//validasi input
	if carId <= 0 {
		return nil, errors.New("invalid car id")
	}

	car, err := service.CarRepository.FindById(ctx, service.DB, carId)
	if err != nil {
		return nil, err
	}

	return helper.ToCarResponse(car), nil
}

// FindAll implements CarService
func (service *CarServiceImpl) FindAll(ctx context.Context, limit, offset int) (*[]dto.CarResponse, int, int, error) {

	cars, totalCount, err := service.CarRepository.FindAll(ctx, service.DB, limit, offset)
	if err != nil {
		return nil, 0, 0, err
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))

	return helper.ToCarResponses(cars), int(totalCount), totalPages, nil
}

// Delete implements CarService
func (service *CarServiceImpl) Delete(ctx context.Context, carId uint) error {

	//validasi input
	if carId <= 0 {
		return errors.New("invalid car id")
	}
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	if _, err := service.CarRepository.FindById(ctx, tx, carId); err != nil {
		return err
	}

	if err := service.CarRepository.Delete(ctx, tx, carId); err != nil {
		return err
	}

	return nil
}
