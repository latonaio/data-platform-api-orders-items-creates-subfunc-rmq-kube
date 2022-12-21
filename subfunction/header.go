package subfunction

import (
	api_input_reader "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"database/sql"
	"fmt"
	"time"
)

func (f *SubFunction) BuyerSellerDetection(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.BuyerSellerDetection, error) {
	var buyerOrSeller string
	metaData := psdc.MetaData

	if *metaData.BusinessPartnerID == *sdc.Orders.Buyer && *metaData.BusinessPartnerID != *sdc.Orders.Seller {
		buyerOrSeller = "Buyer"
	} else if *metaData.BusinessPartnerID != *sdc.Orders.Buyer && *metaData.BusinessPartnerID == *sdc.Orders.Seller {
		buyerOrSeller = "Seller"
	} else {
		return nil, fmt.Errorf("business_partnerがBuyerまたはSellerと一致しません")
	}

	buyerSellerDetection := psdc.ConvertToBuyerSellerDetection(sdc, buyerOrSeller)

	return buyerSellerDetection, nil
}

func (f *SubFunction) HeaderBPCustomerSupplier(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.HeaderBPCustomerSupplier, error) {
	var rows *sql.Rows
	var err error

	buyerSellerDetection := psdc.BuyerSellerDetection
	if buyerSellerDetection.BuyerOrSeller == "Seller" {
		rows, err = f.db.Query(
			`SELECT BusinessPartner, Customer, Currency, Incoterms, PaymentTerms, PaymentMethod, BPAccountAssignmentGroup
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_customer_data
		WHERE (BusinessPartner, Customer) = (?, ?);`, buyerSellerDetection.BusinessPartnerID, buyerSellerDetection.Buyer,
		)
		if err != nil {
			return nil, err
		}
		psdc.HeaderBPCustomer, err = psdc.ConvertToHeaderBPCustomer(rows)
		if err != nil {
			return nil, err
		}

		rows, err = f.db.Query(
			`SELECT BusinessPartner, Supplier, Currency, Incoterms, PaymentTerms, PaymentMethod, BPAccountAssignmentGroup
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_supplier_data
		WHERE (BusinessPartner, Supplier) = (?, ?);`, buyerSellerDetection.Buyer, buyerSellerDetection.Seller,
		)
		if err != nil {
			return nil, err
		}
		psdc.HeaderBPSupplier, err = psdc.ConvertToHeaderBPSupplier(rows)
		if err != nil {
			return nil, err
		}
	} else if buyerSellerDetection.BuyerOrSeller == "Buyer" {
		rows, err = f.db.Query(
			`SELECT BusinessPartner, Supplier, Currency, Incoterms, PaymentTerms, PaymentMethod, BPAccountAssignmentGroup
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_supplier_data
		WHERE (BusinessPartner, Supplier) = (?, ?);`, buyerSellerDetection.BusinessPartnerID, buyerSellerDetection.Seller,
		)
		if err != nil {
			return nil, err
		}
		psdc.HeaderBPSupplier, err = psdc.ConvertToHeaderBPSupplier(rows)
		if err != nil {
			return nil, err
		}

		rows, err = f.db.Query(
			`SELECT BusinessPartner, Customer, Currency, Incoterms, PaymentTerms, PaymentMethod, BPAccountAssignmentGroup
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_customer_data
		WHERE (BusinessPartner, Customer) = (?, ?);`, buyerSellerDetection.Seller, buyerSellerDetection.Buyer,
		)
		if err != nil {
			return nil, err
		}
		psdc.HeaderBPCustomer, err = psdc.ConvertToHeaderBPCustomer(rows)
		if err != nil {
			return nil, err
		}
	}

	data := psdc.ConvertToHeaderBPCustomerSupplier()

	return data, err
}

func (f *SubFunction) CalculateOrderID(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.CalculateOrderID, error) {
	metaData := psdc.MetaData
	dataKey := psdc.ConvertToCalculateOrderIDKey()

	dataKey.ServiceLabel = metaData.ServiceLabel

	rows, err := f.db.Query(
		`SELECT ServiceLabel, FieldNameWithNumberRange, LatestNumber
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_number_range_latest_number_data
		WHERE (ServiceLabel, FieldNameWithNumberRange) = (?, ?);`, dataKey.ServiceLabel, dataKey.FieldNameWithNumberRange,
	)
	if err != nil {
		return nil, err
	}

	dataQueryGets, err := psdc.ConvertToCalculateOrderIDQueryGets(rows)
	if err != nil {
		return nil, err
	}

	orderIDLatestNumber := dataQueryGets.OrderIDLatestNumber
	orderID := *dataQueryGets.OrderIDLatestNumber + 1

	data := psdc.ConvertToCalculateOrderID(orderIDLatestNumber, orderID)

	return data, err
}

func (f *SubFunction) PricingDate(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.PricingDate {
	var data *api_processing_data_formatter.PricingDate

	if sdc.Orders.PricingDate != nil {
		if *sdc.Orders.PricingDate != "" {
			data = psdc.ConvertToPricingDate(*sdc.Orders.PricingDate)
		}
	} else {
		data = psdc.ConvertToPricingDate(GetDateStr())
	}

	return data
}

func GetDateStr() string {
	day := time.Now()
	return day.Format("2006-01-02")
}
