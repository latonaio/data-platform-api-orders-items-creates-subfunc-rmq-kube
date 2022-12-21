package requests

type ItemNetWeight struct {
	Product                 string   `json:"Product"`
	ProductNetWeight        *float32 `json:"ProductNetWeight"`
	OrderQuantityInBaseUnit *float32 `json:"OrderQuantityInBaseUnit"`
	ItemNetWeight           *float32 `json:"ItemNetWeight"`
}
