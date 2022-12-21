package api_processing_data_formatter

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	MetaData                     *MetaData                       `json:"MetaData"`
	BuyerSellerDetection         *BuyerSellerDetection           `json:"BuyerSellerDetection"`
	HeaderBPCustomer             *HeaderBPCustomer               `json:"HeaderBPCustomer"`
	HeaderBPSupplier             *HeaderBPSupplier               `json:"HeaderBPSupplier"`
	HeaderBPCustomerSupplier     *HeaderBPCustomerSupplier       `json:"HeaderBPCustomerSupplier"`
	CalculateOrderID             *CalculateOrderID               `json:"CalculateOrderID"`
	PricingDate                  *PricingDate                    `json:"PricingDate"`
	HeaderPartnerFunction        *[]HeaderPartnerFunction        `json:"HeaderPartnerFunction"`
	HeaderPartnerBPGeneral       *[]HeaderPartnerBPGeneral       `json:"HeaderPartnerBPGeneral"`
	HeaderPartnerPlant           *[]HeaderPartnerPlant           `json:"HeaderPartnerPlant"`
	OrderItem                    *[]OrderItem                    `json:"OrderItem"`
	ItemBPTaxClassification      *ItemBPTaxClassification        `json:"ItemBPTaxClassification"`
	ItemProductTaxClassification *[]ItemProductTaxClassification `json:"ItemProductTaxClassification"`
	TaxCode                      *[]TaxCode                      `json:"TaxCode"`
	ItemPMGeneral                *[]ItemPMGeneral                `json:"ItemPMGeneral"`
	ItemProductDescription       *[]ItemProductDescription       `json:"ItemProductDescription"`
	StockConfirmationPlant       *StockConfirmationPlant         `json:"StockConfirmationPlant"`
	ProductMasterBPPlant         *[]ProductMasterBPPlant         `json:"ProductMasterBPPlant"`
	ProductionPlant              *[]ProductionPlant              `json:"ProductionPlant"`
	PaymentMethod                *PaymentMethod                  `json:"PaymentMethod"`
	ItemGrossWeight              *[]ItemGrossWeight              `json:"ItemGrossWeight"`
	ItemNetWeight                *[]ItemNetWeight                `json:"ItemNetWeight"`
	TaxRate                      *[]TaxRate                      `json:"TaxRate"`
	NetAmount                    *[]NetAmount                    `json:"NetAmount"`
	TaxAmount                    *[]TaxAmount                    `json:"TaxAmount"`
	GrossAmount                  *[]GrossAmount                  `json:"GrossAmount"`
	PriceMaster                  *[]PriceMaster                  `json:"PriceMaster"`
	ConditionAmount              *[]ConditionAmount              `json:"ConditionAmount"`
}

// Initializer
type MetaData struct {
	BusinessPartnerID *int   `json:"business_partner"`
	ServiceLabel      string `json:"service_label"`
}

// Header
type BuyerSellerDetection struct {
	BusinessPartnerID *int   `json:"business_partner"`
	ServiceLabel      string `json:"service_label"`
	Buyer             *int   `json:"Buyer"`
	Seller            *int   `json:"Seller"`
	BuyerOrSeller     string `json:"BuyerOrSeller"`
}

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

type CalculateOrderIDKey struct {
	ServiceLabel             string `json:"service_label"`
	FieldNameWithNumberRange string `json:"FieldNameWithNumberRange"`
}

type CalculateOrderIDQueryGets struct {
	ServiceLabel             string `json:"service_label"`
	FieldNameWithNumberRange string `json:"FieldNameWithNumberRange"`
	OrderIDLatestNumber      *int   `json:"OrderIDLatestNumber"`
}

type CalculateOrderID struct {
	OrderIDLatestNumber *int `json:"OrderIDLatestNumber"`
	OrderID             int  `json:"OrderID"`
}

type PricingDate struct {
	PricingDate string `json:"PricingDate"`
}

// HeaderPartner
type HeaderPartnerFunctionKey struct {
	BusinessPartnerID  *int `json:"business_partner"`
	CustomerOrSupplier *int `json:"CustomerOrSupplier"`
}

type HeaderPartnerFunction struct {
	BusinessPartnerID int     `json:"business_partner"`
	PartnerCounter    int     `json:"PartnerCounter"`
	PartnerFunction   *string `json:"PartnerFunction"`
	BusinessPartner   *int    `json:"BusinessPartner"`
	DefaultPartner    *bool   `json:"DefaultPartner"`
}

type HeaderPartnerBPGeneral struct {
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     string  `json:"BusinessPartnerName"`
	Country                 string  `json:"Country"`
	Language                string  `json:"Language"`
	AddressID               *int    `json:"AddressID"`
}

// HeaderPartnerPlant
type HeaderPartnerPlantKey struct {
	BusinessPartnerID              *int    `json:"business_partner"`
	CustomerOrSupplier             *int    `json:"CustomerOrSupplier"`
	PartnerCounter                 int     `json:"PartnerCounter"`
	PartnerFunction                *string `json:"PartnerFunction"`
	PartnerFunctionBusinessPartner *int    `json:"PartnerFunctionBusinessPartner"`
}

type HeaderPartnerPlant struct {
	BusinessPartner               int     `json:"BusinessPartner"`
	PartnerFunction               string  `json:"PartnerFunction"`
	PlantCounter                  int     `json:"PlantCounter"`
	Plant                         *string `json:"Plant"`
	DefaultPlant                  *bool   `json:"DefaultPlant"`
	DefaultStockConfirmationPlant *bool   `json:"DefaultStockConfirmationPlant"`
}

// Item
type OrderItem struct {
	OrderItemNumber int `json:"OrderItemNumber"`
}

