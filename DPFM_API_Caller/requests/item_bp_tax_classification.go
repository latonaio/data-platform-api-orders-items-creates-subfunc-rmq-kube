package requests

type ItemBPTaxClassification struct {
	BusinessPartnerID   int     `json:"business_partner"`
	CustomerOrSupplier  int     `json:"CustomerOrSupplier"`
	DepartureCountry    string  `json:"DepartureCountry"`
	BPTaxClassification *string `json:"BPTaxClassification"`
}
