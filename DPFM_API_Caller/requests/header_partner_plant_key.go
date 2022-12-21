package requests

type HeaderPartnerPlantKey struct {
	BusinessPartnerID              *int    `json:"business_partner"`
	CustomerOrSupplier             *int    `json:"CustomerOrSupplier"`
	PartnerCounter                 int     `json:"PartnerCounter"`
	PartnerFunction                *string `json:"PartnerFunction"`
	PartnerFunctionBusinessPartner *int    `json:"PartnerFunctionBusinessPartner"`
}
