package requests

type ItemProductTaxClassification struct {
	Product                  string  `json:"Product"`
	BusinessPartnerID        int     `json:"business_partner"`
	Country                  string  `json:"Country"`
	TaxCategory              string  `json:"TaxCategory"`
	ProductTaxClassification *string `json:"ProductTaxClassification"`
}
