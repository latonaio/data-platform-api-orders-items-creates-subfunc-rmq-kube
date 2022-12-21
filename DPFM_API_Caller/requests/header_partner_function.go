package requests

type HeaderPartnerFunction struct {
	BusinessPartnerID int     `json:"business_partner"`
	PartnerCounter    int     `json:"PartnerCounter"`
	PartnerFunction   *string `json:"PartnerFunction"`
	BusinessPartner   *int    `json:"BusinessPartner"`
	DefaultPartner    *bool   `json:"DefaultPartner"`
}
