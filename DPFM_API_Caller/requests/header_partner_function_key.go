package requests

type HeaderPartnerFunctionKey struct {
	BusinessPartnerID  *int `json:"business_partner"`
	CustomerOrSupplier *int `json:"CustomerOrSupplier"`
}
