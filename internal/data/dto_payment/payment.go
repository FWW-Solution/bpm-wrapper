package dtopayment

type GenerateInvoiceRequest struct {
	CaseID      int64  `json:"case_id"`
	CodeBooking string `json:"code_booking"`
}

type DoPaymentRequest struct {
	CaseID         int64   `json:"case_id"`
	CodePayment    string  `json:"code_payment"`
	PaymentMethod  string  `json:"payment_method"`
	PaymentAmmount float64 `json:"payment_ammount"`
}
