package subfunction

import (
	api_input_reader "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"database/sql"
	"math"
	"strconv"
	"strings"

	"golang.org/x/xerrors"
)

func (f *SubFunction) PriceMaster(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.PriceMaster, error) {
	var args []interface{}
	var rows *sql.Rows
	var err error

	buyerSellerDetection := psdc.BuyerSellerDetection
	dataKey := psdc.ConvertToPriceMasterKey(sdc)

	for _, v := range sdc.Orders.Item {
		if v.ItemPricingElement[0].ConditionAmount == nil {
			dataKey.Product = append(dataKey.Product, v.Product)
		}
	}

	dataKey.ConditionValidityEndDate = psdc.PricingDate.PricingDate
	dataKey.ConditionValidityStartDate = psdc.PricingDate.PricingDate

	repeat := strings.Repeat("?,", len(dataKey.Product)-1) + "?"
	for _, v := range dataKey.Product {
		args = append(args, v)
	}

	if buyerSellerDetection.BuyerOrSeller == "Seller" {
		dataKey.CustomerOrSupplier = buyerSellerDetection.Buyer
		args = append(args, dataKey.BusinessPartnerID, dataKey.CustomerOrSupplier, dataKey.ConditionValidityEndDate, dataKey.ConditionValidityStartDate)
		rows, err = f.db.Query(
			`SELECT BusinessPartner, Product, Customer, ConditionValidityEndDate, ConditionValidityEndDate,
			ConditionRecord, ConditionSequentialNumber, ConditionType, ConditionRateValue
			FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_price_master_price_master_data
			WHERE Product IN ( `+repeat+` )
			AND (BusinessPartner, Customer) = (?, ?)
			AND ConditionValidityEndDate >= ?
			AND ConditionValidityStartDate <= ?;`, args...,
		)
		if err != nil {
			return nil, err
		}
	} else if buyerSellerDetection.BuyerOrSeller == "Buyer" {
		dataKey.CustomerOrSupplier = buyerSellerDetection.Seller
		args = append(args, dataKey.BusinessPartnerID, dataKey.CustomerOrSupplier, dataKey.ConditionValidityEndDate, dataKey.ConditionValidityStartDate)
		rows, err = f.db.Query(
			`SELECT BusinessPartner, Product, Supplier, ConditionValidityEndDate, ConditionValidityEndDate,
			ConditionRecord, ConditionSequentialNumber, ConditionType, ConditionRateValue
			FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_price_master_price_master_data
			WHERE Product IN ( `+repeat+` )
			AND (BusinessPartner, Supplier) = (?, ?)
			AND ConditionValidityEndDate >= ?
			AND ConditionValidityStartDate <= ?;`, args...,
		)
		if err != nil {
			return nil, err
		}
	}

	data, err := psdc.ConvertToPriceMaster(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) ConditionAmount(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.ConditionAmount, error) {
	var data []api_processing_data_formatter.ConditionAmount

	priceMaster := psdc.PriceMaster
	priceMasterMap := make(map[string]api_processing_data_formatter.PriceMaster, len(*priceMaster))
	for _, v := range *priceMaster {
		priceMasterMap[v.Product] = v
	}

	for _, v := range sdc.Orders.Item {
		if v.ItemPricingElement[0].ConditionAmount == nil {
			product := v.Product
			conditionQuantity := v.OrderQuantityInBaseUnit
			conditionRateValue := priceMasterMap[v.Product].ConditionRateValue
			conditionAmount, err := CalculateConditionAmount(conditionQuantity, conditionRateValue)
			if err != nil {
				return nil, err
			}

			datum := psdc.ConvertToConditionAmount(product, conditionQuantity, conditionAmount)
			data = append(data, *datum)
		}
	}

	return &data, nil
}

func Float32DecimalDigit(f float32) int {
	s := strconv.FormatFloat(float64(f), 'f', -1, 32)

	i := strings.Index(s, ".")
	if i == -1 {
		return 0
	}

	return len(s[i+1:])
}

func ParseFloat32Ptr(f float32) *float32 {
	return &f
}

func CalculateConditionAmount(conditionRateValue *float32, conditionQuantity *float32) (*float32, error) {
	if conditionRateValue == nil || conditionQuantity == nil {
		return nil, xerrors.Errorf("ConditionRateValueまたはConditionQuantityがnullです。")
	}

	digit := Float32DecimalDigit(*conditionRateValue)
	mul := *conditionRateValue * *conditionQuantity

	s := math.Round(float64(mul)*math.Pow10(digit)) / math.Pow10(digit)
	res := ParseFloat32Ptr(float32(s))

	return res, nil
}
