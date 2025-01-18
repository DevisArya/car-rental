package models

type Driver struct {
	DriverID    uint   `gorm:"type:serial;primaryKey;autoIncrement;not null" json:"driver_id"`
	Name        string `gorm:"type:varchar(100);not null" json:"name"`
	Nik         string `gorm:"type:varchar(255);uniqueIndex;not null" json:"nik"`
	PhoneNumber string `gorm:"type:varchar(13);uniqueIndex;not null" json:"phone_number"`
	DailyCost   int    `gorm:"not null"`
}

// nama tabel di PostgreeSQL
func (Driver) TableName() string {
	return "drivers"
}
