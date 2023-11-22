package dtoticket

type RequestRedeemTicketToBPM struct {
	CaseID         int64               `json:"case_id"`
	PassengersInfo []PassengerInfoData `json:"passengers_info"`
	CodeTicket     string              `json:"code_ticket"`
}

type PassengerInfoData struct {
	IDNumber        string `json:"id_number"`
	VaccineStatus   string `json:"vaccine_status"`
	BookingDetailID int64  `json:"booking_detail_id"`
}

type RequestUpdateTicket struct {
	CodeTicket         string `json:"code_ticket"`
	BookingDetailID    int64  `json:"booking_detail_id"`
	IsEligibleToFlight bool   `json:"is_eligible_to_flight"`
}
