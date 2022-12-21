package requests

type ProductionPlant struct {
	PartnerFunction string  `json:"PartnerFunction"`
	DefaultPlant    *bool   `json:"DefaultPlant"`
	BusinessPartner int     `json:"BusinessPartner"`
	Plant           *string `json:"Plant"`
}
