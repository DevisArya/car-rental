package helper

import (
	"github.com/DevisArya/car-rental/dto"
	"github.com/DevisArya/car-rental/models"
)

func ToCustomerResponse(customer *models.Customer) *dto.CustomerResponse {
	return &dto.CustomerResponse{
		CustomerID:   customer.CustomerID,
		Name:         customer.Name,
		Nik:          customer.Nik,
		PhoneNumber:  customer.PhoneNumber,
		MembershipID: customer.MembershipID,
		Membership:   customer.Membership,
	}
}

func ToCustomerResponses(customers *[]models.Customer) *[]dto.CustomerResponse {

	var customerResponses []dto.CustomerResponse

	for _, val := range *customers {
		customerResponses = append(customerResponses, *ToCustomerResponse(&val))
	}
	return &customerResponses
}

func ToCarResponse(car *models.Car) *dto.CarResponse {
	return &dto.CarResponse{
		CarID:     car.CarID,
		Name:      car.Name,
		Stock:     car.Stock,
		DailyRent: car.DailyRent,
	}
}

func ToCarResponses(cars *[]models.Car) *[]dto.CarResponse {

	var carResponses []dto.CarResponse

	for _, val := range *cars {
		carResponses = append(carResponses, *ToCarResponse(&val))
	}
	return &carResponses
}

func ToBookingResponse(booking *models.Booking) *dto.BookingResponse {
	return &dto.BookingResponse{
		BookingID:       booking.BookingID,
		CustomerID:      booking.CustomerID,
		CarID:           booking.CarID,
		DriverID:        booking.DriverID,
		BookingTypeID:   booking.BookingTypeID,
		StartDate:       booking.StartDate,
		EndDate:         booking.EndDate,
		TotalCost:       booking.TotalCost,
		Finished:        booking.Finished,
		Discount:        booking.Discount,
		TotalDriverCost: booking.TotalDriverCost,
	}
}

func ToBookingResponses(bookings *[]models.Booking) *[]dto.BookingResponse {

	var bookingResponses []dto.BookingResponse

	for _, val := range *bookings {
		bookingResponses = append(bookingResponses, *ToBookingResponse(&val))
	}
	return &bookingResponses
}

func ToMembershipResponse(membership *models.Membership) *dto.MembershipResponse {
	return &dto.MembershipResponse{
		MembershipID:   membership.MembershipID,
		MembershipName: membership.MembershipName,
		Discount:       membership.Discount,
	}
}

func ToMembershipResponses(memberships *[]models.Membership) *[]dto.MembershipResponse {

	var membershipResponses []dto.MembershipResponse

	for _, val := range *memberships {
		membershipResponses = append(membershipResponses, *ToMembershipResponse(&val))
	}
	return &membershipResponses
}

func ToDriverResponse(driver *models.Driver) *dto.DriverResponse {
	return &dto.DriverResponse{
		DriverID:    driver.DriverID,
		Name:        driver.Name,
		Nik:         driver.Nik,
		PhoneNumber: driver.PhoneNumber,
		DailyCost:   driver.DailyCost,
	}
}

func ToDriverResponses(drivers *[]models.Driver) *[]dto.DriverResponse {

	var driverResponses []dto.DriverResponse

	for _, val := range *drivers {
		driverResponses = append(driverResponses, *ToDriverResponse(&val))
	}
	return &driverResponses
}

func ToDriverIncentiveResponse(driverIncentive *models.DriverIncentive) *dto.DriverIncentiveResponse {
	return &dto.DriverIncentiveResponse{
		DriverIncentiveID: driverIncentive.DriverIncentiveID,
		BookingID:         driverIncentive.BookingID,
		Incentive:         driverIncentive.Incentive,
	}
}

func ToDriverIncentiveResponses(driverIncentives *[]models.DriverIncentive) *[]dto.DriverIncentiveResponse {

	var driverIncentiveResponses []dto.DriverIncentiveResponse

	for _, val := range *driverIncentives {
		driverIncentiveResponses = append(driverIncentiveResponses, *ToDriverIncentiveResponse(&val))
	}
	return &driverIncentiveResponses
}
