package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-material-document-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
			MaterialDocumentYear:       data.MaterialDocumentYear,
			MaterialDocument:           data.MaterialDocument,
			InventoryTransactionType:   data.InventoryTransactionType,
			DocumentDate:               data.DocumentDate,
			PostingDate:                data.PostingDate,
			CreationDate:               data.CreationDate,
			CreationTime:               data.CreationTime,
			MaterialDocumentHeaderText: data.MaterialDocumentHeaderText,
			ReferenceDocument:          data.ReferenceDocument,
			GoodsMovementCode:          data.GoodsMovementCode,
			ToItem:                     data.ToItem.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) ([]Item, error) {
	pm := &responses.Item{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	item := make([]Item, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		item = append(item, Item{
			MaterialDocumentYear:           data.MaterialDocumentYear,
			MaterialDocument:               data.MaterialDocument,
			MaterialDocumentItem:           data.MaterialDocumentItem,
			Material:                       data.Material,
			Plant:                          data.Plant,
			StorageLocation:                data.StorageLocation,
			Batch:                          data.Batch,
			GoodsMovementType:              data.GoodsMovementType,
			InventoryStockType:             data.InventoryStockType,
			InventoryValuationType:         data.InventoryValuationType,
			InventorySpecialStockType:      data.InventorySpecialStockType,
			Supplier:                       data.Supplier,
			Customer:                       data.Customer,
			SalesOrder:                     data.SalesOrder,
			SalesOrderItem:                 data.SalesOrderItem,
			SalesOrderScheduleLine:         data.SalesOrderScheduleLine,
			PurchaseOrder:                  data.PurchaseOrder,
			PurchaseOrderItem:              data.PurchaseOrderItem,
			WBSElement:                     data.WBSElement,
			ManufacturingOrder:             data.ManufacturingOrder,
			ManufacturingOrderItem:         data.ManufacturingOrderItem,
			GoodsMovementRefDocType:        data.GoodsMovementRefDocType,
			GoodsMovementReasonCode:        data.GoodsMovementReasonCode,
			Delivery:                       data.Delivery,
			DeliveryItem:                   data.DeliveryItem,
			AccountAssignmentCategory:      data.AccountAssignmentCategory,
			CostCenter:                     data.CostCenter,
			ControllingArea:                data.ControllingArea,
			CostObject:                     data.CostObject,
			ProfitabilitySegment:           data.ProfitabilitySegment,
			ProfitCenter:                   data.ProfitCenter,
			GLAccount:                      data.GLAccount,
			FunctionalArea:                 data.FunctionalArea,
			MaterialBaseUnit:               data.MaterialBaseUnit,
			QuantityInBaseUnit:             data.QuantityInBaseUnit,
			EntryUnit:                      data.EntryUnit,
			QuantityInEntryUnit:            data.QuantityInEntryUnit,
			CompanyCodeCurrency:            data.CompanyCodeCurrency,
			FiscalYear:                     data.FiscalYear,
			FiscalYearPeriod:               data.FiscalYearPeriod,
			IssgOrRcvgMaterial:             data.IssgOrRcvgMaterial,
			IssgOrRcvgBatch:                data.IssgOrRcvgBatch,
			IssuingOrReceivingPlant:        data.IssuingOrReceivingPlant,
			IssuingOrReceivingStorageLoc:   data.IssuingOrReceivingStorageLoc,
			IssuingOrReceivingStockType:    data.IssuingOrReceivingStockType,
			IssgOrRcvgSpclStockInd:         data.IssgOrRcvgSpclStockInd,
			IssuingOrReceivingValType:      data.IssuingOrReceivingValType,
			IsCompletelyDelivered:          data.IsCompletelyDelivered,
			MaterialDocumentItemText:       data.MaterialDocumentItemText,
			UnloadingPointName:             data.UnloadingPointName,
			ShelfLifeExpirationDate:        data.ShelfLifeExpirationDate,
			ManufactureDate:                data.ManufactureDate,
			SerialNumbersAreCreatedAutomly: data.SerialNumbersAreCreatedAutomly,
			Reservation:                    data.Reservation,
			ReservationItem:                data.ReservationItem,
			ReservationIsFinallyIssued:     data.ReservationIsFinallyIssued,
			IsAutomaticallyCreated:         data.IsAutomaticallyCreated,
			GoodsMovementIsCancelled:       data.GoodsMovementIsCancelled,
			ReversedMaterialDocumentYear:   data.ReversedMaterialDocumentYear,
			ReversedMaterialDocument:       data.ReversedMaterialDocument,
			ReversedMaterialDocumentItem:   data.ReversedMaterialDocumentItem,
			ReferenceDocumentFiscalYear:    data.ReferenceDocumentFiscalYear,
			InvtryMgmtRefDocumentItem:      data.InvtryMgmtRefDocumentItem,
			InvtryMgmtReferenceDocument:    data.InvtryMgmtReferenceDocument,
			MaterialDocumentPostingType:    data.MaterialDocumentPostingType,
		})
	}

	return item, nil
}

func ConvertToToItem(raw []byte, l *logger.Logger) ([]ToItem, error) {
	pm := &responses.ToItem{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItem. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toItem := make([]ToItem, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItem = append(toItem, ToItem{
			MaterialDocumentYear:           data.MaterialDocumentYear,
			MaterialDocument:               data.MaterialDocument,
			MaterialDocumentItem:           data.MaterialDocumentItem,
			Material:                       data.Material,
			Plant:                          data.Plant,
			StorageLocation:                data.StorageLocation,
			Batch:                          data.Batch,
			GoodsMovementType:              data.GoodsMovementType,
			InventoryStockType:             data.InventoryStockType,
			InventoryValuationType:         data.InventoryValuationType,
			InventorySpecialStockType:      data.InventorySpecialStockType,
			Supplier:                       data.Supplier,
			Customer:                       data.Customer,
			SalesOrder:                     data.SalesOrder,
			SalesOrderItem:                 data.SalesOrderItem,
			SalesOrderScheduleLine:         data.SalesOrderScheduleLine,
			PurchaseOrder:                  data.PurchaseOrder,
			PurchaseOrderItem:              data.PurchaseOrderItem,
			WBSElement:                     data.WBSElement,
			ManufacturingOrder:             data.ManufacturingOrder,
			ManufacturingOrderItem:         data.ManufacturingOrderItem,
			GoodsMovementRefDocType:        data.GoodsMovementRefDocType,
			GoodsMovementReasonCode:        data.GoodsMovementReasonCode,
			Delivery:                       data.Delivery,
			DeliveryItem:                   data.DeliveryItem,
			AccountAssignmentCategory:      data.AccountAssignmentCategory,
			CostCenter:                     data.CostCenter,
			ControllingArea:                data.ControllingArea,
			CostObject:                     data.CostObject,
			ProfitabilitySegment:           data.ProfitabilitySegment,
			ProfitCenter:                   data.ProfitCenter,
			GLAccount:                      data.GLAccount,
			FunctionalArea:                 data.FunctionalArea,
			MaterialBaseUnit:               data.MaterialBaseUnit,
			QuantityInBaseUnit:             data.QuantityInBaseUnit,
			EntryUnit:                      data.EntryUnit,
			QuantityInEntryUnit:            data.QuantityInEntryUnit,
			CompanyCodeCurrency:            data.CompanyCodeCurrency,
			FiscalYear:                     data.FiscalYear,
			FiscalYearPeriod:               data.FiscalYearPeriod,
			IssgOrRcvgMaterial:             data.IssgOrRcvgMaterial,
			IssgOrRcvgBatch:                data.IssgOrRcvgBatch,
			IssuingOrReceivingPlant:        data.IssuingOrReceivingPlant,
			IssuingOrReceivingStorageLoc:   data.IssuingOrReceivingStorageLoc,
			IssuingOrReceivingStockType:    data.IssuingOrReceivingStockType,
			IssgOrRcvgSpclStockInd:         data.IssgOrRcvgSpclStockInd,
			IssuingOrReceivingValType:      data.IssuingOrReceivingValType,
			IsCompletelyDelivered:          data.IsCompletelyDelivered,
			MaterialDocumentItemText:       data.MaterialDocumentItemText,
			UnloadingPointName:             data.UnloadingPointName,
			ShelfLifeExpirationDate:        data.ShelfLifeExpirationDate,
			ManufactureDate:                data.ManufactureDate,
			SerialNumbersAreCreatedAutomly: data.SerialNumbersAreCreatedAutomly,
			Reservation:                    data.Reservation,
			ReservationItem:                data.ReservationItem,
			ReservationIsFinallyIssued:     data.ReservationIsFinallyIssued,
			IsAutomaticallyCreated:         data.IsAutomaticallyCreated,
			GoodsMovementIsCancelled:       data.GoodsMovementIsCancelled,
			ReversedMaterialDocumentYear:   data.ReversedMaterialDocumentYear,
			ReversedMaterialDocument:       data.ReversedMaterialDocument,
			ReversedMaterialDocumentItem:   data.ReversedMaterialDocumentItem,
			ReferenceDocumentFiscalYear:    data.ReferenceDocumentFiscalYear,
			InvtryMgmtRefDocumentItem:      data.InvtryMgmtRefDocumentItem,
			InvtryMgmtReferenceDocument:    data.InvtryMgmtReferenceDocument,
			MaterialDocumentPostingType:    data.MaterialDocumentPostingType,
		})
	}

	return toItem, nil
}
