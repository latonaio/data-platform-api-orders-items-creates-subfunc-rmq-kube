package requests

type ItemBPTaxClassificationKey struct {
	OrderID            *int   `json:"OrderID"`
	BusinessPartnerID  *int   `json:"business_partner"`
	CustomerOrSupplier *int   `json:"CustomerOrSupplier"`
	DepartureCountry   string `json:"DepartureCountry"`
}
