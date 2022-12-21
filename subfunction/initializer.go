package subfunction

import (
	"context"
	api_input_reader "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Output_Formatter"
	api_processing_data_formatter "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"sync"

	database "github.com/latonaio/golang-mysql-network-connector"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type SubFunction struct {
	ctx context.Context
	db  *database.Mysql
	l   *logger.Logger
}

func NewSubFunction(ctx context.Context, db *database.Mysql, l *logger.Logger) *SubFunction {
	return &SubFunction{
		ctx: ctx,
		db:  db,
		l:   l,
	}
}

func (f *SubFunction) MetaData(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.MetaData {
	metaData := psdc.ConvertToMetaData(sdc)

	return metaData
}

func (f *SubFunction) CreateSdc(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) error {
	var err error
	var e error

	wg := sync.WaitGroup{}
	wg.Add(5)

	psdc.MetaData = f.MetaData(sdc, psdc)

	// 1-0. 入力ファイルのbusiness_partnerがBuyerであるかSellerであるかの判断
	psdc.BuyerSellerDetection, err = f.BuyerSellerDetection(sdc, psdc)
	if err != nil {
		return err
	}

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-1. ビジネスパートナ 得意先データ/仕入先データ の取得
		psdc.HeaderBPCustomerSupplier, e = f.HeaderBPCustomerSupplier(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 5-17. PaymentMethod  // 1-1
		psdc.PaymentMethod = f.PaymentMethod(sdc, psdc)
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-2. OrderID
		psdc.CalculateOrderID, e = f.CalculateOrderID(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 5-0. OrderItem
		psdc.OrderItem = f.OrderItem(sdc, psdc)
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 5-2. 品目マスタ一般データの取得
		psdc.ItemPMGeneral, e = f.ItemPMGeneral(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 5-3 OrderItemText  // 5-2
		psdc.ItemProductDescription, e = f.ItemProductDescription(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 5-19. ItemNetWeight  // 5-2
		psdc.ItemNetWeight = f.ItemNetWeight(sdc, psdc)

		// 5-18. ItemGrossWeight  // 5-2
		psdc.ItemGrossWeight = f.ItemGrossWeight(sdc, psdc)
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 2-1. ビジネスパートナマスタの取引先機能データの取得
		psdc.HeaderPartnerFunction, e = f.HeaderPartnerFunction(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 2-2. ビジネスパートナの一般データの取得  // 2-1
		psdc.HeaderPartnerBPGeneral, e = f.HeaderPartnerBPGeneral(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 5-1-1. BPTaxClassification  // 2-2
		psdc.ItemBPTaxClassification, e = f.ItemBPTaxClassification(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 5-1-2 ProductTaxClassification
		psdc.ItemProductTaxClassification, e = f.ItemProductTaxClassification(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 5-1-3 TaxCode  // 5-1-1, 5-1-2
		psdc.TaxCode, e = f.TaxCode(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 4-1. ビジネスパートナマスタの取引先プラントデータの取得  // 2-1
		psdc.HeaderPartnerPlant, e = f.HeaderPartnerPlant(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 5-4-1 StockConfirmationPlant  // 4-1
		psdc.StockConfirmationPlant = f.StockConfirmationPlant(sdc, psdc)

		// TODO: 仕様変更に対応する必要あり
		// 5-5-1 ProductMasterBPPlant  // 4-1
		psdc.ProductMasterBPPlant, e = f.ProductMasterBPPlant(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 5-6-1 GetPartnerPlantData  // 4-1
		psdc.ProductionPlant = f.ProductionPlant(sdc, psdc)

		// 1-8. PricingDate
		psdc.PricingDate = f.PricingDate(sdc, psdc)

		// 8-1. 価格マスタデータの取得(入力ファイルの[ConditionAmount]がnullである場合)  // 1-8
		psdc.PriceMaster, e = f.PriceMaster(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 8-2. 価格の計算(入力ファイルの[ConditionAmount]がnullである場合)  // 8-1
		psdc.ConditionAmount, e = f.ConditionAmount(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 5-21. NetAmount  // 8-2
		psdc.NetAmount = f.NetAmount(sdc, psdc)

		// 5-20. TaxRateの計算  // 5-1-3
		psdc.TaxRate, e = f.TaxRate(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 5-22. TaxAmount  // 5-1-3, 5-20, 5-21
		psdc.TaxAmount, e = f.TaxAmount(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 5-23. GrossAmount  // 5-21, 5-22
		psdc.GrossAmount, e = f.GrossAmount(sdc, psdc)
		if e != nil {
			err = e
			return
		}

	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	f.l.Info(psdc)
	osdc, err = f.SetValue(sdc, psdc, osdc)
	if err != nil {
		return err
	}

	return nil
}
