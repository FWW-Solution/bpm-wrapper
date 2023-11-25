package dtobooking

import "time"

type StartProcessBookingRequest struct {
	CodeBooking    string    `json:"code_booking"`
	PaymentExpired time.Time `json:"payment_expired"`
}

type RequestUpdateBooking struct {
	CodeBooking string `json:"code_booking"`
	Status      string `json:"status"`
}
