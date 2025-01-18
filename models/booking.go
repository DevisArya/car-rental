package models

import "time"

type Booking struct {
	BookingID  uint64    `gorm:"type:serial;primaryKey;autoIncrement;not null" json:"booking_id"`
	CustomerID uint64    `gorm:"type:bigint;not null" json:"customer_id"`
	CarID      uint64    `gorm:"type:bigint;not null" json:"car_id"`
	StartDate  time.Time `gorm:"type:date;not null" json:"start_date"`
	EndDate    time.Time `gorm:"type:date;not null" json:"end_date"`
	TotalCost  uint64    `gorm:"type:bigint" json:"total_cost"`
	Finished   bool      `gorm:"default:false" json:"finished"`
}

// nama tabel di PostgreeSQL
func (Booking) TableName() string {
	return "bookings"
}
