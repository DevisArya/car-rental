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

type DriverServiceImpl struct {
	DriverRepository repository.DriverRepository
	DB               *gorm.DB
	validate         *validator.Validate
}

func NewDriverService(driverRepository repository.DriverRepository, DB *gorm.DB, validate *validator.Validate) DriverService {

	return &DriverServiceImpl{
		DriverRepository: driverRepository,
		DB:               DB,
		validate:         validate,
	}
}

// Create implements DriverService
func (service *DriverServiceImpl) Create(ctx context.Context, request *dto.DriverRequest) (*dto.DriverResponse, error) {

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
	custData := models.Driver{
		Name:        request.Name,
		Nik:         request.Nik,
		PhoneNumber: request.PhoneNumber,
		DailyCost:   request.DailyCost,
	}

	noRecord, err := service.DriverRepository.FindByNikAndPhoneNumber(ctx, tx, custData.PhoneNumber, custData.Nik)
	if err != nil {
		return nil, err
	}

	if !noRecord {
		return nil, helper.NewValidationError(http.StatusBadRequest, []string{"nik or phone number already used"})
	}

	driver, err := service.DriverRepository.Create(ctx, tx, &custData)
	if err != nil {
		return nil, err
	}

	return helper.ToDriverResponse(driver), nil
}

// Update implements DriverService
func (service *DriverServiceImpl) Update(ctx context.Context, request *dto.DriverRequest, driverId uint) error {
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
	custData := models.Driver{
		DriverID:    driverId,
		Name:        request.Name,
		Nik:         request.Nik,
		PhoneNumber: request.PhoneNumber,
		DailyCost:   request.DailyCost,
	}

	// cek apakah driver dengan id ini ada
	driver, err := service.DriverRepository.FindById(ctx, tx, custData.DriverID)
	if err != nil {
		return err
	}

	//jika nik tidak sama dengan sebelumnya
	if driver.Nik != request.Nik {
		custData.Nik = request.Nik

		//cek apakah tidak ada nik
		noRecord, err := service.DriverRepository.FindByNik(ctx, tx, request.Nik)
		if err != nil {
			return err
		}

		//jika ada
		if !noRecord {
			return helper.NewValidationError(http.StatusBadRequest, []string{"nik already used"})
		}
	}

	//jika nomor hp tidak sama dengan sebelumnya
	if driver.PhoneNumber != request.PhoneNumber {
		custData.PhoneNumber = request.PhoneNumber

		//cek apakah tidak ada nomor hp
		noRecord, err := service.DriverRepository.FindByPhoneNumber(ctx, tx, request.PhoneNumber)
		if err != nil {
			return err
		}

		//jika ada
		if !noRecord {
			return helper.NewValidationError(http.StatusBadRequest, []string{"phone number already used"})
		}
	}

	if err := service.DriverRepository.Update(ctx, tx, &custData); err != nil {
		return err
	}

	return nil
}

// FindById implements DriverService
func (service *DriverServiceImpl) FindById(ctx context.Context, driverId uint) (*dto.DriverResponse, error) {

	//validasi input
	if driverId <= 0 {
		return nil, errors.New("invalid driver id")
	}

	driver, err := service.DriverRepository.FindById(ctx, service.DB, driverId)
	if err != nil {
		return nil, err
	}

	return helper.ToDriverResponse(driver), nil
}

// FindAll implements DriverService
func (service *DriverServiceImpl) FindAll(ctx context.Context, limit, offset int) (*[]dto.DriverResponse, int, int, error) {

	drivers, totalCount, err := service.DriverRepository.FindAll(ctx, service.DB, limit, offset)
	if err != nil {
		return nil, 0, 0, err
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))

	return helper.ToDriverResponses(drivers), int(totalCount), totalPages, nil
}

// Delete implements DriverService
func (service *DriverServiceImpl) Delete(ctx context.Context, driverId uint) error {

	//validasi input
	if driverId <= 0 {
		return errors.New("invalid driver id")
	}
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	if _, err := service.DriverRepository.FindById(ctx, tx, driverId); err != nil {
		return err
	}

	if err := service.DriverRepository.Delete(ctx, tx, driverId); err != nil {
		return err
	}

	return nil
}
