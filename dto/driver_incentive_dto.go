package dto

import (
	"time"
)

type FindDriverIncentiveRequest struct {
	StartDate time.Time `json:"start_date" validate:"required"`
	EndDate   time.Time `json:"end_date" validate:"required"`
}

type DriverIncentiveResponse struct {
	DriverIncentiveID uint   `json:"driver_incentive_id"`
	BookingID         uint64 `json:"booking_id"`
	Incentive         uint   `json:"incentive"`
}
