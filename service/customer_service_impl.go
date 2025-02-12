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

type CustomerServiceImpl struct {
	CustomerRepository repository.CustomerRepository
	MembershipService  MembershipService
	DB                 *gorm.DB
	validate           *validator.Validate
}

func NewCustomerService(
	customerRepository repository.CustomerRepository,
	membershipService MembershipService,
	DB *gorm.DB,
	validate *validator.Validate) CustomerService {

	return &CustomerServiceImpl{
		CustomerRepository: customerRepository,
		MembershipService:  membershipService,
		DB:                 DB,
		validate:           validate,
	}
}

// Create implements CustomerService
func (service *CustomerServiceImpl) Create(ctx context.Context, request *dto.CustomerRequest) (*dto.CustomerResponse, error) {

	//validasi input
	if err := service.validate.Struct(request); err != nil {

		var validationMessages []string
		for _, e := range err.(validator.ValidationErrors) {
			validationMessages = append(validationMessages, fmt.Sprintf("%s: %s", e.Field(), e.Tag()))
		}

		return nil, helper.NewValidationError(http.StatusBadRequest, validationMessages)
	}

	//cek apakah id membership ada
	if request.MembershipID != nil {
		_, err := service.MembershipService.FindById(ctx, *request.MembershipID)
		if err != nil {
			return nil, err
		}
	}

	//melakukan database transaksional
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	//menyiapkan data untuk disimpan
	custData := models.Customer{
		Name:         request.Name,
		Nik:          request.Nik,
		PhoneNumber:  request.PhoneNumber,
		MembershipID: request.MembershipID,
	}

	noRecord, err := service.CustomerRepository.FindByNikAndPhoneNumber(ctx, tx, custData.PhoneNumber, custData.Nik)
	if err != nil {
		return nil, err
	}

	if !noRecord {
		return nil, helper.NewValidationError(http.StatusBadRequest, []string{"nik or phone number already used"})
	}

	customer, err := service.CustomerRepository.Create(ctx, tx, &custData)
	if err != nil {
		return nil, err
	}

	return helper.ToCustomerResponse(customer), nil
}

// Update implements CustomerService
func (service *CustomerServiceImpl) Update(ctx context.Context, request *dto.CustomerRequest, customerId uint) error {
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
	custData := models.Customer{
		CustomerID:   customerId,
		Name:         request.Name,
		PhoneNumber:  request.PhoneNumber,
		MembershipID: request.MembershipID,
	}

	// cek apakah customer dengan id ini ada
	customer, err := service.CustomerRepository.FindById(ctx, tx, custData.CustomerID)
	if err != nil {
		return err
	}

	//jika nik tidak sama dengan sebelumnya
	if customer.Nik != request.Nik {

		custData.Nik = request.Nik

		//cek apakah tidak ada nik
		noRecord, err := service.CustomerRepository.FindByNik(ctx, tx, request.Nik)
		if err != nil {
			return err
		}

		fmt.Println(noRecord)

		//jika ada
		if !noRecord {
			return helper.NewValidationError(http.StatusBadRequest, []string{"nik already used"})
		}
	}

	//jika nomor hp tidak sama dengan sebelumnya
	if customer.PhoneNumber != request.PhoneNumber {
		custData.PhoneNumber = request.PhoneNumber

		//cek apakah tidak ada nomor hp
		noRecord, err := service.CustomerRepository.FindByPhoneNumber(ctx, tx, request.PhoneNumber)
		if err != nil {
			return err
		}

		//jika ada
		if !noRecord {
			return helper.NewValidationError(http.StatusBadRequest, []string{"phone number already used"})
		}
	}

	if err := service.CustomerRepository.Update(ctx, tx, &custData); err != nil {
		return err
	}

	return nil
}

// FindById implements CustomerService
func (service *CustomerServiceImpl) FindById(ctx context.Context, customerId uint) (*dto.CustomerResponse, error) {

	//validasi input
	if customerId <= 0 {
		return nil, errors.New("invalid customer id")
	}

	customer, err := service.CustomerRepository.FindById(ctx, service.DB, customerId)
	if err != nil {
		return nil, err
	}

	return helper.ToCustomerResponse(customer), nil
}

// FindAll implements CustomerService
func (service *CustomerServiceImpl) FindAll(ctx context.Context, limit, offset int) (*[]dto.CustomerResponse, int, int, error) {

	customers, totalCount, err := service.CustomerRepository.FindAll(ctx, service.DB, limit, offset)
	if err != nil {
		return nil, 0, 0, err
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))

	return helper.ToCustomerResponses(customers), int(totalCount), totalPages, nil
}

// Delete implements CustomerService
func (service *CustomerServiceImpl) Delete(ctx context.Context, customerId uint) error {

	//validasi input
	if customerId <= 0 {
		return errors.New("invalid customer id")
	}
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	if _, err := service.CustomerRepository.FindById(ctx, tx, customerId); err != nil {
		return err
	}

	if err := service.CustomerRepository.Delete(ctx, tx, customerId); err != nil {
		return err
	}

	return nil
}
