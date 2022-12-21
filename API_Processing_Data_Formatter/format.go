package api_processing_data_formatter

import (
	api_input_reader "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Input_Reader"
	"data-platform-api-orders-items-creates-subfunc-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

// Initializer
func (psdc *SDC) ConvertToMetaData(sdc *api_input_reader.SDC) *MetaData {
	pm := &requests.MetaData{
		BusinessPartnerID: sdc.BusinessPartnerID,
		ServiceLabel:      sdc.ServiceLabel,
	}

	data := pm
	res := MetaData{
		BusinessPartnerID: data.BusinessPartnerID,
		ServiceLabel:      data.ServiceLabel,
	}

	return &res
}

// Header
func (psdc *SDC) ConvertToBuyerSellerDetection(sdc *api_input_reader.SDC, buyerOrSeller string) *BuyerSellerDetection {
	pm := &requests.BuyerSellerDetection{
		BusinessPartnerID: sdc.BusinessPartnerID,
		ServiceLabel:      sdc.ServiceLabel,
		Buyer:             sdc.Orders.Buyer,
		Seller:            sdc.Orders.Seller,
	}

	pm.BuyerOrSeller = buyerOrSeller

	data := pm
	res := BuyerSellerDetection{
		BusinessPartnerID: data.BusinessPartnerID,
		ServiceLabel:      data.ServiceLabel,
		Buyer:             data.Buyer,
		Seller:            data.Seller,
		BuyerOrSeller:     data.BuyerOrSeller,
	}

	return &res
}

func (psdc *SDC) ConvertToHeaderBPCustomer(rows *sql.Rows) (*HeaderBPCustomer, error) {
	pm := &requests.HeaderBPCustomer{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_business_partner_customer_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartnerID,
			&pm.Customer,
			&pm.TransactionCurrency,
			&pm.Incoterms,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.BPAccountAssignmentGroup,
		)
		if err != nil {
			return nil, err
		}
	}
	data := pm

	res := HeaderBPCustomer{
		OrderID:                  data.OrderID,
		BusinessPartnerID:        data.BusinessPartnerID,
		Customer:                 data.Customer,
		TransactionCurrency:      data.TransactionCurrency,
		Incoterms:                data.Incoterms,
		PaymentTerms:             data.PaymentTerms,
		PaymentMethod:            data.PaymentMethod,
		BPAccountAssignmentGroup: data.BPAccountAssignmentGroup,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToHeaderBPSupplier(rows *sql.Rows) (*HeaderBPSupplier, error) {
	pm := &requests.HeaderBPSupplier{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_business_partner_supplier_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartnerID,
			&pm.Supplier,
			&pm.TransactionCurrency,
			&pm.Incoterms,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.BPAccountAssignmentGroup,
		)
		if err != nil {
			return nil, err
		}
	}

	data := pm
	res := HeaderBPSupplier{
		OrderID:                  data.OrderID,
		BusinessPartnerID:        data.BusinessPartnerID,
		Supplier:                 data.Supplier,
		TransactionCurrency:      data.TransactionCurrency,
		Incoterms:                data.Incoterms,
		PaymentTerms:             data.PaymentTerms,
		PaymentMethod:            data.PaymentMethod,
		BPAccountAssignmentGroup: data.BPAccountAssignmentGroup,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToHeaderBPCustomerSupplier() *HeaderBPCustomerSupplier {
	pm := &requests.HeaderBPCustomerSupplier{}

	if psdc.BuyerSellerDetection.BuyerOrSeller == "Seller" {
		pm.BusinessPartnerID = psdc.HeaderBPCustomer.BusinessPartnerID
		pm.CustomerOrSupplier = psdc.HeaderBPCustomer.Customer
		pm.TransactionCurrency = psdc.HeaderBPCustomer.TransactionCurrency
		pm.Incoterms = psdc.HeaderBPCustomer.Incoterms
		pm.PaymentTerms = psdc.HeaderBPCustomer.PaymentTerms
		pm.PaymentMethod = psdc.HeaderBPCustomer.PaymentMethod
		pm.BPAccountAssignmentGroup = psdc.HeaderBPCustomer.BPAccountAssignmentGroup
	} else if psdc.BuyerSellerDetection.BuyerOrSeller == "Buyer" {
		pm.BusinessPartnerID = psdc.HeaderBPSupplier.BusinessPartnerID
		pm.CustomerOrSupplier = psdc.HeaderBPSupplier.Supplier
		pm.TransactionCurrency = psdc.HeaderBPSupplier.TransactionCurrency
		pm.Incoterms = psdc.HeaderBPSupplier.Incoterms
		pm.PaymentTerms = psdc.HeaderBPSupplier.PaymentTerms
		pm.PaymentMethod = psdc.HeaderBPSupplier.PaymentMethod
		pm.BPAccountAssignmentGroup = psdc.HeaderBPSupplier.BPAccountAssignmentGroup
	}

	data := pm
	res := HeaderBPCustomerSupplier{
		OrderID:                  data.OrderID,
		BusinessPartnerID:        data.BusinessPartnerID,
		CustomerOrSupplier:       data.CustomerOrSupplier,
		TransactionCurrency:      data.TransactionCurrency,
		Incoterms:                data.Incoterms,
		PaymentTerms:             data.PaymentTerms,
		PaymentMethod:            data.PaymentMethod,
		BPAccountAssignmentGroup: data.BPAccountAssignmentGroup,
	}

	return &res
}

func (psdc *SDC) ConvertToCalculateOrderIDKey() *CalculateOrderIDKey {
	pm := &requests.CalculateOrderIDKey{
		FieldNameWithNumberRange: "OrderID",
	}

	data := pm
	res := CalculateOrderIDKey{
		ServiceLabel:             data.ServiceLabel,
		FieldNameWithNumberRange: data.FieldNameWithNumberRange,
	}

	return &res
}

func (psdc *SDC) ConvertToCalculateOrderIDQueryGets(rows *sql.Rows) (*CalculateOrderIDQueryGets, error) {
	pm := &requests.CalculateOrderIDQueryGets{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_number_range_latest_number_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.ServiceLabel,
			&pm.FieldNameWithNumberRange,
			&pm.OrderIDLatestNumber,
		)
		if err != nil {
			return nil, err
		}
	}

	data := pm
	res := CalculateOrderIDQueryGets{
		ServiceLabel:             data.ServiceLabel,
		FieldNameWithNumberRange: data.FieldNameWithNumberRange,
		OrderIDLatestNumber:      data.OrderIDLatestNumber,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToCalculateOrderID(orderIDLatestNumber *int, orderID int) *CalculateOrderID {
	pm := &requests.CalculateOrderID{}

	pm.OrderIDLatestNumber = orderIDLatestNumber
	pm.OrderID = orderID

	data := pm
	res := CalculateOrderID{
		OrderIDLatestNumber: data.OrderIDLatestNumber,
		OrderID:             data.OrderID,
	}

	return &res
}

func (psdc *SDC) ConvertToPricingDate(inputPricingDate string) *PricingDate {
	pm := &requests.PricingDate{}

	pm.PricingDate = inputPricingDate

	data := pm
	res := PricingDate{
		PricingDate: data.PricingDate,
	}

	return &res
}

// HeaderPartner
func (psdc *SDC) ConvertToHeaderPartnerFunctionKey() *HeaderPartnerFunctionKey {
	pm := &requests.HeaderPartnerFunctionKey{}

	data := pm
	res := HeaderPartnerFunctionKey{
		BusinessPartnerID:  data.BusinessPartnerID,
		CustomerOrSupplier: data.CustomerOrSupplier,
	}

	return &res
}

func (psdc *SDC) ConvertToHeaderPartnerFunction(rows *sql.Rows) (*[]HeaderPartnerFunction, error) {
	var res []HeaderPartnerFunction

	for i := 0; true; i++ {
		pm := &requests.HeaderPartnerFunction{}
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_business_partner_customer_partner_function_data'または'data_platform_business_partner_supplier_partner_function_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartnerID,
			&pm.PartnerCounter,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.DefaultPartner)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, HeaderPartnerFunction{
			BusinessPartnerID: data.BusinessPartnerID,
			PartnerCounter:    data.PartnerCounter,
			PartnerFunction:   data.PartnerFunction,
			BusinessPartner:   data.BusinessPartner,
			DefaultPartner:    data.DefaultPartner,
		})
	}

	return &res, nil
}

func (psdc *SDC) ConvertToHeaderPartnerBPGeneral(rows *sql.Rows) (*[]HeaderPartnerBPGeneral, error) {
	var res []HeaderPartnerBPGeneral

	for i := 0; true; i++ {
		pm := &requests.HeaderPartnerBPGeneral{}
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_business_partner_general_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.BusinessPartnerFullName,
			&pm.BusinessPartnerName,
			&pm.Country,
			&pm.Language,
			&pm.AddressID,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, HeaderPartnerBPGeneral{
			BusinessPartner:         data.BusinessPartner,
			BusinessPartnerFullName: data.BusinessPartnerFullName,
			BusinessPartnerName:     data.BusinessPartnerName,
			Country:                 data.Country,
			Language:                data.Language,
			AddressID:               data.AddressID,
		})
	}

	return &res, nil
}

// HeaderPartnerPlant
func (psdc *SDC) ConvertToHeaderPartnerPlantKey(length int) *[]HeaderPartnerPlantKey {
	var res []HeaderPartnerPlantKey

	for i := 0; i < length; i++ {
		pm := &requests.HeaderPartnerPlantKey{}

		data := pm
		res = append(res, HeaderPartnerPlantKey{
			BusinessPartnerID:              data.BusinessPartnerID,
			CustomerOrSupplier:             data.CustomerOrSupplier,
			PartnerCounter:                 data.PartnerCounter,
			PartnerFunction:                data.PartnerFunction,
			PartnerFunctionBusinessPartner: data.PartnerFunctionBusinessPartner,
		})
	}

	return &res
}

func (psdc *SDC) ConvertToHeaderPartnerPlant(rows *sql.Rows) (*[]HeaderPartnerPlant, error) {
	var res []HeaderPartnerPlant

	for i := 0; true; i++ {
		pm := &requests.HeaderPartnerPlant{}
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_business_partner_supplier_partner_plant_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.PartnerFunction,
			&pm.PlantCounter,
			&pm.Plant,
			&pm.DefaultPlant,
			&pm.DefaultStockConfirmationPlant,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, HeaderPartnerPlant{
			BusinessPartner:               data.BusinessPartner,
			PartnerFunction:               data.PartnerFunction,
			PlantCounter:                  data.PlantCounter,
			Plant:                         data.Plant,
			DefaultPlant:                  data.DefaultPlant,
			DefaultStockConfirmationPlant: data.DefaultStockConfirmationPlant,
		})
	}

	return &res, nil
}

