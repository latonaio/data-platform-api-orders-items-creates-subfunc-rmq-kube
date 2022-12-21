package requests

type CalculateOrderIDQueryGets struct {
	ServiceLabel             string `json:"service_label"`
	FieldNameWithNumberRange string `json:"FieldNameWithNumberRange"`
	OrderIDLatestNumber      *int   `json:"OrderIDLatestNumber"`
}
