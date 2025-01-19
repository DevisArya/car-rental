package dto

import "github.com/DevisArya/car-rental/models"

type CustomerRequest struct {
	Name         string `json:"name" validate:"required,min=3,max=100"`
	Nik          string `json:"nik" validate:"required"`
	PhoneNumber  string `json:"phone_number" validate:"required,min=10,max=13,numeric"`
	MembershipID *uint  `json:"membership_id"`
}

type CustomerResponse struct {
	CustomerID   uint   `json:"customer_id"`
	MembershipID *uint  `json:"membership_id"`
	Name         string `json:"name"`
	Nik          string `json:"nik"`
	PhoneNumber  string `json:"phone_number"`
	Membership   *models.Membership
}
