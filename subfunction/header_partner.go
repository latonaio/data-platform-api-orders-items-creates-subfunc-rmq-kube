package subfunction

import (
	api_input_reader "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"database/sql"
	"strings"
)

func (f *SubFunction) HeaderPartnerFunction(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.HeaderPartnerFunction, error) {
	var rows *sql.Rows
	var err error

	buyerSellerDetection := psdc.BuyerSellerDetection
	dataKey := psdc.ConvertToHeaderPartnerFunctionKey()

	dataKey.BusinessPartnerID = buyerSellerDetection.BusinessPartnerID

	if buyerSellerDetection.BuyerOrSeller == "Seller" {
		dataKey.CustomerOrSupplier = buyerSellerDetection.Buyer
		rows, err = f.db.Query(
			`SELECT BusinessPartner, PartnerCounter, PartnerFunction, PartnerFunctionBusinessPartner, DefaultPartner
			FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_customer_partner_function_data
			WHERE (BusinessPartner, Customer) = (?, ?);`, dataKey.BusinessPartnerID, dataKey.CustomerOrSupplier,
		)
		if err != nil {
			return nil, err
		}
	} else if buyerSellerDetection.BuyerOrSeller == "Buyer" {
		dataKey.CustomerOrSupplier = buyerSellerDetection.Seller
		rows, err = f.db.Query(
			`SELECT BusinessPartner, PartnerCounter, PartnerFunction, PartnerFunctionBusinessPartner, DefaultPartner
			FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_supplier_partner_function_data
			WHERE (BusinessPartner, Supplier) = (?, ?);`, dataKey.BusinessPartnerID, dataKey.CustomerOrSupplier,
		)
		if err != nil {
			return nil, err
		}
	}

	data, err := psdc.ConvertToHeaderPartnerFunction(rows)
	if err != nil {
		return nil, err
	}

	return data, err

}

func (f *SubFunction) HeaderPartnerBPGeneral(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.HeaderPartnerBPGeneral, error) {
	var args []interface{}

	headerPartnerFunction := psdc.HeaderPartnerFunction
	repeat := strings.Repeat("?,", len(*headerPartnerFunction)-1) + "?"
	for _, tag := range *headerPartnerFunction {
		args = append(args, tag.BusinessPartner)
	}

	rows, err := f.db.Query(
		`SELECT BusinessPartner, BusinessPartnerFullName, BusinessPartnerName, Country, Language, AddressID
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_general_data
		WHERE BusinessPartner IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToHeaderPartnerBPGeneral(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}
