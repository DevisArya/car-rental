package models

type Membership struct {
	MembershipID   uint    `gorm:"type:serial;primaryKey;autoIncrement;not null" json:"membership_id"`
	MembershipName string  `gorm:"not null" json:"membership_name"`
	Discount       float64 `gorm:"type:decimal(5,2);not null" json:"discount"`
}

func (Membership) TableName() string {
	return "memberships"
}
