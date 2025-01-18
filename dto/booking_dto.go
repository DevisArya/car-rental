package dto

import "time"

type BookingRequest struct {
	CustomerID uint64    `json:"customer_id" validate:"required"`
	CarID      uint64    `json:"car_id" validate:"required"`
	StartDate  time.Time `json:"start_date" validate:"required"`
	EndDate    time.Time `json:"end_date" validate:"required"`
}

type BookingResponse struct {
	BookingID  uint64    `json:"booking_id"`
	CustomerID uint64    `json:"customer_id"`
	CarID      uint64    `json:"car_id"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	TotalCost  uint64    `json:"total_cost"`
	Finished   bool      `json:"finished"`
}
