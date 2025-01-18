package helper

import (
	"github.com/DevisArya/car-rental/dto"
	"github.com/DevisArya/car-rental/models"
)

func ToCustomerResponse(customer *models.Customer) *dto.CustomerResponse {
	return &dto.CustomerResponse{
		CustomerID:  customer.CustomerID,
		Name:        customer.Name,
		Nik:         customer.Nik,
		PhoneNumber: customer.PhoneNumber,
	}
}

func ToCustomerResponses(customers *[]models.Customer) *[]dto.CustomerResponse {

	var customerResponses []dto.CustomerResponse

	for _, val := range *customers {
		customerResponses = append(customerResponses, *ToCustomerResponse(&val))
	}
	return &customerResponses
}
