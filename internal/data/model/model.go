package model

type LoginBonitaBPM struct {
	BonitaToken string `json:"bonita_token" example:"ed27cbeb-9953-4d77-b5a2-1f62a6c2e0bb"` // Bonita token uuidv4
	BonitaAuth  string `json:"bonita_auth" example:"C5385BFEE2969D9E46F0160C1952B0F1"`      // Bonita auth
}
