package models

type DriverIncentive struct {
	DriverIncentiveID uint   `gorm:"type:serial;primaryKey;autoIncrement;not null" json:"driver_incentive_id"`
	BookingID         uint64 `gorm:"not null;uniqueIndex" json:"booking_id"`
	Incentive         uint   `gorm:"not null" json:"incentive"`
}

// nama tabel di PostgreeSQL
func (DriverIncentive) TableName() string {
	return "driver_incentives"
}
