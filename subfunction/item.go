package subfunction

import (
	api_input_reader "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"database/sql"
	"fmt"
	"math"
	"strings"

	"golang.org/x/xerrors"
)

func (f *SubFunction) OrderItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *[]api_processing_data_formatter.OrderItem {

	data := psdc.ConvertToOrderItem(sdc)

	return data
}

func (f *SubFunction) ItemBPTaxClassification(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.ItemBPTaxClassification, error) {
	var rows *sql.Rows
	var err error

	buyerSellerDetection := psdc.BuyerSellerDetection
	dataKey := psdc.ConvertToItemBPTaxClassificationKey()
	if err != nil {
		return nil, err
	}

	dataKey.BusinessPartnerID = buyerSellerDetection.BusinessPartnerID

	headerPartnerBPGeneralMap := make(map[int]api_processing_data_formatter.HeaderPartnerBPGeneral, len(*psdc.HeaderPartnerFunction))
	for _, v := range *psdc.HeaderPartnerBPGeneral {
		headerPartnerBPGeneralMap[v.BusinessPartner] = v
	}

	for _, v := range *psdc.HeaderPartnerFunction {
		if *v.PartnerFunction == "BUYER" {
			dataKey.DepartureCountry = headerPartnerBPGeneralMap[*v.BusinessPartner].Country
			break
		}
	}

	if buyerSellerDetection.BuyerOrSeller == "Seller" {
		dataKey.CustomerOrSupplier = buyerSellerDetection.Buyer
		rows, err = f.db.Query(
			`SELECT BusinessPartner, Customer, DepartureCountry, BPTaxClassification
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_customer_tax_data
		WHERE (BusinessPartner, Customer, DepartureCountry) = (?, ?, ?);`, dataKey.BusinessPartnerID, dataKey.CustomerOrSupplier, dataKey.DepartureCountry,
		)
		if err != nil {
			return nil, err
		}
	} else if buyerSellerDetection.BuyerOrSeller == "Buyer" {
		dataKey.CustomerOrSupplier = buyerSellerDetection.Seller
		rows, err = f.db.Query(
			`SELECT BusinessPartner, Supplier, DepartureCountry, BPTaxClassification
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_supplier_tax_data
		WHERE (BusinessPartner, Supplier, DepartureCountry) = (?, ?, ?);`, dataKey.BusinessPartnerID, dataKey.CustomerOrSupplier, dataKey.DepartureCountry,
		)
		if err != nil {
			return nil, err
		}
	}

	data, err := psdc.ConvertToItemBPTaxClassification(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) ItemProductTaxClassification(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.ItemProductTaxClassification, error) {
	var args []interface{}
	var err error

	buyerSellerDetection := psdc.BuyerSellerDetection
	item := sdc.Orders.Item
	dataKey := psdc.ConvertToItemProductTaxClassificationKey(len(item))
	if err != nil {
		return nil, err
	}

	for i, v := range sdc.Orders.Item {
		(*dataKey)[i].Product = v.Product
	}

	for i := range *dataKey {
		(*dataKey)[i].BusinessPartnerID = buyerSellerDetection.BusinessPartnerID
	}

	headerPartnerBPGeneralMap := make(map[int]api_processing_data_formatter.HeaderPartnerBPGeneral, len(*psdc.HeaderPartnerFunction))
	for _, v := range *psdc.HeaderPartnerBPGeneral {
		headerPartnerBPGeneralMap[v.BusinessPartner] = v
	}

	var country string
	for _, v := range *psdc.HeaderPartnerFunction {
		if buyerSellerDetection.BuyerOrSeller == "Seller" {
			if *v.PartnerFunction == "BUYER" {
				country = headerPartnerBPGeneralMap[*v.BusinessPartner].Country
				break
			}
		} else if buyerSellerDetection.BuyerOrSeller == "Buyer" {
			if *v.PartnerFunction == "SELLER" {
				country = headerPartnerBPGeneralMap[*v.BusinessPartner].Country
				break
			}
		}
	}

	for i := range *dataKey {
		(*dataKey)[i].Country = country
	}

	repeat := strings.Repeat("(?, ?, ?, ?),", len(*dataKey)-1) + "(?, ?, ?, ?)" // keyの数
	for _, tag := range *dataKey {
		args = append(args, tag.Product, tag.BusinessPartnerID, tag.Country, tag.TaxCategory)
	}

	rows, err := f.db.Query(
		`SELECT Product, BusinessPartner, Country, TaxCategory, ProductTaxClassification
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_tax_data
		WHERE (Product, BusinessPartner, Country, TaxCategory) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		fmt.Printf("err = %+v \n", err)
		return nil, err
	}

	data, err := psdc.ConvertToItemProductTaxClassification(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) TaxCode(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.TaxCode, error) {
	var data []api_processing_data_formatter.TaxCode

	orderType := sdc.Orders.OrderType
	bpTaxClassification := psdc.ItemBPTaxClassification.BPTaxClassification

	for _, v := range *psdc.ItemProductTaxClassification {
		var taxCode string
		product := v.Product
		productTaxClassification := v.ProductTaxClassification
		if bpTaxClassification != nil && productTaxClassification != nil {
			switch {
			case *bpTaxClassification == "1" && *productTaxClassification == "1" && orderType == "販売系":
				taxCode = "A1"
			case *bpTaxClassification == "0" && *productTaxClassification == "0" && orderType == "販売系":
				taxCode = "A0"
			case *bpTaxClassification == "0" && *productTaxClassification == "1" && orderType == "販売系":
				taxCode = "A0"
			case *bpTaxClassification == "1" && *productTaxClassification == "0" && orderType == "販売系":
				taxCode = "A0"
			case *bpTaxClassification == "1" && *productTaxClassification == "1" && orderType == "購買系":
				taxCode = "V1"
			case *bpTaxClassification == "0" && *productTaxClassification == "0" && orderType == "購買系":
				taxCode = "V0"
			case *bpTaxClassification == "0" && *productTaxClassification == "1" && orderType == "購買系":
				taxCode = "V0"
			case *bpTaxClassification == "1" && *productTaxClassification == "0" && orderType == "購買系":
				taxCode = "V0"
			default:
				return nil, xerrors.Errorf("TaxCodeが決定できません。")
			}
		}

		datum := psdc.ConvertToTaxCode(bpTaxClassification, productTaxClassification, product, orderType, taxCode)
		data = append(data, *datum)
	}

	return &data, nil
}

func (f *SubFunction) ItemPMGeneral(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.ItemPMGeneral, error) {
	var args []interface{}
	repeat := strings.Repeat("?,", len(sdc.Orders.Item)-1) + "?"
	for _, tag := range sdc.Orders.Item {
		args = append(args, tag.Product)
	}

	rows, err := f.db.Query(
		`SELECT Product, ProductStandardID, ProductGroup, BaseUnit, WeightUnit, GrossWeight, NetWeight, ProductAccountAssignmentGroup, CountryOfOrigin, CountryOfOriginLanguage, ItemCategory
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_general_data
		WHERE Product IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		fmt.Printf("err = %+v \n", err)
		return nil, err
	}
	data, err := psdc.ConvertToItemPMGeneral(sdc, rows)
	if err != nil {
		fmt.Printf("err = %+v \n", err)
		return nil, err
	}

	return data, err
}

func (f *SubFunction) ItemProductDescription(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.ItemProductDescription, error) {
	var args []interface{}

	item := sdc.Orders.Item
	dataKey := psdc.ConvertToItemProductDescriptionKey(len(item))

	for i := range *psdc.ItemPMGeneral {
		(*dataKey)[i].Product = (*psdc.ItemPMGeneral)[i].Product
		(*dataKey)[i].Language = *(*psdc.ItemPMGeneral)[i].CountryOfOriginLanguage
	}

	repeat := strings.Repeat("(?, ?),", len(*dataKey)-1) + "(?, ?)"
	for _, tag := range *dataKey {
		args = append(args, tag.Product, tag.Language)
	}

	rows, err := f.db.Query(
		`SELECT Product, Language, ProductDescription
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_product_description_data
		WHERE (Product, Language) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		fmt.Printf("err = %+v \n", err)
		return nil, err
	}
	data, err := psdc.ConvertToItemProductDescription(sdc, rows)
	if err != nil {
		fmt.Printf("err = %+v \n", err)
		return nil, err
	}

	return data, err
}

func (f *SubFunction) StockConfirmationPlant(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.StockConfirmationPlant {

	data := psdc.ConvertToStockConfirmationPlant(sdc)

	return data
}

func (f *SubFunction) ProductMasterBPPlant(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.ProductMasterBPPlant, error) {
	var args []interface{}

	item := sdc.Orders.Item
	getPartnerPlantData := psdc.StockConfirmationPlant
	dataKey := psdc.ConvertToProductMasterBPPlantKey(len(item))

	for i, v := range item {
		(*dataKey)[i].Product = v.Product
		(*dataKey)[i].BusinessPartner = &getPartnerPlantData.BusinessPartner
		(*dataKey)[i].Plant = getPartnerPlantData.Plant
	}

	repeat := strings.Repeat("(?,?,?),", len(*dataKey)-1) + "(?,?,?)"
	for _, tag := range *dataKey {
		args = append(args, tag.Product, tag.BusinessPartner, tag.Plant)
	}

	rows, err := f.db.Query(
		`SELECT Product, BusinessPartner, Plant, IssuingDeliveryUnit, ReceivingDeliveryUnit, IsBatchManagementRequired
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_bp_plant_data
		WHERE (Product, BusinessPartner, Plant) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		fmt.Printf("err = %+v \n", err)
		return nil, err
	}
	data, err := psdc.ConvertToProductMasterBPPlant(sdc, rows)
	if err != nil {
		fmt.Printf("err = %+v \n", err)
		return nil, err
	}

	return data, err
}

func (f *SubFunction) ProductionPlant(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *[]api_processing_data_formatter.ProductionPlant {
	data := psdc.ConvertToProductionPlant(sdc)

	return data
}

func (f *SubFunction) PaymentMethod(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.PaymentMethod {
	paymentMethod := psdc.HeaderBPCustomerSupplier.PaymentMethod

	data := psdc.ConvertToPaymentMethod(paymentMethod)

	return data
}

func (f *SubFunction) ItemNetWeight(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *[]api_processing_data_formatter.ItemNetWeight {
	var data []api_processing_data_formatter.ItemNetWeight

	item := sdc.Orders.Item
	itemMap := make(map[string]api_input_reader.Item, len(item))
	for _, v := range item {
		itemMap[v.Product] = v
	}

	for _, v := range *psdc.ItemPMGeneral {
		product := v.Product
		productNetWeight := v.ProductNetWeight
		orderQuantityInBaseUnit := itemMap[product].OrderQuantityInBaseUnit
		itemNetWeight := ParseFloat32Ptr(*productNetWeight * *orderQuantityInBaseUnit)

		datum := psdc.ConvertToItemNetWeight(product, productNetWeight, orderQuantityInBaseUnit, itemNetWeight)
		data = append(data, *datum)
	}

	return &data
}

func (f *SubFunction) ItemGrossWeight(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *[]api_processing_data_formatter.ItemGrossWeight {
	var data []api_processing_data_formatter.ItemGrossWeight

	item := sdc.Orders.Item
	itemMap := make(map[string]api_input_reader.Item, len(item))
	for _, v := range item {
		itemMap[v.Product] = v
	}

	for _, v := range *psdc.ItemPMGeneral {
		product := v.Product
		productGrossWeight := v.ProductGrossWeight
		orderQuantityInBaseUnit := itemMap[product].OrderQuantityInBaseUnit
		itemGrossWeight := ParseFloat32Ptr(*productGrossWeight * *orderQuantityInBaseUnit)

		datum := psdc.ConvertToItemGrossWeight(product, productGrossWeight, orderQuantityInBaseUnit, itemGrossWeight)
		data = append(data, *datum)
	}

	return &data
}

func (f *SubFunction) TaxRate(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.TaxRate, error) {
	var args []interface{}

	dataKey := psdc.ConvertToTaxRateKey()

	for _, v := range *psdc.TaxCode {
		dataKey.TaxCode = append(dataKey.TaxCode, v.TaxCode)
	}

	dataKey.ValidityEndDate = GetDateStr()
	dataKey.ValidityStartDate = GetDateStr()

	repeat := strings.Repeat("?,", len(dataKey.TaxCode)-1) + "?"
	args = append(args, dataKey.Country)
	for _, v := range dataKey.TaxCode {
		args = append(args, v)
	}
	args = append(args, dataKey.ValidityEndDate, dataKey.ValidityStartDate)

	rows, err := f.db.Query(
		`SELECT Country, TaxCode, ValidityEndDate, ValidityStartDate, TaxRate
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_tax_code_tax_rate_data
		WHERE Country = ?
		AND TaxCode IN ( `+repeat+` )
		AND ValidityEndDate >= ?
		AND ValidityStartDate <= ?;`, args...,
	)
	if err != nil {
		fmt.Printf("err = %+v \n", err)
		return nil, err
	}

	data, err := psdc.ConvertToTaxRate(rows)
	if err != nil {
		fmt.Printf("err = %+v \n", err)
		return nil, err
	}

	return data, err
}

func (f *SubFunction) NetAmount(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *[]api_processing_data_formatter.NetAmount {
	conditionAmount := psdc.ConditionAmount

	data := psdc.ConvertToNetAmount(conditionAmount)

	return data
}

func (f *SubFunction) TaxAmount(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.TaxAmount, error) {
	var data []api_processing_data_formatter.TaxAmount

	item := sdc.Orders.Item
	itemMap := make(map[string]api_input_reader.Item, len(item))
	for _, v := range item {
		itemMap[v.Product] = v
	}

	taxRate := psdc.TaxRate
	taxRateMap := make(map[string]api_processing_data_formatter.TaxRate, len(*taxRate))
	for _, v := range *taxRate {
		taxRateMap[v.TaxCode] = v
	}

	netAmount := psdc.NetAmount
	netAmountMap := make(map[string]api_processing_data_formatter.NetAmount, len(*netAmount))
	for _, v := range *netAmount {
		netAmountMap[v.Product] = v
	}

	for _, v := range *psdc.TaxCode {
		var taxAmount *float32
		if v.TaxCode == "A1" || v.TaxCode == "V1" {
			taxAmount, _ = CalculateTaxAmount(taxRateMap[v.TaxCode].TaxRate, netAmountMap[v.Product].NetAmount)
		} else {
			taxAmount = ParseFloat32Ptr(0)
		}

		if itemMap[v.Product].TaxAmount == nil {
			datum := psdc.ConvertToTaxAmount(v.Product, v.TaxCode, taxRateMap[v.TaxCode].TaxRate, netAmountMap[v.Product].NetAmount, taxAmount)
			data = append(data, *datum)
		} else {
			datum := psdc.ConvertToTaxAmount(v.Product, v.TaxCode, taxRateMap[v.TaxCode].TaxRate, netAmountMap[v.Product].NetAmount, itemMap[v.Product].TaxAmount)
			data = append(data, *datum)
			if math.Abs(float64(*taxAmount-*itemMap[v.Product].TaxAmount)) >= 2 {
				return nil, xerrors.Errorf("TaxAmountについて入力ファイルの値と計算結果の差の絶対値が2以上の明細が一つ以上存在します。")
			}
		}
	}

	return &data, nil
}

func CalculateTaxAmount(taxRate *float32, netAmount *float32) (*float32, error) {
	if taxRate == nil || netAmount == nil {
		return nil, xerrors.Errorf("TaxRateまたはNetAmountがnullです。")
	}

	digit := Float32DecimalDigit(*netAmount)
	mul := *netAmount * *taxRate / 100

	s := math.Round(float64(mul)*math.Pow10(digit)) / math.Pow10(digit)
	res := ParseFloat32Ptr(float32(s))

	return res, nil
}

func (f *SubFunction) GrossAmount(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.GrossAmount, error) {
	var data []api_processing_data_formatter.GrossAmount

	item := sdc.Orders.Item
	itemMap := make(map[string]api_input_reader.Item, len(item))
	for _, v := range item {
		itemMap[v.Product] = v
	}

	for _, v := range *psdc.TaxAmount {
		grossAmount := ParseFloat32Ptr(*v.NetAmount + *v.TaxAmount)

		if itemMap[v.Product].GrossAmount == nil {
			datum := psdc.ConvertToGrossAmount(v.Product, v.NetAmount, v.TaxAmount, grossAmount)
			data = append(data, *datum)
		} else {
			datum := psdc.ConvertToGrossAmount(v.Product, v.NetAmount, v.TaxAmount, itemMap[v.Product].GrossAmount)
			data = append(data, *datum)
			if math.Abs(float64(*grossAmount-*itemMap[v.Product].GrossAmount)) >= 2 {
				return nil, xerrors.Errorf("GrossAmountについて入力ファイルの値と計算結果の差の絶対値が2以上の明細が一つ以上存在します。")
			}
		}
	}

	return &data, nil
}
