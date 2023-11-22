package dtonotification

type SendEmailRequest struct {
	Route        string   `json:"route" validate:"required"`
	EmailAddress string   `json:"email_address" validate:"email"`
	To           string   `json:"to" validate:"required,email"`
	Cc           string   `json:"cc" validate:"email"`
	Bcc          string   `json:"bcc" validate:"email"`
	Subject      string   `json:"subject" validate:"required"`
	Body         string   `json:"body" validate:"required"`
	Attachments  []string `json:"attachments" validate:"required"`
}

type Request struct {
	CodeBooking string `json:"code_booking"`
	Route       string `json:"route"`
}
