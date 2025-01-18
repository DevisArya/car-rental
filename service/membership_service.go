package service

import (
	"context"

	"github.com/DevisArya/car-rental/dto"
)

type MembershipService interface {
	Create(ctx context.Context, request *dto.MembershipRequest) (*dto.MembershipResponse, error)
	Update(ctx context.Context, request *dto.MembershipRequest, membershipId uint) error
	Delete(ctx context.Context, membershipId uint) error
	FindById(ctx context.Context, membershipId uint) (*dto.MembershipResponse, error)
	FindAll(ctx context.Context, limit, offset int) (*[]dto.MembershipResponse, int, int, error)
}
