package requests

type HeaderBPCustomerSupplier struct {
	OrderID                  *int    `json:"OrderID"`
	BusinessPartnerID        int     `json:"business_partner"`
	CustomerOrSupplier       int     `json:"CustomerOrSupplier"`
	TransactionCurrency      *string `json:"TransactionCurrency"`
	Incoterms                *string `json:"Incoterms"`
	PaymentTerms             *string `json:"PaymentTerms"`
	PaymentMethod            *string `json:"PaymentMethod"`
	BPAccountAssignmentGroup *string `json:"BPAccountAssignmentGroup"`
}
