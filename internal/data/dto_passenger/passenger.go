package dtopassenger

type Passenger struct {
	IDNumber           string `json:"id_number"`
	VaccineStatus      string `json:"vaccine_status"`
	IsVerifiedDukcapil bool   `json:"is_verified_dukcapil"`
	CaseID             int64  `json:"case_id"`
}
