package requests

type TaxCode struct {
	Product                  string  `json:"Product"`
	BPTaxClassification      *string `json:"BPTaxClassification"`
	ProductTaxClassification *string `json:"ProductTaxClassification"`
	OrderType                string  `json:"OrderType"`
	TaxCode                  string  `json:"TaxCode"`
}
