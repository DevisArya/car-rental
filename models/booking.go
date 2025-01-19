package models

import "time"

type Booking struct {
	BookingID       uint64    `gorm:"primaryKey;autoIncrement;not null" json:"booking_id"`
	CustomerID      uint64    `gorm:"type:bigint;not null" json:"customer_id"`
	CarID           uint64    `gorm:"type:bigint;not null" json:"car_id"`
	DriverID        *uint     `json:"driver_id"`
	BookingTypeID   uint64    `gorm:"not null" json:"booking_type_id"`
	StartDate       time.Time `gorm:"type:date;not null" json:"start_date"`
	EndDate         time.Time `gorm:"type:date;not null" json:"end_date"`
	TotalCost       uint64    `gorm:"type:bigint" json:"total_cost"`
	Finished        bool      `gorm:"default:false" json:"finished"`
	Discount        int       `gorm:"type:int" json:"discount"`
	TotalDriverCost int       `gorm:"type:int" json:"total_driver_count"`

	DriverIncentive DriverIncentive `gorm:"foreignKey:BookingID;references:BookingID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	BookingType BookingType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Customer    Customer    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Car         Car         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Driver      *Driver     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

// nama tabel di PostgreeSQL
func (Booking) TableName() string {
	return "bookings"
}
