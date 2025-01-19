package dto

type MembershipRequest struct {
	MembershipName string `json:"membership_name" validate:"required"`
	Discount       int    `json:"discount" validate:"required,gte=0,lte=100"`
}

type MembershipResponse struct {
	MembershipID   uint   `json:"membership_id"`
	MembershipName string `json:"membership_name"`
	Discount       int    `json:"discount"`
}
