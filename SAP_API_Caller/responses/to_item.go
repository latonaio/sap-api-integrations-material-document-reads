package responses

type ToItem struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			MaterialDocumentYear           string      `json:"MaterialDocumentYear"`
			MaterialDocument               string      `json:"MaterialDocument"`
			MaterialDocumentItem           string      `json:"MaterialDocumentItem"`
			Material                       string      `json:"Material"`
			Plant                          string      `json:"Plant"`
			StorageLocation                string      `json:"StorageLocation"`
			Batch                          string      `json:"Batch"`
			GoodsMovementType              string      `json:"GoodsMovementType"`
			InventoryStockType             string      `json:"InventoryStockType"`
			InventoryValuationType         string      `json:"InventoryValuationType"`
			InventorySpecialStockType      string      `json:"InventorySpecialStockType"`
			Supplier                       string      `json:"Supplier"`
			Customer                       string      `json:"Customer"`
			SalesOrder                     string      `json:"SalesOrder"`
			SalesOrderItem                 string      `json:"SalesOrderItem"`
			SalesOrderScheduleLine         string      `json:"SalesOrderScheduleLine"`
			PurchaseOrder                  string      `json:"PurchaseOrder"`
			PurchaseOrderItem              string      `json:"PurchaseOrderItem"`
			WBSElement                     string      `json:"WBSElement"`
			ManufacturingOrder             string      `json:"ManufacturingOrder"`
			ManufacturingOrderItem         string      `json:"ManufacturingOrderItem"`
			GoodsMovementRefDocType        string      `json:"GoodsMovementRefDocType"`
			GoodsMovementReasonCode        string      `json:"GoodsMovementReasonCode"`
			Delivery                       string      `json:"Delivery"`
			DeliveryItem                   string      `json:"DeliveryItem"`
			AccountAssignmentCategory      string      `json:"AccountAssignmentCategory"`
			CostCenter                     string      `json:"CostCenter"`
			ControllingArea                string      `json:"ControllingArea"`
			CostObject                     string      `json:"CostObject"`
			ProfitabilitySegment           string      `json:"ProfitabilitySegment"`
			ProfitCenter                   string      `json:"ProfitCenter"`
			GLAccount                      string      `json:"GLAccount"`
			FunctionalArea                 string      `json:"FunctionalArea"`
			MaterialBaseUnit               string      `json:"MaterialBaseUnit"`
			QuantityInBaseUnit             string      `json:"QuantityInBaseUnit"`
			EntryUnit                      string      `json:"EntryUnit"`
			QuantityInEntryUnit            string      `json:"QuantityInEntryUnit"`
			CompanyCodeCurrency            string      `json:"CompanyCodeCurrency"`
			FiscalYear                     string      `json:"FiscalYear"`
			FiscalYearPeriod               string      `json:"FiscalYearPeriod"`
			IssgOrRcvgMaterial             string      `json:"IssgOrRcvgMaterial"`
			IssgOrRcvgBatch                string      `json:"IssgOrRcvgBatch"`
			IssuingOrReceivingPlant        string      `json:"IssuingOrReceivingPlant"`
			IssuingOrReceivingStorageLoc   string      `json:"IssuingOrReceivingStorageLoc"`
			IssuingOrReceivingStockType    string      `json:"IssuingOrReceivingStockType"`
			IssgOrRcvgSpclStockInd         string      `json:"IssgOrRcvgSpclStockInd"`
			IssuingOrReceivingValType      string      `json:"IssuingOrReceivingValType"`
			IsCompletelyDelivered          bool        `json:"IsCompletelyDelivered"`
			MaterialDocumentItemText       string      `json:"MaterialDocumentItemText"`
			UnloadingPointName             string      `json:"UnloadingPointName"`
			ShelfLifeExpirationDate        string      `json:"ShelfLifeExpirationDate"`
			ManufactureDate                string      `json:"ManufactureDate"`
			SerialNumbersAreCreatedAutomly bool        `json:"SerialNumbersAreCreatedAutomly"`
			Reservation                    string      `json:"Reservation"`
			ReservationItem                string      `json:"ReservationItem"`
			ReservationIsFinallyIssued     bool        `json:"ReservationIsFinallyIssued"`
			IsAutomaticallyCreated         string      `json:"IsAutomaticallyCreated"`
			GoodsMovementIsCancelled       bool        `json:"GoodsMovementIsCancelled"`
			ReversedMaterialDocumentYear   string      `json:"ReversedMaterialDocumentYear"`
			ReversedMaterialDocument       string      `json:"ReversedMaterialDocument"`
			ReversedMaterialDocumentItem   string      `json:"ReversedMaterialDocumentItem"`
			ReferenceDocumentFiscalYear    string      `json:"ReferenceDocumentFiscalYear"`
			InvtryMgmtRefDocumentItem      string      `json:"InvtryMgmtRefDocumentItem"`
			InvtryMgmtReferenceDocument    string      `json:"InvtryMgmtReferenceDocument"`
			MaterialDocumentPostingType    string      `json:"MaterialDocumentPostingType"`
		} `json:"results"`
	} `json:"d"`
}
