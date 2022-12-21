package requests

type ItemProductTaxClassificationKey struct {
	Product           string `json:"Product"`
	BusinessPartnerID *int   `json:"business_partner"`
	Country           string `json:"Country"`
	TaxCategory       string `json:"TaxCategory"`
}
