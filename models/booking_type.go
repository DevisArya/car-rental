package models

type BookingType struct {
	BookingTypeID uint64    `gorm:"primaryKey;autoIncrement;not null" json:"booking_type_id"`
	BookingType   string    `gorm:"type:varchar(100);not null" json:"booking_type"`
	Description   string    `gorm:"type:varchar(255);not null" json:"description"`
	Bookings      []Booking `gorm:"foreignKey:BookingTypeID"`
}

// nama tabel di PostgreeSQL
func (BookingType) TableName() string {
	return "booking_types"
}
