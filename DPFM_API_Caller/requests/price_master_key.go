package requests

type PriceMasterKey struct {
	BusinessPartnerID          *int     `json:"business_partner"`
	Product                    []string `json:"Product"`
	CustomerOrSupplier         *int     `json:"CustomerOrSupplier"`
	ConditionValidityEndDate   string   `json:"ConditionValidityEndDate"`
	ConditionValidityStartDate string   `json:"ConditionValidityStartDate"`
}
