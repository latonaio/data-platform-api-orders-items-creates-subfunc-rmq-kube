# data-platform-api-orders-items-creates-subfunc-rmq-kube
data-platform-api-orders-items-creates-subfunc-rmq-kube は、データ連携基盤において、オーダーAPIサービスのヘッダ登録補助機能を担うマイクロサービスです。

## 動作環境
・ OS: LinuxOS  
・ CPU: ARM/AMD/Intel  

## 対象APIサービス
data-platform-api-orders-items-creates-subfunc-rmq-kube の 対象APIサービスは次の通りです。

*  APIサービス URL: https://xxx.xxx.io/api/API_ORDERS_SRV/creates/

## 本レポジトリ が 対応する データ
data-platform-api-orders-items-creates-subfunc-rmq-kube が対応する データ は、次のものです。

* OrdersHeader（オーダー - ヘッダデータ）
* OrdersHeaderPartner（オーダー - ヘッダ取引先データ）
* OrdersHeaderPartnerPlant（オーダー - ヘッダ取引先プラントデータ）
* OrdersHeaderPartnerContact（オーダー - ヘッダ取引先コンタクトデータ）

## Output
data-platform-api-orders-items-creates-subfunc-rmq-kube では、[golang-logging-library-for-data-platform](https://github.com/latonaio/golang-logging-library-for-data-platform) により、Output として、RabbitMQ へのメッセージを JSON 形式で出力します。以下の項目のうち、"cursor" ～ "time"は、golang-logging-library-for-data-platform による 定型フォーマットの出力結果です。

```
{
  "cursor": "/Users/tatsuya/dev/data-platform-api-orders-creates-subfunc-rmq-kube/main.go#L67",
  "function": "main.callProcess",
  "level": "INFO",
  "message": {
    "connection_key": "request",
    "result": true,
    "redis_key": "abcdefg",
    "filepath": "/var/lib/aion/Data/rededge_sdc/abcdef.json",
    "runtime_session_id": "boi9ar543dg91ipdnspi099u231280ab0v8af0ew",
    "business_partner": 201,
    "service_label": "ORDERS",
    "Orders": {
      "OrderID": 11,
      "OrderDate": "",
      "OrderType": "",
      "Buyer": 101,
      "Seller": 201,
      "CreationDate": "",
      "LastChangeDate": "",
      "ContractType": "",
      "ValidityStartDate": "",
      "ValidityEndDate": "",
      "InvoiceScheduleStartDate": "",
      "InvoiceScheduleEndDate": "",
      "TotalNetAmount": null,
      "TotalTaxAmount": null,
      "TotalGrossAmount": null,
      "OverallDeliveryStatus": "",
      "TotalBlockStatus": null,
      "OverallOrdReltdBillgStatus": "",
      "OverallDocReferenceStatus": "",
      "TransactionCurrency": "",
      "PricingDate": "",
      "PriceDetnExchangeRate": null,
      "RequestedDeliveryDate": "",
      "HeaderCompleteDeliveryIsDefined": null,
      "HeaderBillingBlockReason": null,
      "DeliveryBlockReason": null,
      "Incoterms": "CIF",
      "PaymentTerms": "0001",
      "PaymentMethod": "T",
      "ReferenceDocument": null,
      "ReferenceDocumentItem": null,
      "BPAccountAssignmentGroup": "01",
      "AccountingExchangeRate": null,
      "BillingDocumentDate": "",
      "IsExportImportDelivery": null,
      "HeaderText": "",
      "HeaderPartner": [
        {
          "PartnerFunction": "DELIVERTO",
          "BusinessPartner": 102,
          "BusinessPartnerFullName": "株式会社ABC虎ノ門店",
          "BusinessPartnerName": "ABC虎ノ門店",
          "Organization": "",
          "Country": "JP",
          "Language": "JA",
          "Currency": "JPY",
          "ExternalDocumentID": "",
          "AddressID": 200000,
          "HeaderPartnerContact": null,
          "HeaderPartnerPlant": [
            {
              "Plant": "AB02"
            }
          ]
        },
        {
          "PartnerFunction": "BUYER",
          "BusinessPartner": 101,
          "BusinessPartnerFullName": "株式会社ABC本社",
          "BusinessPartnerName": "ABC本社",
          "Organization": "",
          "Country": "JP",
          "Language": "JA",
          "Currency": "JPY",
          "ExternalDocumentID": "",
          "AddressID": 100000,
          "HeaderPartnerContact": [
            {
              "ContactID": null,
              "ContactPersonName": "",
              "EmailAddress": "",
              "PhoneNumber": "",
              "MobilePhoneNumber": "",
              "FaxNumber": "",
              "ContactTag1": "",
              "ContactTag2": "",
              "ContactTag3": "",
              "ContactTag4": ""
            }
          ],
          "HeaderPartnerPlant": [
            {
              "Plant": "AB01"
            }
          ]
        },
        {
          "PartnerFunction": "SELLER",
          "BusinessPartner": 201,
          "BusinessPartnerFullName": "パン販売株式会社",
          "BusinessPartnerName": "パン販売",
          "Organization": "",
          "Country": "JP",
          "Language": "JA",
          "Currency": "JPY",
          "ExternalDocumentID": "",
          "AddressID": 300000,
          "HeaderPartnerContact": [
            {
              "ContactID": null,
              "ContactPersonName": "",
              "EmailAddress": "",
              "PhoneNumber": "",
              "MobilePhoneNumber": "",
              "FaxNumber": "",
              "ContactTag1": "",
              "ContactTag2": "",
              "ContactTag3": "",
              "ContactTag4": ""
            }
          ],
          "HeaderPartnerPlant": [
            {
              "Plant": "TE01"
            }
          ]
        }
      ],
      "Address": [
        {
          "AddressID": null,
          "PostalCode": "",
          "LocalRegion": "",
          "Country": "",
          "District": "",
          "StreetName": "",
          "CityName": "",
          "Building": "",
          "Floor": null,
          "Room": null
        }
      ],
      "HeaderPDF": [
        {
          "DocType": "",
          "DocVersionID": null,
          "DocID": "",
          "DocIssuerBusinessPartner": null,
          "FileName": ""
        }
      ],
      "Item": [
        {
          "OrderItem": 1,
          "OrderItemCategory": "",
          "OrderItemText": "RobotA",
          "Product": "A3750",
          "ProductStandardID": "CDC",
          "ProductGroup": "01",
          "BaseUnit": "PC",
          "PricingDate": "",
          "PriceDetnExchangeRate": null,
          "RequestedDeliveryDate": "",
          "StockConfirmationPartnerFunction": "",
          "StockConfirmationBusinessPartner": null,
          "StockConfirmationPlant": "",
          "StockConfirmationPlantBatch": "",
          "StockConfirmationPlantBatchValidityStartDate": "",
          "StockConfirmationPlantBatchValidityEndDate": "",
          "ProductIsBatchManagedInStockConfirmationPlant": null,
          "OrderQuantityInBaseUnit": null,
          "OrderQuantityInIssuingUnit": null,
          "OrderQuantityInReceivingUnit": null,
          "OrderIssuingUnit": "",
          "OrderReceivingUnit": "",
          "StockConfirmationPolicy": "",
          "StockConfirmationStatus": "",
          "ConfdDelivQtyInOrderQtyUnit": null,
          "ItemWeightUnit": "G",
          "ProductGrossWeight": 300.05,
          "ItemGrossWeight": null,
          "ProductNetWeight": 300.05,
          "ItemNetWeight": null,
          "NetAmount": null,
          "TaxAmount": null,
          "GrossAmount": null,
          "BillingDocumentDate": "",
          "ProductionPlantPartnerFunction": "",
          "ProductionPlantBusinessPartner": null,
          "ProductionPlant": "",
          "ProductionPlantTimeZone": "",
          "ProductionPlantStorageLocation": "",
          "IssuingPlantPartnerFunction": "",
          "IssuingPlantBusinessPartner": null,
          "IssuingPlant": "",
          "IssuingPlantTimeZone": "",
          "IssuingPlantStorageLocation": "",
          "ReceivingPlantPartnerFunction": "",
          "ReceivingPlantBusinessPartner": null,
          "ReceivingPlant": "",
          "ReceivingPlantTimeZone": "",
          "ReceivingPlantStorageLocation": "",
          "ProductIsBatchManagedInProductionPlant": null,
          "ProductIsBatchManagedInIssuingPlant": null,
          "ProductIsBatchManagedInReceivingPlant": null,
          "BatchMgmtPolicyInProductionPlant": "",
          "BatchMgmtPolicyInIssuingPlant": "",
          "BatchMgmtPolicyInReceivingPlant": "",
          "ProductionPlantBatch": "",
          "IssuingPlantBatch": "",
          "ReceivingPlantBatch": "",
          "ProductionPlantBatchValidityStartDate": "",
          "ProductionPlantBatchValidityEndDate": "",
          "IssuingPlantBatchValidityStartDate": "",
          "IssuingPlantBatchValidityEndDate": "",
          "ReceivingPlantBatchValidityStartDate": "",
          "ReceivingPlantBatchValidityEndDate": "",
          "Incoterms": "",
          "BPTaxClassification": "1",
          "ProductTaxClassification": "",
          "BPAccountAssignmentGroup": "",
          "ProductAccountAssignmentGroup": "01",
          "PaymentTerms": "",
          "PaymentMethod": "",
          "DocumentRjcnReason": null,
          "ItemBillingBlockReason": null,
          "Project": "",
          "AccountingExchangeRate": null,
          "ReferenceDocument": null,
          "ReferenceDocumentItem": null,
          "ItemCompleteDeliveryIsDefined": null,
          "ItemDeliveryStatus": "",
          "IssuingStatus": "",
          "ReceivingStatus": "",
          "BillingStatus": "",
          "TaxCode": "",
          "TaxRate": null,
          "CountryOfOrigin": "JPY",
          "ItemPartner": [
            {
              "PartnerFunction": "",
              "BusinessPartner": null,
              "ItemPartnerPlant": {
                "Plant": ""
              }
            }
          ],
          "ItemPricingElement": [
            {
              "PricingProcedureStep": null,
              "PricingProcedureCounter": null,
              "ConditionType": "",
              "PricingDate": "",
              "ConditionRateValue": null,
              "ConditionCurrency": "",
              "ConditionQuantity": null,
              "ConditionQuantityUnit": "",
              "ConditionRecord": null,
              "ConditionSequentialNumber": null,
              "TaxCode": "",
              "ConditionAmount": null,
              "TransactionCurrency": "",
              "ConditionIsManuallyChanged": null
            }
          ],
          "ItemSchedulingLine": [
            {
              "ScheduleLine": null,
              "Product": "",
              "StockConfirmationPartnerFunction": "",
              "StockConfirmationBusinessPartner": null,
              "StockConfirmationPlant": "",
              "StockConfirmationPlantBatch": "",
              "StockConfirmationPlantBatchValidityStartDate": "",
              "StockConfirmationPlantBatchValidityEndDate": "",
              "ConfirmedDeliveryDate": "",
              "RequestedDeliveryDate": "",
              "OrderQuantityInBaseUnit": null,
              "ConfdOrderQtyByPDTAvailCheck": null,
              "DeliveredQtyInOrderQtyUnit": null,
              "OpenConfdDelivQtyInOrdQtyUnit": null,
              "DelivBlockReasonForSchedLine": null,
              "PlusMinusFlag": ""
            }
          ]
        },
        {
          "OrderItem": 2,
          "OrderItemCategory": "",
          "OrderItemText": "RobotB",
          "Product": "W999",
          "ProductStandardID": "CAC",
          "ProductGroup": "01",
          "BaseUnit": "PC",
          "PricingDate": "",
          "PriceDetnExchangeRate": null,
          "RequestedDeliveryDate": "",
          "StockConfirmationPartnerFunction": "",
          "StockConfirmationBusinessPartner": null,
          "StockConfirmationPlant": "",
          "StockConfirmationPlantBatch": "",
          "StockConfirmationPlantBatchValidityStartDate": "",
          "StockConfirmationPlantBatchValidityEndDate": "",
          "ProductIsBatchManagedInStockConfirmationPlant": null,
          "OrderQuantityInBaseUnit": null,
          "OrderQuantityInIssuingUnit": null,
          "OrderQuantityInReceivingUnit": null,
          "OrderIssuingUnit": "",
          "OrderReceivingUnit": "",
          "StockConfirmationPolicy": "",
          "StockConfirmationStatus": "",
          "ConfdDelivQtyInOrderQtyUnit": null,
          "ItemWeightUnit": "G",
          "ProductGrossWeight": 100.05,
          "ItemGrossWeight": null,
          "ProductNetWeight": 100.05,
          "ItemNetWeight": null,
          "NetAmount": null,
          "TaxAmount": null,
          "GrossAmount": null,
          "BillingDocumentDate": "",
          "ProductionPlantPartnerFunction": "",
          "ProductionPlantBusinessPartner": null,
          "ProductionPlant": "",
          "ProductionPlantTimeZone": "",
          "ProductionPlantStorageLocation": "",
          "IssuingPlantPartnerFunction": "",
          "IssuingPlantBusinessPartner": null,
          "IssuingPlant": "",
          "IssuingPlantTimeZone": "",
          "IssuingPlantStorageLocation": "",
          "ReceivingPlantPartnerFunction": "",
          "ReceivingPlantBusinessPartner": null,
          "ReceivingPlant": "",
          "ReceivingPlantTimeZone": "",
          "ReceivingPlantStorageLocation": "",
          "ProductIsBatchManagedInProductionPlant": null,
          "ProductIsBatchManagedInIssuingPlant": null,
          "ProductIsBatchManagedInReceivingPlant": null,
          "BatchMgmtPolicyInProductionPlant": "",
          "BatchMgmtPolicyInIssuingPlant": "",
          "BatchMgmtPolicyInReceivingPlant": "",
          "ProductionPlantBatch": "",
          "IssuingPlantBatch": "",
          "ReceivingPlantBatch": "",
          "ProductionPlantBatchValidityStartDate": "",
          "ProductionPlantBatchValidityEndDate": "",
          "IssuingPlantBatchValidityStartDate": "",
          "IssuingPlantBatchValidityEndDate": "",
          "ReceivingPlantBatchValidityStartDate": "",
          "ReceivingPlantBatchValidityEndDate": "",
          "Incoterms": "",
          "BPTaxClassification": "1",
          "ProductTaxClassification": "",
          "BPAccountAssignmentGroup": "",
          "ProductAccountAssignmentGroup": "01",
          "PaymentTerms": "",
          "PaymentMethod": "",
          "DocumentRjcnReason": null,
          "ItemBillingBlockReason": null,
          "Project": "",
          "AccountingExchangeRate": null,
          "ReferenceDocument": null,
          "ReferenceDocumentItem": null,
          "ItemCompleteDeliveryIsDefined": null,
          "ItemDeliveryStatus": "",
          "IssuingStatus": "",
          "ReceivingStatus": "",
          "BillingStatus": "",
          "TaxCode": "",
          "TaxRate": null,
          "CountryOfOrigin": "JPY",
          "ItemPartner": [
            {
              "PartnerFunction": "",
              "BusinessPartner": null,
              "ItemPartnerPlant": {
                "Plant": ""
              }
            }
          ],
          "ItemPricingElement": [
            {
              "PricingProcedureStep": null,
              "PricingProcedureCounter": null,
              "ConditionType": "",
              "PricingDate": "",
              "ConditionRateValue": null,
              "ConditionCurrency": "",
              "ConditionQuantity": null,
              "ConditionQuantityUnit": "",
              "ConditionRecord": null,
              "ConditionSequentialNumber": null,
              "TaxCode": "",
              "ConditionAmount": null,
              "TransactionCurrency": "",
              "ConditionIsManuallyChanged": null
            }
          ],
          "ItemSchedulingLine": [
            {
              "ScheduleLine": null,
              "Product": "",
              "StockConfirmationPartnerFunction": "",
              "StockConfirmationBusinessPartner": null,
              "StockConfirmationPlant": "",
              "StockConfirmationPlantBatch": "",
              "StockConfirmationPlantBatchValidityStartDate": "",
              "StockConfirmationPlantBatchValidityEndDate": "",
              "ConfirmedDeliveryDate": "",
              "RequestedDeliveryDate": "",
              "OrderQuantityInBaseUnit": null,
              "ConfdOrderQtyByPDTAvailCheck": null,
              "DeliveredQtyInOrderQtyUnit": null,
              "OpenConfdDelivQtyInOrdQtyUnit": null,
              "DelivBlockReasonForSchedLine": null,
              "PlusMinusFlag": ""
            }
          ]
        }
      ]
    },
    "api_schema": "DPFMOrdersCreates",
    "accepter": [
      "All"
    ],
    "order_id": null,
    "deleted": false
  },
  "runtime_session_id": "boi9ar543dg91ipdnspi099u231280ab0v8af0ew",
  "time": "2022-11-08T11:21:23+09:00"
}
```