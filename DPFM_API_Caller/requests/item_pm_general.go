package requests

type ItemPMGeneral struct {
	Product                       string   `json:"Product"`
	ProductStandardID             *string  `json:"ProductStandardID"`
	ProductGroup                  *string  `json:"ProductGroup"`
	BaseUnit                      *string  `json:"BaseUnit"`
	ItemWeightUnit                *string  `json:"ItemWeightUnit"`
	ProductGrossWeight            *float32 `json:"ProductGrossWeight"`
	ProductNetWeight              *float32 `json:"ProductNetWeight"`
	ProductAccountAssignmentGroup *string  `json:"ProductAccountAssignmentGroup"`
	CountryOfOrigin               *string  `json:"CountryOfOrigin"`
	CountryOfOriginLanguage       *string  `json:"CountryOfOriginLanguage"`
	ItemCategory                  *string  `json:"ItemCategory"`
}
