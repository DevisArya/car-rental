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

func ToCarResponse(car *models.Car) *dto.CarResponse {
	return &dto.CarResponse{
		CarID:     car.CarID,
		Name:      car.Name,
		Stock:     car.Stock,
		DailyRent: car.DailyRent,
	}
}

func ToCarResponses(cars *[]models.Car) *[]dto.CarResponse {

	var carResponses []dto.CarResponse

	for _, val := range *cars {
		carResponses = append(carResponses, *ToCarResponse(&val))
	}
	return &carResponses
}
