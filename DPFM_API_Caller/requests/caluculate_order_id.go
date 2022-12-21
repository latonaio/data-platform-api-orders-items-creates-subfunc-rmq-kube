package requests

type CalculateOrderID struct {
	OrderIDLatestNumber *int `json:"OrderIDLatestNumber"`
	OrderID             int  `json:"OrderID"`
}
