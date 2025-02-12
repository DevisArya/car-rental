package dto

type DriverRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Nik         string `json:"nik" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required,min=10,max=13,numeric"`
	DailyCost   int    `json:"daily_cost" validate:"required,min=1"`
}

type DriverResponse struct {
	DriverID    uint   `json:"driver_id"`
	Name        string `json:"name"`
	Nik         string `json:"nik"`
	PhoneNumber string `json:"phone_number"`
	DailyCost   int    `json:"daily_cost"`
}
