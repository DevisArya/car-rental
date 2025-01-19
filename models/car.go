package models

type Car struct {
	CarID     uint      `gorm:"type:serial;primaryKey;autoIncrement;not null" json:"car_id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Stock     int16     `gorm:"type:smallint;not null;default:0" json:"stock"`
	DailyRent int       `gorm:"type:int;not null" json:"daily_rent"`
	Bookings  []Booking `gorm:"foreignKey:CarID"`
}

// nama tabel di PostgreeSQL
func (Car) TableName() string {
	return "cars"
}
