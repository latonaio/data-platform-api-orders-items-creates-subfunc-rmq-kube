package requests

type TaxRateKey struct {
	Country           string   `json:"Country"`
	TaxCode           []string `json:"TaxCode"`
	ValidityEndDate   string   `json:"ValidityEndDate"`
	ValidityStartDate string   `json:"ValidityStartDate"`
}
