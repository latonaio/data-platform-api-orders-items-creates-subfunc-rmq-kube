package dpfm_api_output_formatter

import (
	api_input_reader "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"encoding/json"
	"reflect"

	"golang.org/x/xerrors"
)

func ConvertToItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]Item, error) {
	var err error
	calculateOrderID := psdc.CalculateOrderID
	orderItem := psdc.OrderItem
	itemPMGeneral := psdc.ItemPMGeneral
	itemBPTaxClassification := psdc.ItemBPTaxClassification
	itemProductDescription := psdc.ItemProductDescription
	itemProductTaxClassification := psdc.ItemProductTaxClassification
	stockConfirmationPlant := psdc.StockConfirmationPlant
	paymentMethod := psdc.PaymentMethod

	productMasterBPPlantMap := StructArrayToMap(*psdc.ProductMasterBPPlant, "Product")
	itemNetWeightMap := StructArrayToMap(*psdc.ItemNetWeight, "Product")
	itemGrossWeightMap := StructArrayToMap(*psdc.ItemGrossWeight, "Product")
	netAmountMap := StructArrayToMap(*psdc.NetAmount, "Product")
	taxAmountMap := StructArrayToMap(*psdc.TaxAmount, "Product")
	grossAmountMap := StructArrayToMap(*psdc.GrossAmount, "Product")

	// productionPlantMap := StructArrayToMap(*psdc.ProductionPlant, "Product")

	res := make([]Item, 0, len(*itemPMGeneral))
	for i, v := range *itemPMGeneral {
		item := Item{}
		inputItem := sdc.Orders.Item[0]

		item, err = jsonTypeConversion[Item](inputItem)
		if err != nil {
			return nil, err
		}

		item, err = jsonTypeConversion[Item](v)
		if err != nil {
			return nil, err
		}

		item.OrderID = *calculateOrderID.OrderIDLatestNumber
		item.OrderItem = (*orderItem)[i].OrderItemNumber
		item.OrderItemText = *(*itemProductDescription)[i].ProductDescription
		item.BPTaxClassification = *itemBPTaxClassification.BPTaxClassification
		item.ProductTaxClassification = *(*itemProductTaxClassification)[i].ProductTaxClassification
		if v.ItemCategory != nil {
			if *v.ItemCategory == "INVP" {
				item.StockConfirmationPartnerFunction = &stockConfirmationPlant.PartnerFunction
				item.StockConfirmationBusinessPartner = &stockConfirmationPlant.BusinessPartner
				item.StockConfirmationPlant = stockConfirmationPlant.Plant

				item.OrderIssuingUnit = *productMasterBPPlantMap[v.Product].IssuingDeliveryUnit
				item.OrderReceivingUnit = *productMasterBPPlantMap[v.Product].ReceivingDeliveryUnit
				item.ProductIsBatchManagedInStockConfirmationPlant = *productMasterBPPlantMap[v.Product].IsBatchManagementRequired

				// item.ProductionPlantPartnerFunction = productionPlantMap[].PartnerFunction
				// item.ProductionPlantBusinessPartner = productionPlantMap[].BusinessPartner
				// item.ProductionPlant = productionPlantMap[].Plant
			}
		}
		item.PaymentMethod = *paymentMethod.PaymentMethod
		item.ItemNetWeight = itemNetWeightMap[v.Product].ItemNetWeight
		item.ItemGrossWeight = itemGrossWeightMap[v.Product].ItemGrossWeight
		item.NetAmount = netAmountMap[v.Product].NetAmount
		item.TaxAmount = taxAmountMap[v.Product].TaxAmount
		item.GrossAmount = grossAmountMap[v.Product].GrossAmount

		res = append(res, item)
	}

	return &res, nil
}

func ConvertToItemPricingElement(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]ItemPricingElement, error) {
	calculateOrderID := psdc.CalculateOrderID
	orderItem := psdc.OrderItem
	conditionAmount := psdc.ConditionAmount
	conditionAmountMap := make(map[string]api_processing_data_formatter.ConditionAmount, len(*conditionAmount))
	for _, v := range *conditionAmount {
		conditionAmountMap[v.Product] = v
	}

	priceMaster := psdc.PriceMaster

	res := make([]ItemPricingElement, 0, len(*priceMaster))

	for i, v := range *priceMaster {
		itemPricingElement := ItemPricingElement{}
		inputItemPricingElement := sdc.Orders.Item[0].ItemPricingElement[0]
		inputData, err := json.Marshal(inputItemPricingElement)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(inputData, &itemPricingElement)
		if err != nil {
			return nil, err
		}

		data, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &itemPricingElement)
		if err != nil {
			return nil, err
		}

		itemPricingElement.OrderID = *calculateOrderID.OrderIDLatestNumber
		itemPricingElement.OrderItem = (*orderItem)[i].OrderItemNumber
		itemPricingElement.ConditionQuantity = *conditionAmountMap[v.Product].ConditionQuantity
		itemPricingElement.ConditionAmount = *conditionAmountMap[v.Product].ConditionAmount
		itemPricingElement.ConditionIsManuallyChanged = *conditionAmountMap[v.Product].ConditionIsManuallyChanged

		res = append(res, itemPricingElement)
	}

	return &res, nil
}

func StructArrayToMap[T any](data []T, key string) map[any]T {
	res := make(map[any]T, len(data))

	for _, value := range data {
		m := StructToMap[T](&value, key)
		for k, v := range m {
			res[k] = v
		}
	}

	return res
}

func StructToMap[T any](data interface{}, key string) map[any]T {
	res := make(map[any]T)
	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()

	for i := 0; i < size; i++ {
		field := elem.Type().Field(i).Name
		value := elem.Field(i).Interface()
		if field == key {
			res[value], _ = jsonTypeConversion[T](elem.Interface())
		}
	}

	return res
}

func jsonTypeConversion[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