// Item
func (psdc *SDC) ConvertToOrderItem(sdc *api_input_reader.SDC) *[]OrderItem {
	var res []OrderItem

	for i := range sdc.Orders.Item {
		pm := &requests.OrderItem{}

		pm.OrderItemNumber = i + 1

		data := pm
		res = append(res, OrderItem{
			OrderItemNumber: data.OrderItemNumber,
		})
	}

	return &res
}

func (psdc *SDC) ConvertToItemBPTaxClassificationKey() *ItemBPTaxClassificationKey {
	pm := &requests.ItemBPTaxClassificationKey{}

	data := pm
	res := ItemBPTaxClassificationKey{
		BusinessPartnerID:  data.BusinessPartnerID,
		CustomerOrSupplier: data.CustomerOrSupplier,
		DepartureCountry:   data.DepartureCountry,
	}

	return &res
}

func (psdc *SDC) ConvertToItemBPTaxClassification(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*ItemBPTaxClassification, error) {
	pm := &requests.ItemBPTaxClassification{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_business_partner_customer_tax_data'または'data_platform_business_partner_supplier_tax_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartnerID,
			&pm.CustomerOrSupplier,
			&pm.DepartureCountry,
			&pm.BPTaxClassification,
		)
		if err != nil {
			return nil, err
		}
	}

	data := pm
	res := ItemBPTaxClassification{
		BusinessPartnerID:   data.BusinessPartnerID,
		CustomerOrSupplier:  data.CustomerOrSupplier,
		DepartureCountry:    data.DepartureCountry,
		BPTaxClassification: data.BPTaxClassification,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToItemProductTaxClassificationKey(length int) *[]ItemProductTaxClassificationKey {
	var res []ItemProductTaxClassificationKey

	for i := 0; i < length; i++ {
		pm := &requests.ItemProductTaxClassificationKey{
			TaxCategory: "MWST",
		}

		data := pm
		res = append(res, ItemProductTaxClassificationKey{
			Product:           data.Product,
			BusinessPartnerID: data.BusinessPartnerID,
			Country:           data.Country,
			TaxCategory:       data.TaxCategory,
		})
	}
	return &res
}

func (psdc *SDC) ConvertToItemProductTaxClassification(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]ItemProductTaxClassification, error) {
	var itemProductTaxClassification []ItemProductTaxClassification

	for i := 0; true; i++ {
		pm := &requests.ItemProductTaxClassification{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_product_master_tax_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.Product,
			&pm.BusinessPartnerID,
			&pm.Country,
			&pm.TaxCategory,
			&pm.ProductTaxClassification,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		itemProductTaxClassification = append(itemProductTaxClassification, ItemProductTaxClassification{
			Product:                  data.Product,
			BusinessPartnerID:        data.BusinessPartnerID,
			Country:                  data.Country,
			TaxCategory:              data.TaxCategory,
			ProductTaxClassification: data.ProductTaxClassification,
		})
	}

	return &itemProductTaxClassification, nil
}

func (psdc *SDC) ConvertToTaxCode(bpTaxClassification, productTaxClassification *string, product, orderType, taxCode string) *TaxCode {
	pm := requests.TaxCode{}

	pm.Product = product
	pm.BPTaxClassification = bpTaxClassification
	pm.ProductTaxClassification = productTaxClassification
	pm.OrderType = orderType
	pm.TaxCode = taxCode

	data := pm
	res := TaxCode{
		Product:                  data.Product,
		BPTaxClassification:      data.BPTaxClassification,
		ProductTaxClassification: data.ProductTaxClassification,
		OrderType:                data.OrderType,
		TaxCode:                  data.TaxCode,
	}

	return &res
}

func (psdc *SDC) ConvertToItemPMGeneral(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]ItemPMGeneral, error) {
	var res []ItemPMGeneral

	for i := 0; true; i++ {
		pm := &requests.ItemPMGeneral{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_product_master_general_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.Product,
			&pm.ProductStandardID,
			&pm.ProductGroup,
			&pm.BaseUnit,
			&pm.ItemWeightUnit,
			&pm.ProductGrossWeight,
			&pm.ProductNetWeight,
			&pm.ProductAccountAssignmentGroup,
			&pm.CountryOfOrigin,
			&pm.CountryOfOriginLanguage,
			&pm.ItemCategory,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		res = append(res, ItemPMGeneral{
			Product:                       data.Product,
			ProductStandardID:             data.ProductStandardID,
			ProductGroup:                  data.ProductGroup,
			BaseUnit:                      data.BaseUnit,
			ItemWeightUnit:                data.ItemWeightUnit,
			ProductGrossWeight:            data.ProductGrossWeight,
			ProductNetWeight:              data.ProductNetWeight,
			ProductAccountAssignmentGroup: data.ProductAccountAssignmentGroup,
			CountryOfOrigin:               data.CountryOfOrigin,
			CountryOfOriginLanguage:       data.CountryOfOriginLanguage,
			ItemCategory:                  data.ItemCategory,
		})
	}

	return &res, nil
}

func (psdc *SDC) ConvertToItemProductDescriptionKey(length int) *[]ItemProductDescriptionKey {
	var res []ItemProductDescriptionKey

	for i := 0; i < length; i++ {
		pm := &requests.ItemProductDescriptionKey{}

		data := pm
		res = append(res, ItemProductDescriptionKey{
			Product:  data.Product,
			Language: data.Language,
		})
	}

	return &res
}

func (psdc *SDC) ConvertToItemProductDescription(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]ItemProductDescription, error) {
	var res []ItemProductDescription

	for i := 0; true; i++ {
		pm := &requests.ItemProductDescription{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_product_master_product_description_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.Product,
			&pm.Language,
			&pm.ProductDescription,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		res = append(res, ItemProductDescription{
			Product:            data.Product,
			Language:           data.Language,
			ProductDescription: data.ProductDescription,
		})
	}

	return &res, nil
}

func (psdc *SDC) ConvertToStockConfirmationPlant(sdc *api_input_reader.SDC) *StockConfirmationPlant {
	pm := &requests.StockConfirmationPlant{}

	for _, v := range *psdc.HeaderPartnerPlant {
		if *v.DefaultStockConfirmationPlant {
			pm = &requests.StockConfirmationPlant{
				BusinessPartner:               v.BusinessPartner,
				PartnerFunction:               v.PartnerFunction,
				Plant:                         v.Plant,
				DefaultStockConfirmationPlant: v.DefaultStockConfirmationPlant,
			}
			break
		}
	}

	data := pm
	res := StockConfirmationPlant{
		BusinessPartner:               data.BusinessPartner,
		PartnerFunction:               data.PartnerFunction,
		Plant:                         data.Plant,
		DefaultStockConfirmationPlant: data.DefaultStockConfirmationPlant,
	}
	return &res
}

func (psdc *SDC) ConvertToProductMasterBPPlantKey(length int) *[]ProductMasterBPPlantKey {
	var res []ProductMasterBPPlantKey

	for i := 0; i < length; i++ {
		pm := &requests.ProductMasterBPPlantKey{}

		data := pm
		res = append(res, ProductMasterBPPlantKey{
			Product:         data.Product,
			BusinessPartner: data.BusinessPartner,
			Plant:           data.Plant,
		})
	}

	return &res
}

func (psdc *SDC) ConvertToProductMasterBPPlant(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]ProductMasterBPPlant, error) {
	var res []ProductMasterBPPlant

	for i := 0; true; i++ {
		pm := &requests.ProductMasterBPPlant{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_product_master_bp_plant_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.Product,
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.IssuingDeliveryUnit,
			&pm.ReceivingDeliveryUnit,
			&pm.IsBatchManagementRequired,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		res = append(res, ProductMasterBPPlant{
			Product:                   data.Product,
			BusinessPartner:           data.BusinessPartner,
			Plant:                     data.Plant,
			IssuingDeliveryUnit:       data.IssuingDeliveryUnit,
			ReceivingDeliveryUnit:     data.ReceivingDeliveryUnit,
			IsBatchManagementRequired: data.IsBatchManagementRequired,
		})
	}

	return &res, nil
}

func (psdc *SDC) ConvertToProductionPlant(
	sdc *api_input_reader.SDC,
) *[]ProductionPlant {
	var res []ProductionPlant

	for _, v := range *psdc.HeaderPartnerPlant {
		if (v.PartnerFunction == "MANUFACTURER") && *v.DefaultPlant {
			pm := &requests.ProductionPlant{
				PartnerFunction: v.PartnerFunction,
				DefaultPlant:    v.DefaultPlant,
				BusinessPartner: v.BusinessPartner,
				Plant:           v.Plant,
			}

			data := pm
			res = append(res, ProductionPlant{
				PartnerFunction: data.PartnerFunction,
				DefaultPlant:    data.DefaultPlant,
				BusinessPartner: data.BusinessPartner,
				Plant:           data.Plant,
			})
		}
	}

	return &res
}

func (psdc *SDC) ConvertToPaymentMethod(paymentMethod *string) *PaymentMethod {
	pm := &requests.PaymentMethod{}

	pm.PaymentMethod = paymentMethod

	data := pm
	res := PaymentMethod{
		PaymentMethod: data.PaymentMethod,
	}

	return &res
}

func (psdc *SDC) ConvertToItemNetWeight(product string, productNetWeight, orderQuantityInBaseUnit, itemNetWeght *float32) *ItemNetWeight {
	pm := &requests.ItemNetWeight{}

	pm.Product = product
	pm.ProductNetWeight = productNetWeight
	pm.OrderQuantityInBaseUnit = orderQuantityInBaseUnit
	pm.ItemNetWeight = itemNetWeght

	data := pm
	res := ItemNetWeight{
		Product:                 data.Product,
		ProductNetWeight:        data.ProductNetWeight,
		OrderQuantityInBaseUnit: data.OrderQuantityInBaseUnit,
		ItemNetWeight:           data.ItemNetWeight,
	}

	return &res
}

func (psdc *SDC) ConvertToItemGrossWeight(product string, productGrossWeight, orderQuantityInBaseUnit, itemGrossWeght *float32) *ItemGrossWeight {
	pm := &requests.ItemGrossWeight{}

	pm.Product = product
	pm.ProductGrossWeight = productGrossWeight
	pm.OrderQuantityInBaseUnit = orderQuantityInBaseUnit
	pm.ItemGrossWeight = itemGrossWeght

	data := pm
	res := ItemGrossWeight{
		Product:                 data.Product,
		ProductGrossWeight:      data.ProductGrossWeight,
		OrderQuantityInBaseUnit: data.OrderQuantityInBaseUnit,
		ItemGrossWeight:         data.ItemGrossWeight,
	}

	return &res
}

func (psdc *SDC) ConvertToTaxRateKey() *TaxRateKey {
	pm := &requests.TaxRateKey{
		Country: "JP",
	}

	data := pm
	res := TaxRateKey{
		Country:           data.Country,
		TaxCode:           data.TaxCode,
		ValidityEndDate:   data.ValidityEndDate,
		ValidityStartDate: data.ValidityStartDate,
	}

	return &res
}

func (psdc *SDC) ConvertToTaxRate(rows *sql.Rows) (*[]TaxRate, error) {
	var res []TaxRate

	for i := 0; true; i++ {
		pm := &requests.TaxRate{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_price_master_price_master_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.Country,
			&pm.TaxCode,
			&pm.ValidityEndDate,
			&pm.ValidityStartDate,
			&pm.TaxRate,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		res = append(res, TaxRate{
			Country:           data.Country,
			TaxCode:           data.TaxCode,
			ValidityEndDate:   data.ValidityEndDate,
			ValidityStartDate: data.ValidityStartDate,
			TaxRate:           data.TaxRate,
		})
	}

	return &res, nil
}

func (psdc *SDC) ConvertToNetAmount(conditionAmount *[]ConditionAmount) *[]NetAmount {
	var res []NetAmount

	for _, v := range *conditionAmount {
		pm := &requests.NetAmount{}

		pm.Product = v.Product
		pm.NetAmount = v.ConditionAmount

		data := pm
		res = append(res, NetAmount{
			Product:   data.Product,
			NetAmount: data.NetAmount,
		})
	}

	return &res
}

func (psdc *SDC) ConvertToTaxAmount(product, taxCode string, taxRate, netAmount, taxAmount *float32) *TaxAmount {
	pm := &requests.TaxAmount{}

	pm.Product = product
	pm.TaxCode = taxCode
	pm.TaxRate = taxRate
	pm.NetAmount = netAmount
	pm.TaxAmount = taxAmount

	data := pm
	res := TaxAmount{
		Product:   data.Product,
		TaxCode:   data.TaxCode,
		TaxRate:   data.TaxRate,
		NetAmount: data.NetAmount,
		TaxAmount: data.TaxAmount,
	}

	return &res
}

func (psdc *SDC) ConvertToGrossAmount(product string, netAmount, taxAmount, grossAmount *float32) *GrossAmount {
	pm := &requests.GrossAmount{}

	pm.Product = product
	pm.NetAmount = netAmount
	pm.TaxAmount = taxAmount
	pm.GrossAmount = grossAmount

	data := pm
	res := GrossAmount{
		Product:     data.Product,
		NetAmount:   data.NetAmount,
		TaxAmount:   data.TaxAmount,
		GrossAmount: data.GrossAmount,
	}

	return &res
}

// ItemPricingElement
func (psdc *SDC) ConvertToPriceMasterKey(sdc *api_input_reader.SDC) *PriceMasterKey {
	pm := &requests.PriceMasterKey{
		BusinessPartnerID: sdc.BusinessPartnerID,
	}

	data := pm
	res := PriceMasterKey{
		BusinessPartnerID:          data.BusinessPartnerID,
		Product:                    data.Product,
		CustomerOrSupplier:         data.CustomerOrSupplier,
		ConditionValidityEndDate:   data.ConditionValidityEndDate,
		ConditionValidityStartDate: data.ConditionValidityStartDate,
	}

	return &res
}

func (psdc *SDC) ConvertToPriceMaster(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]PriceMaster, error) {
	var res []PriceMaster

	for i := 0; true; i++ {
		pm := &requests.PriceMaster{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_price_master_price_master_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartnerID,
			&pm.Product,
			&pm.CustomerOrSupplier,
			&pm.ConditionValidityEndDate,
			&pm.ConditionValidityStartDate,
			&pm.ConditionRecord,
			&pm.ConditionSequentialNumber,
			&pm.ConditionType,
			&pm.ConditionRateValue,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		res = append(res, PriceMaster{
			BusinessPartnerID:          data.BusinessPartnerID,
			Product:                    data.Product,
			CustomerOrSupplier:         data.CustomerOrSupplier,
			ConditionValidityEndDate:   data.ConditionValidityEndDate,
			ConditionValidityStartDate: data.ConditionValidityStartDate,
			ConditionRecord:            data.ConditionRecord,
			ConditionSequentialNumber:  data.ConditionSequentialNumber,
			ConditionType:              data.ConditionType,
			ConditionRateValue:         data.ConditionRateValue,
		})
	}

	return &res, nil
}

func (psdc *SDC) ConvertToConditionAmount(product string, conditionQuantity *float32, conditionAmount *float32) *ConditionAmount {
	pm := &requests.ConditionAmount{
		ConditionIsManuallyChanged: GetBoolPtr(false),
	}

	pm.Product = product
	pm.ConditionQuantity = conditionQuantity
	pm.ConditionAmount = conditionAmount

	data := pm
	res := ConditionAmount{
		Product:                    data.Product,
		ConditionQuantity:          data.ConditionQuantity,
		ConditionAmount:            data.ConditionAmount,
		ConditionIsManuallyChanged: data.ConditionIsManuallyChanged,
	}

	return &res
}

func GetBoolPtr(b bool) *bool {
	return &b
}
