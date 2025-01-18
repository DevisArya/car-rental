package dto

type CarRequest struct {
	Name      string `json:"name" validate:"required,max=100"`
	Stock     int16  `json:"stock" validate:"required,min=0"`
	DailyRent int    `json:"daily_rent" validate:"required,min=1"`
}

type CarResponse struct {
	CarID     uint   `json:"car_id"`
	Name      string `json:"name"`
	Stock     int16  `json:"stock"`
	DailyRent int    `json:"daily_rent"`
}
