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

type MembershipServiceImpl struct {
	MembershipRepository repository.MembershipRepository
	DB                   *gorm.DB
	validate             *validator.Validate
}

func NewMembershipService(membershipRepository repository.MembershipRepository, DB *gorm.DB, validate *validator.Validate) MembershipService {

	return &MembershipServiceImpl{
		MembershipRepository: membershipRepository,
		DB:                   DB,
		validate:             validate,
	}
}

// Create implements MembershipService
func (service *MembershipServiceImpl) Create(ctx context.Context, request *dto.MembershipRequest) (*dto.MembershipResponse, error) {

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
	membershipData := models.Membership{
		MembershipName: request.MembershipName,
		Discount:       request.Discount,
	}

	membership, err := service.MembershipRepository.Create(ctx, tx, &membershipData)
	if err != nil {
		return nil, err
	}

	return helper.ToMembershipResponse(membership), nil
}

// Update implements MembershipService
func (service *MembershipServiceImpl) Update(ctx context.Context, request *dto.MembershipRequest, membershipId uint) error {
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
	membershipData := models.Membership{
		MembershipID:   membershipId,
		MembershipName: request.MembershipName,
		Discount:       request.Discount,
	}
	// cek apakah membership dengan id ini ada
	_, err := service.MembershipRepository.FindById(ctx, tx, membershipData.MembershipID)
	if err != nil {
		return err
	}

	if err := service.MembershipRepository.Update(ctx, tx, &membershipData); err != nil {
		return err
	}

	return nil
}

// FindById implements MembershipService
func (service *MembershipServiceImpl) FindById(ctx context.Context, membershipId uint) (*dto.MembershipResponse, error) {

	//validasi input
	if membershipId <= 0 {
		return nil, errors.New("invalid membership id")
	}

	membership, err := service.MembershipRepository.FindById(ctx, service.DB, membershipId)
	if err != nil {
		return nil, err
	}

	return helper.ToMembershipResponse(membership), nil
}

// FindAll implements MembershipService
func (service *MembershipServiceImpl) FindAll(ctx context.Context, limit, offset int) (*[]dto.MembershipResponse, int, int, error) {

	memberships, totalCount, err := service.MembershipRepository.FindAll(ctx, service.DB, limit, offset)
	if err != nil {
		return nil, 0, 0, err
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))

	return helper.ToMembershipResponses(memberships), int(totalCount), totalPages, nil
}

// Delete implements MembershipService
func (service *MembershipServiceImpl) Delete(ctx context.Context, membershipId uint) error {

	//validasi input
	if membershipId <= 0 {
		return errors.New("invalid membership id")
	}
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	if _, err := service.MembershipRepository.FindById(ctx, tx, membershipId); err != nil {
		return err
	}

	if err := service.MembershipRepository.Delete(ctx, tx, membershipId); err != nil {
		return err
	}

	return nil
}
