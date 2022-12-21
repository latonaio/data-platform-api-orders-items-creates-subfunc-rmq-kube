package requests

type BuyerSellerDetection struct {
	BusinessPartnerID *int   `json:"business_partner"`
	ServiceLabel      string `json:"service_label"`
	Buyer             *int   `json:"Buyer"`
	Seller            *int   `json:"Seller"`
	BuyerOrSeller     string `json:"BuyerOrSeller"`
}
