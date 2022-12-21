package subfunction

import (
	api_input_reader "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"database/sql"
	"strings"
)

func (f *SubFunction) HeaderPartnerPlant(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.HeaderPartnerPlant, error) {
	var args []interface{}
	var rows *sql.Rows
	var err error

	buyerSellerDetection := psdc.BuyerSellerDetection
	headerPartnerFunction := psdc.HeaderPartnerFunction
	dataKey := psdc.ConvertToHeaderPartnerPlantKey(len(*headerPartnerFunction))
	if err != nil {
		return nil, err
	}

	for i, v := range *headerPartnerFunction {
		(*dataKey)[i].BusinessPartnerID = buyerSellerDetection.BusinessPartnerID
		if buyerSellerDetection.BuyerOrSeller == "Seller" {
			(*dataKey)[i].CustomerOrSupplier = buyerSellerDetection.Buyer
		} else if buyerSellerDetection.BuyerOrSeller == "Buyer" {
			(*dataKey)[i].CustomerOrSupplier = buyerSellerDetection.Seller
		}
		(*dataKey)[i].PartnerCounter = v.PartnerCounter
		(*dataKey)[i].PartnerFunction = v.PartnerFunction
		(*dataKey)[i].PartnerFunctionBusinessPartner = v.BusinessPartner
	}

	repeat := strings.Repeat("(?,?,?,?,?),", len(*dataKey)-1) + "(?,?,?,?,?)"
	for _, tag := range *dataKey {
		args = append(
			args,
			tag.BusinessPartnerID,
			tag.CustomerOrSupplier,
			tag.PartnerCounter,
			tag.PartnerFunction,
			tag.PartnerFunctionBusinessPartner)
	}

	if buyerSellerDetection.BuyerOrSeller == "Seller" {
		rows, err = f.db.Query(
			`SELECT PartnerFunctionBusinessPartner, PartnerFunction, PlantCounter, Plant, DefaultPlant, DefaultStockConfirmationPlant
				FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_customer_partner_plant_data
				WHERE (BusinessPartner, Customer, PartnerCounter, PartnerFunction, PartnerFunctionBusinessPartner) IN ( `+repeat+` );`, args...,
		)
		if err != nil {
			return nil, err
		}
	} else if buyerSellerDetection.BuyerOrSeller == "Buyer" {
		rows, err = f.db.Query(
			`SELECT PartnerFunctionBusinessPartner, PartnerFunction, PlantCounter, Plant, DefaultPlant, DefaultStockConfirmationPlant
				FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_supplier_partner_plant_data
				WHERE (BusinessPartner, Supplier, PartnerCounter, PartnerFunction, PartnerFunctionBusinessPartner) IN ( `+repeat+` );`, args...,
		)
		if err != nil {
			return nil, err
		}
	}

	data, err := psdc.ConvertToHeaderPartnerPlant(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}
