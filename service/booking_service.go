package service

import (
	"context"

	"github.com/DevisArya/car-rental/dto"
)

type BookingService interface {
	Create(ctx context.Context, request *dto.BookingRequest) (*dto.BookingResponse, error)
	Update(ctx context.Context, request *dto.BookingRequest, bookingId uint) error
	UpdateStatus(ctx context.Context, bookingId uint) error
	Delete(ctx context.Context, bookingId uint) error
	FindById(ctx context.Context, bookingId uint) (*dto.BookingResponse, error)
	FindAll(ctx context.Context, limit, offset int) (*[]dto.BookingResponse, int, int, error)
}
