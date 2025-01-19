package models

type Membership struct {
	MembershipID   uint       `gorm:"type:serial;primaryKey;autoIncrement;not null" json:"membership_id"`
	MembershipName string     `gorm:"not null" json:"membership_name"`
	Discount       int        `gorm:"type:int;not null" json:"discount"`
	Customer       []Customer `gorm:"foreignKey:MembershipID"`
}

func (Membership) TableName() string {
	return "memberships"
}
