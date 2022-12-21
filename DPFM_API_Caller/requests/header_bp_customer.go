package requests

type HeaderBPCustomer struct {
	OrderID                  *int    `json:"OrderID"`
	BusinessPartnerID        int     `json:"business_partner"`
	Customer                 int     `json:"Customer"`
	TransactionCurrency      *string `json:"TransactionCurrency"`
	Incoterms                *string `json:"Incoterms"`
	PaymentTerms             *string `json:"PaymentTerms"`
	PaymentMethod            *string `json:"PaymentMethod"`
	BPAccountAssignmentGroup *string `json:"BPAccountAssignmentGroup"`
}
