package sap_api_input_reader

type EC_MC struct {
	ConnectionKey    string `json:"connection_key"`
	Result           bool   `json:"result"`
	RedisKey         string `json:"redis_key"`
	Filepath         string `json:"filepath"`
	MaterialDocument struct {
		MaterialDocument string `json:"document_no"`
		Plant            string `json:"deliver_to"`
		OrderQuantity    string `json:"quantity"`
		PickedQuantity   string `json:"picked_quantity"`
		NetPriceAmount   string `json:"price"`
		Batch            string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema                string   `json:"api_schema"`
	Accepter                 []string `json:"accepter"`
	MaterialCode             string   `json:"material_code"`
	Supplier                 string   `json:"plant/supplier"`
	Stock                    string   `json:"stock"`
	SalesOrderType           string   `json:"document_type"`
	MaterialDocumentNo       string   `json:"document_no"`
	ScheduleLineDeliveryDate string   `json:"planned_date"`
	ValidatedDate            string   `json:"validated_date"`
	Deleted                  bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey    string `json:"connection_key"`
	Result           bool   `json:"result"`
	RedisKey         string `json:"redis_key"`
	Filepath         string `json:"filepath"`
	MaterialDocument struct {
		MaterialDocumentYear       string `json:"MaterialDocumentYear"`
		MaterialDocument           string `json:"MaterialDocument"`
		InventoryTransactionType   string `json:"InventoryTransactionType"`
		DocumentDate               string `json:"DocumentDate"`
		PostingDate                string `json:"PostingDate"`
		CreationDate               string `json:"CreationDate"`
		CreationTime               string `json:"CreationTime"`
		MaterialDocumentHeaderText string `json:"MaterialDocumentHeaderText"`
		ReferenceDocument          string `json:"ReferenceDocument"`
		GoodsMovementCode          string `json:"GoodsMovementCode"`
		MaterialDocumentItem       struct {
			MaterialDocumentItem           string `json:"MaterialDocumentItem"`
			Material                       string `json:"Material"`
			Plant                          string `json:"Plant"`
			StorageLocation                string `json:"StorageLocation"`
			Batch                          string `json:"Batch"`
			GoodsMovementType              string `json:"GoodsMovementType"`
			InventoryStockType             string `json:"InventoryStockType"`
			InventoryValuationType         string `json:"InventoryValuationType"`
			InventorySpecialStockType      string `json:"InventorySpecialStockType"`
			Supplier                       string `json:"Supplier"`
			Customer                       string `json:"Customer"`
			SalesOrder                     string `json:"SalesOrder"`
			SalesOrderItem                 string `json:"SalesOrderItem"`
			SalesOrderScheduleLine         string `json:"SalesOrderScheduleLine"`
			PurchaseOrder                  string `json:"PurchaseOrder"`
			PurchaseOrderItem              string `json:"PurchaseOrderItem"`
			WBSElement                     string `json:"WBSElement"`
			ManufacturingOrder             string `json:"ManufacturingOrder"`
			ManufacturingOrderItem         string `json:"ManufacturingOrderItem"`
			GoodsMovementRefDocType        string `json:"GoodsMovementRefDocType"`
			GoodsMovementReasonCode        string `json:"GoodsMovementReasonCode"`
			Delivery                       string `json:"Delivery"`
			DeliveryItem                   string `json:"DeliveryItem"`
			AccountAssignmentCategory      string `json:"AccountAssignmentCategory"`
			CostCenter                     string `json:"CostCenter"`
			ControllingArea                string `json:"ControllingArea"`
			CostObject                     string `json:"CostObject"`
			ProfitabilitySegment           string `json:"ProfitabilitySegment"`
			ProfitCenter                   string `json:"ProfitCenter"`
			GLAccount                      string `json:"GLAccount"`
			FunctionalArea                 string `json:"FunctionalArea"`
			MaterialBaseUnit               string `json:"MaterialBaseUnit"`
			QuantityInBaseUnit             string `json:"QuantityInBaseUnit"`
			EntryUnit                      string `json:"EntryUnit"`
			QuantityInEntryUnit            string `json:"QuantityInEntryUnit"`
			CompanyCodeCurrency            string `json:"CompanyCodeCurrency"`
			FiscalYear                     string `json:"FiscalYear"`
			FiscalYearPeriod               string `json:"FiscalYearPeriod"`
			IssgOrRcvgMaterial             string `json:"IssgOrRcvgMaterial"`
			IssgOrRcvgBatch                string `json:"IssgOrRcvgBatch"`
			IssuingOrReceivingPlant        string `json:"IssuingOrReceivingPlant"`
			IssuingOrReceivingStorageLoc   string `json:"IssuingOrReceivingStorageLoc"`
			IssuingOrReceivingStockType    string `json:"IssuingOrReceivingStockType"`
			IssgOrRcvgSpclStockInd         string `json:"IssgOrRcvgSpclStockInd"`
			IssuingOrReceivingValType      string `json:"IssuingOrReceivingValType"`
			IsCompletelyDelivered          bool   `json:"IsCompletelyDelivered"`
			MaterialDocumentItemText       string `json:"MaterialDocumentItemText"`
			UnloadingPointName             string `json:"UnloadingPointName"`
			ShelfLifeExpirationDate        string `json:"ShelfLifeExpirationDate"`
			ManufactureDate                string `json:"ManufactureDate"`
			SerialNumbersAreCreatedAutomly bool   `json:"SerialNumbersAreCreatedAutomly"`
			Reservation                    string `json:"Reservation"`
			ReservationItem                string `json:"ReservationItem"`
			ReservationIsFinallyIssued     bool   `json:"ReservationIsFinallyIssued"`
			IsAutomaticallyCreated         string `json:"IsAutomaticallyCreated"`
			GoodsMovementIsCancelled       bool   `json:"GoodsMovementIsCancelled"`
			ReversedMaterialDocumentYear   string `json:"ReversedMaterialDocumentYear"`
			ReversedMaterialDocument       string `json:"ReversedMaterialDocument"`
			ReversedMaterialDocumentItem   string `json:"ReversedMaterialDocumentItem"`
			ReferenceDocumentFiscalYear    string `json:"ReferenceDocumentFiscalYear"`
			InvtryMgmtRefDocumentItem      string `json:"InvtryMgmtRefDocumentItem"`
			InvtryMgmtReferenceDocument    string `json:"InvtryMgmtReferenceDocument"`
			MaterialDocumentPostingType    string `json:"MaterialDocumentPostingType"`
		} `json:"MaterialDocumentItem"`
	} `json:"MaterialDocument"`
	APISchema          string   `json:"api_schema"`
	Accepter           []string `json:"accepter"`
	MaterialDocumentNo string   `json:"material_document"`
	Deleted            bool     `json:"deleted"`
}
