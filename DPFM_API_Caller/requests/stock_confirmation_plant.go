package requests

type StockConfirmationPlant struct {
	BusinessPartner               int     `json:"BusinessPartner"`
	PartnerFunction               string  `json:"PartnerFunction"`
	Plant                         *string `json:"Plant"`
	DefaultStockConfirmationPlant *bool   `json:"DefaultStockConfirmationPlant"`
}
