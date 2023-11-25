package dtopayment

import "time"

type GenerateInvoiceRequest struct {
	CaseID      int64  `json:"case_id"`
	CodeBooking string `json:"code_booking"`
}

type RequestUpdatePayment struct {
	InvoiceNumber string `json:"invoice_number"`
	Status        string `json:"status"`
	PaymentMethod string `json:"payment_method"`
}

type DoPaymentRequest struct {
	CaseID         int64     `json:"case_id"`
	InvoiceNumber  string    `json:"invoice_number"`
	PaymentMethod  string    `json:"payment_method"`
	PaymentAmount  float64   `json:"payment_amount"`
	BookingExpired time.Time `json:"booking_expired"`
}
