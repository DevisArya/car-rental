package dto

type CustomerRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Nik         string `json:"nik" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required,min=10,max=13,numeric"`
}

type CustomerResponse struct {
	CustomerID  uint   `json:"customer_id"`
	Name        string `json:"name"`
	Nik         string `json:"nik"`
	PhoneNumber string `json:"phone_number"`
}
