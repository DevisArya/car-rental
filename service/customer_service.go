package service

import (
	"context"

	"github.com/DevisArya/car-rental/dto"
)

type CustomerService interface {
	Create(ctx context.Context, request *dto.CustomerRequest) (*dto.CustomerResponse, error)
	Update(ctx context.Context, request *dto.CustomerRequest, customerId uint) error
	Delete(ctx context.Context, customerId uint) error
	FindById(ctx context.Context, customerId uint) (*dto.CustomerResponse, error)
	FindAll(ctx context.Context, limit, offset int) (*[]dto.CustomerResponse, int, int, error)
}
