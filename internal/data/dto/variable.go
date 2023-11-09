package dto

type Variable struct {
	IsEligible     bool   `json:"is_eligible"`
	Role           string `json:"role_actor"`
	Status         string `json:"status"`
	IncidentNumber string `json:"incident_number"`
}
