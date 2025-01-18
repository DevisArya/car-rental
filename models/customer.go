package models

type Customer struct {
	CustomerID  uint   `gorm:"type:serial;primaryKey;autoIncrement;not null" json:"customer_id"`
	Name        string `gorm:"type:varchar(100);not null" json:"name"`
	Nik         string `gorm:"type:varchar(255);uniqueIndex;not null" json:"nik"`
	PhoneNumber string `gorm:"type:varchar(13);uniqueIndex;not null" json:"phone_number"`
}

// nama tabel di PostgreeSQL
func (Customer) TableName() string {
	return "customers"
}
