package requests

type HeaderPartnerBPGeneral struct {
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     string  `json:"BusinessPartnerName"`
	Country                 string  `json:"Country"`
	Language                string  `json:"Language"`
	AddressID               *int    `json:"AddressID"`
}
