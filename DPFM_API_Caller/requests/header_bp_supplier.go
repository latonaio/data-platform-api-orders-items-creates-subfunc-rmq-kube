package requests

type HeaderBPSupplier struct {
	OrderID                  *int    `json:"OrderID"`
	BusinessPartnerID        int     `json:"business_partner"`
	Supplier                 int     `json:"Supplier"`
	TransactionCurrency      *string `json:"TransactionCurrency"`
	Incoterms                *string `json:"Incoterms"`
	PaymentTerms             *string `json:"PaymentTerms"`
	PaymentMethod            *string `json:"PaymentMethod"`
	BPAccountAssignmentGroup *string `json:"BPAccountAssignmentGroup"`
}
