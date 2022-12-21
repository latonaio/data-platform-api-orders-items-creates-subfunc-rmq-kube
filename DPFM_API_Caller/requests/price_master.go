package requests

type PriceMaster struct {
	BusinessPartnerID          int      `json:"business_partner"`
	Product                    string   `json:"Product"`
	CustomerOrSupplier         *int     `json:"CustomerOrSupplier"`
	ConditionValidityEndDate   string   `json:"ConditionValidityEndDate"`
	ConditionValidityStartDate string   `json:"ConditionValidityStartDate"`
	ConditionRecord            int      `json:"ConditionRecord"`
	ConditionSequentialNumber  int      `json:"ConditionSequentialNumber"`
	ConditionType              string   `json:"ConditionType"`
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
}
