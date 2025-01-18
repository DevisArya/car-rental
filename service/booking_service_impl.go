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

type BookingServiceImpl struct {
	BookingRepository repository.BookingRepository
	CarService        CarService
	DB                *gorm.DB
	validate          *validator.Validate
}

func NewBookingService(
	bookingRepository repository.BookingRepository,
	carService CarService,
	DB *gorm.DB,
	validate *validator.Validate) BookingService {

	return &BookingServiceImpl{
		BookingRepository: bookingRepository,
		CarService:        carService,
		DB:                DB,
		validate:          validate,
	}
}

// Create implements BookingService
func (service *BookingServiceImpl) Create(ctx context.Context, request *dto.BookingRequest) (*dto.BookingResponse, error) {

	//validasi input
	if err := service.validate.Struct(request); err != nil {

		var validationMessages []string
		for _, e := range err.(validator.ValidationErrors) {
			validationMessages = append(validationMessages, fmt.Sprintf("%s: %s", e.Field(), e.Tag()))
		}

		return nil, helper.NewValidationError(http.StatusBadRequest, validationMessages)
	}

	if request.EndDate.Before(request.StartDate) {
		return nil, helper.NewValidationError(http.StatusBadRequest, []string{"end date cannot be earlier than the start date"})
	}
	//melakukan database transaksional
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	//ambil harga sewa perhari
	car, err := service.CarService.FindById(ctx, uint(request.CarID))
	if err != nil {
		return nil, err
	}

	//hitung biaya sewa perhari * jumlah hari
	difference := request.EndDate.Sub(request.StartDate)
	days := int(difference.Hours()/24) + 1

	totalCost := uint64(days) * uint64(car.DailyRent)

	//menyiapkan data untuk disimpan
	custData := models.Booking{
		CustomerID: request.CustomerID,
		CarID:      request.CarID,
		StartDate:  request.StartDate,
		EndDate:    request.EndDate,
		Finished:   false,
		TotalCost:  totalCost,
	}

	booking, err := service.BookingRepository.Create(ctx, tx, &custData)
	if err != nil {
		return nil, err
	}

	// lock and update stock
	if err := service.CarService.SelectForUpdateCarStock(ctx, tx, uint(custData.CarID), -1); err != nil {
		return nil, err
	}

	return helper.ToBookingResponse(booking), nil
}

// Update implements BookingService
func (service *BookingServiceImpl) Update(ctx context.Context, request *dto.BookingRequest, bookingId uint) error {
	//validasi input
	if err := service.validate.Struct(request); err != nil {

		var validationMessages []string
		for _, e := range err.(validator.ValidationErrors) {
			validationMessages = append(validationMessages, fmt.Sprintf("%s: %s", e.Field(), e.Tag()))
		}

		return helper.NewValidationError(http.StatusBadRequest, validationMessages)
	}

	if request.EndDate.Before(request.StartDate) {
		return helper.NewValidationError(http.StatusBadRequest, []string{"end date cannot be earlier than the start date"})
	}
	//melakukan database transaksional
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	// cek apakah booking dengan id ini apakah ada
	booking, err := service.BookingRepository.FindById(ctx, tx, uint(bookingId))
	if err != nil {
		return err
	}

	//ambil harga sewa perhari
	car, err := service.CarService.FindById(ctx, uint(request.CarID))
	if err != nil {
		return err
	}

	//hitung biaya sewa perhari * jumlah hari
	difference := request.EndDate.Sub(request.StartDate)
	days := int(difference.Hours()/24) + 1

	totalCost := uint64(days) * uint64(car.DailyRent)

	//menyiapkan data untuk disimpan
	custData := models.Booking{
		BookingID:  uint64(bookingId),
		CustomerID: request.CustomerID,
		CarID:      request.CarID,
		StartDate:  request.StartDate,
		EndDate:    request.EndDate,
		TotalCost:  totalCost,
	}

	if err := service.BookingRepository.Update(ctx, tx, &custData); err != nil {
		return err
	}

	//jika car id berbeda dengan sebelumnya maka perlu update stock
	if booking.CarID != custData.CarID {
		// lock and update stock (kurangi stock car yang baru)
		if err := service.CarService.SelectForUpdateCarStock(ctx, tx, uint(custData.CarID), -1); err != nil {
			return err
		}

		// lock and update stock (balikan stock car yang lama)
		if err := service.CarService.SelectForUpdateCarStock(ctx, tx, uint(booking.CarID), 1); err != nil {
			return err
		}
	}

	return nil
}

// Update implements BookingService
func (service *BookingServiceImpl) UpdateStatus(ctx context.Context, bookingId uint) error {

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	// cek apakah booking dengan id ini ada
	booking, err := service.BookingRepository.FindById(ctx, tx, uint(bookingId))
	if err != nil {
		return err
	}

	//menyiapkan data untuk disimpan
	custData := models.Booking{
		Finished: true,
	}

	if err := service.BookingRepository.Update(ctx, tx, &custData); err != nil {
		return err
	}

	// lock and update stock (update stock jika mobil sudah dikembalikan)
	if err := service.CarService.SelectForUpdateCarStock(ctx, tx, uint(booking.CarID), 1); err != nil {
		return err
	}

	return nil
}

// FindById implements BookingService
func (service *BookingServiceImpl) FindById(ctx context.Context, bookingId uint) (*dto.BookingResponse, error) {

	//validasi input
	if bookingId <= 0 {
		return nil, errors.New("invalid booking id")
	}

	booking, err := service.BookingRepository.FindById(ctx, service.DB, bookingId)
	if err != nil {
		return nil, err
	}

	return helper.ToBookingResponse(booking), nil
}

// FindAll implements BookingService
func (service *BookingServiceImpl) FindAll(ctx context.Context, limit, offset int) (*[]dto.BookingResponse, int, int, error) {

	bookings, totalCount, err := service.BookingRepository.FindAll(ctx, service.DB, limit, offset)
	if err != nil {
		return nil, 0, 0, err
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))

	return helper.ToBookingResponses(bookings), int(totalCount), totalPages, nil
}

// Delete implements BookingService
func (service *BookingServiceImpl) Delete(ctx context.Context, bookingId uint) error {

	//validasi input
	if bookingId <= 0 {
		return errors.New("invalid booking id")
	}
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	if _, err := service.BookingRepository.FindById(ctx, tx, bookingId); err != nil {
		return err
	}

	if err := service.BookingRepository.Delete(ctx, tx, bookingId); err != nil {
		return err
	}

	return nil
}
