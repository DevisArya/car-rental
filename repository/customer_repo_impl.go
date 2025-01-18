package repository

import (
	"context"

	"github.com/DevisArya/car-rental/models"
	"gorm.io/gorm"
)

type CustomerRepositoryImpl struct {
}

func NewCustomerRepository() CustomerRepository {
	return &CustomerRepositoryImpl{}
}

// Save implements CustomerRepository
func (*CustomerRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, dataCustomer *models.Customer) (*models.Customer, error) {

	if err := tx.WithContext(ctx).
		Create(dataCustomer).
		Error; err != nil {
		return nil, err
	}

	return dataCustomer, nil
}

// Update implements CustomerRepository
func (*CustomerRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, dataCustomer *models.Customer) error {

	if err := tx.WithContext(ctx).
		Model(&models.Customer{}).
		Where("customer_id = ?", dataCustomer.CustomerID).
		Updates(dataCustomer).
		Error; err != nil {
		return err
	}

	return nil
}

// Delete implements CustomerRepository
func (*CustomerRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, customerId uint) error {

	if err := tx.WithContext(ctx).
		Delete(&models.Customer{}, customerId).
		Error; err != nil {
		return err
	}

	return nil
}

// FindById implements CustomerRepository
func (*CustomerRepositoryImpl) FindById(ctx context.Context, db *gorm.DB, customerId uint) (*models.Customer, error) {
	var customer models.Customer

	if err := db.WithContext(ctx).
		First(&customer, customerId).
		Error; err != nil {
		return nil, err
	}

	return &customer, nil
}

// FindAll implements CustomerRepository
func (*CustomerRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB, limit int, offset int) (*[]models.Customer, int64, error) {
	var customers []models.Customer
	var totalCount int64

	if err := db.WithContext(ctx).
		Model(&models.Customer{}).
		Count(&totalCount).
		Error; err != nil {
		return nil, 0, err
	}

	if err := db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&customers).
		Error; err != nil {
		return nil, 0, err
	}

	if len(customers) == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}

	return &customers, totalCount, nil
}

// FindByNikAndPhoneNumber implements CustomerRepository
func (*CustomerRepositoryImpl) FindByNikAndPhoneNumber(ctx context.Context, db *gorm.DB, PhoneNumber string, Nik string) (bool, error) {
	var customer models.Customer

	err := db.WithContext(ctx).
		Where("nik = ? OR phone_number = ?", Nik, PhoneNumber).
		First(&customer).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return true, nil
		}
		return false, err
	}

	return false, nil
}