type ItemBPTaxClassificationKey struct {
	BusinessPartnerID  *int   `json:"business_partner"`
	CustomerOrSupplier *int   `json:"CustomerOrSupplier"`
	DepartureCountry   string `json:"DepartureCountry"`
}

type ItemBPTaxClassification struct {
	BusinessPartnerID   int     `json:"business_partner"`
	CustomerOrSupplier  int     `json:"CustomerOrSupplier"`
	DepartureCountry    string  `json:"DepartureCountry"`
	BPTaxClassification *string `json:"BPTaxClassification"`
}

type ItemProductTaxClassificationKey struct {
	Product           string `json:"Product"`
	BusinessPartnerID *int   `json:"business_partner"`
	Country           string `json:"Country"`
	TaxCategory       string `json:"TaxCategory"`
}

type ItemProductTaxClassification struct {
	Product                  string  `json:"Product"`
	BusinessPartnerID        int     `json:"business_partner"`
	Country                  string  `json:"Country"`
	TaxCategory              string  `json:"TaxCategory"`
	ProductTaxClassification *string `json:"ProductTaxClassification"`
}

type TaxCode struct {
	Product                  string  `json:"Product"`
	BPTaxClassification      *string `json:"BPTaxClassification"`
	ProductTaxClassification *string `json:"ProductTaxClassification"`
	OrderType                string  `json:"OrderType"`
	TaxCode                  string  `json:"TaxCode"`
}

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

type ItemProductDescriptionKey struct {
	Product  string `json:"Product"`
	Language string `json:"Language"`
}

type ItemProductDescription struct {
	Product            string  `json:"Product"`
	Language           string  `json:"Language"`
	ProductDescription *string `json:"ProductDescription"`
}

type StockConfirmationPlant struct {
	BusinessPartner               int     `json:"BusinessPartner"`
	PartnerFunction               string  `json:"PartnerFunction"`
	Plant                         *string `json:"Plant"`
	DefaultStockConfirmationPlant *bool   `json:"DefaultStockConfirmationPlant"`
}

type ProductMasterBPPlantKey struct {
	Product         string  `json:"Product"`
	BusinessPartner *int    `json:"BusinessPartner"`
	Plant           *string `json:"Plant"`
}

type ProductMasterBPPlant struct {
	Product                   string  `json:"Product"`
	BusinessPartner           int     `json:"BusinessPartner"`
	Plant                     string  `json:"Plant"`
	IssuingDeliveryUnit       *string `json:"IssuingDeliveryUnit"`
	ReceivingDeliveryUnit     *string `json:"ReceivingDeliveryUnit"`
	IsBatchManagementRequired *bool   `json:"IsBatchManagementRequired"`
}

type ProductionPlant struct {
	PartnerFunction string  `json:"PartnerFunction"`
	DefaultPlant    *bool   `json:"DefaultPlant"`
	BusinessPartner int     `json:"BusinessPartner"`
	Plant           *string `json:"Plant"`
}

type PaymentMethod struct {
	PaymentMethod *string `json:"PaymentMethod"`
}

type ItemGrossWeight struct {
	Product                 string   `json:"Product"`
	ProductGrossWeight      *float32 `json:"ProductGrossWeight"`
	OrderQuantityInBaseUnit *float32 `json:"OrderQuantityInBaseUnit"`
	ItemGrossWeight         *float32 `json:"ItemGrossWeight"`
}

type ItemNetWeight struct {
	Product                 string   `json:"Product"`
	ProductNetWeight        *float32 `json:"ProductNetWeight"`
	OrderQuantityInBaseUnit *float32 `json:"OrderQuantityInBaseUnit"`
	ItemNetWeight           *float32 `json:"ItemNetWeight"`
}

type TaxRateKey struct {
	Country           string   `json:"Country"`
	TaxCode           []string `json:"TaxCode"`
	ValidityEndDate   string   `json:"ValidityEndDate"`
	ValidityStartDate string   `json:"ValidityStartDate"`
}

type TaxRate struct {
	Country           string   `json:"Country"`
	TaxCode           string   `json:"TaxCode"`
	ValidityEndDate   string   `json:"ValidityEndDate"`
	ValidityStartDate *string  `json:"ValidityStartDate"`
	TaxRate           *float32 `json:"TaxRate"`
}

type NetAmount struct {
	Product   string   `json:"Product"`
	NetAmount *float32 `json:"NetAmount"`
}

type TaxAmount struct {
	Product   string   `json:"Product"`
	TaxCode   string   `json:"TaxCode"`
	TaxRate   *float32 `json:"TaxRate"`
	NetAmount *float32 `json:"NetAmount"`
	TaxAmount *float32 `json:"TaxAmount"`
}

type GrossAmount struct {
	Product     string   `json:"Product"`
	NetAmount   *float32 `json:"NetAmount"`
	TaxAmount   *float32 `json:"TaxAmount"`
	GrossAmount *float32 `json:"GrossAmount"`
}

// ItemPricingElement
type PriceMasterKey struct {
	BusinessPartnerID          *int     `json:"business_partner"`
	Product                    []string `json:"Product"`
	CustomerOrSupplier         *int     `json:"CustomerOrSupplier"`
	ConditionValidityEndDate   string   `json:"ConditionValidityEndDate"`
	ConditionValidityStartDate string   `json:"ConditionValidityStartDate"`
}

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

type ConditionAmount struct {
	Product                    string   `json:"Product"`
	ConditionQuantity          *float32 `json:"ConditionQuantity"`
	ConditionAmount            *float32 `json:"ConditionAmount"`
	ConditionIsManuallyChanged *bool    `json:"ConditionIsManuallyChanged"`
}
