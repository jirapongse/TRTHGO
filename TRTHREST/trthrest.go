package TRTHREST

import (
	"encoding/json"
	"reflect"
	"time"
)

//TickHistoryMarketDepthViewOptions : This is an enumeration for TickHistoryMarketDepthViewOptions
type TickHistoryMarketDepthViewOptions int

//TickHistorySort : This is an enumeration for TickHistorySort
type TickHistorySort int

//TickHistoryTimeOptions :  This is an enumeration for TickHistoryTimeOptions
type TickHistoryTimeOptions int

//ReportDateRangeType : This is an enumeration for ReportDateRangeType
type ReportDateRangeType int

//PreviewMode : This is an enumeration for PreviewMode
type PreviewMode int

//TickHistoryExtractByMode : This is an enumeration for TickHistoryExtractByMode
type TickHistoryExtractByMode int

//Available Enumerations for TickHistoryMarketDepthViewOptions
const (
	RawMarketByPrice TickHistoryMarketDepthViewOptions = iota
	RawMarketByOrder
	RawMarketMaker
	LegacyLevel2
	NormalizedLL2
)

//Available Enumerations for TickHistoryExtractByMode
const (
	Ric TickHistoryExtractByMode = iota
	Entity
)

//Available Enumerations for PreviewMode
const (
	None PreviewMode = iota
	Content
	Instrument
)

//Available Enumerations for TickHistorySort
const (
	SingleByRic TickHistorySort = iota
	SingleByTimestamp
)

//Available Enumerations for TickHistoryTimeOptions
const (
	LocalExchangeTime TickHistoryTimeOptions = iota
	GmtUtc
)

//Available Enumerations for ReportDateRangeType
const (
	NoRange ReportDateRangeType = iota
	Init
	Range
	Delta
	Last
)

//Enumeration String of tickHistoryExtractByMode enumeration used by Marshaller while encoding to JSON
var tickHistoryExtractByMode = [...]string{"Ric", "Entity"}

//Enumeration String of previewMode enumeration used by Marshaller while encoding to JSON
var previewMode = [...]string{"None", "Content", "Instrument"}

//Enumeration String of reportDateRangeType enumeration used by Marshaller while encoding to JSON
var reportDateRangeType = [...]string{
	"NoRange",
	"Init",
	"Range",
	"Delta",
	"Last",
}

//Enumeration String of tickHistoryMarketDepthViewOptions enumeration used by Marshaller while encoding to JSON
var tickHistoryMarketDepthViewOptions = [...]string{
	"RawMarketByPrice",
	"RawMarketByOrder",
	"RawMarketMaker",
	"LegacyLevel2",
	"NormalizedLL2",
}

//Enumeration String of tickHistorySort enumeration used by Marshaller while encoding to JSON
var tickHistorySort = [...]string{
	"SingleByRic",
	"SingleByTimestamp",
}

//Enumeration String of tickHistoryTimeOptions enumeration used by Marshaller while encoding to JSON
var tickHistoryTimeOptions = [...]string{
	"LocalExchangeTime",
	"GmtUtc",
}

//MarshalText : JSON Marshaller for TickHistoryExtractByMode enumeration.
//It uses tickHistoryExtractByMode string array variable to convert int (enum) to string
func (d TickHistoryExtractByMode) MarshalText() ([]byte, error) {
	return []byte(tickHistoryExtractByMode[d]), nil
}

//MarshalText : JSON Marshaller for PreviewMode enumeration.
//It uses previewMode string array variable to convert int (enum) to string
func (d PreviewMode) MarshalText() ([]byte, error) {
	return []byte(previewMode[d]), nil
}

//MarshalText : JSON Marshaller for ReportDateRangeType enumeration.
//It uses reportDateRangeType string array variable to convert int (enum) to string
func (d ReportDateRangeType) MarshalText() ([]byte, error) {
	return []byte(reportDateRangeType[d]), nil
}

//MarshalText : JSON Marshaller for TickHistoryMarketDepthViewOptions enumeration.
//It uses tickHistoryMarketDepthViewOptions string array variable to convert int (enum) to string
func (d TickHistoryMarketDepthViewOptions) MarshalText() ([]byte, error) {
	return []byte(tickHistoryMarketDepthViewOptions[d]), nil
}

//MarshalText : JSON Marshaller for TickHistorySort enumeration.
//It uses tickHistorySort string array variable to convert int (enum) to string
func (d TickHistorySort) MarshalText() ([]byte, error) {
	return []byte(tickHistorySort[d]), nil
}

//MarshalText : JSON Marshaller for TickHistoryTimeOptions enumeration.
//It uses tickHistoryTimeOptions string array variable to convert int (enum) to string
func (d TickHistoryTimeOptions) MarshalText() ([]byte, error) {
	return []byte(tickHistoryTimeOptions[d]), nil
}

//InstrumentIdentifier : defined type for instrument used in InstrumentIdentifierList. This type will be encoded to Json by Marshaller
type InstrumentIdentifier struct {
	Identifier     string
	IdentifierType string
}

//InstrumentValidationOptions : defined type for InstrumentValidationOptions used in InstrumentIdentifierList. This type will be encoded to Json by Marshaller
type InstrumentValidationOptions struct {
	AllowOpenAccessInstruments          bool `json:",omitempty"`
	AllowHistoricalInstruments          bool `json:",omitempty"`
	ExcludeFinrAsPricingSourceForBonds  bool `json:",omitempty"`
	UseExchangeCodeInsteadOfLipper      bool `json:",omitempty"`
	UseUsQuoteInsteadOfCanadian         bool `json:",omitempty"`
	UseConsolidatedQuoteSourceForUsa    bool `json:",omitempty"`
	UseConsolidatedQuoteSourceForCanada bool `json:",omitempty"`
	UseDebtOverEquity                   bool `json:",omitempty"`
}

//TickHistoryMarketDepthCondition : defined type for TickHistoryMarketDepthCondition used in TickHistoryMarketDepthExtractionRequest. This type will be encoded to Json by Marshaller
type TickHistoryMarketDepthCondition struct {
	View                TickHistoryMarketDepthViewOptions
	NumberOfLevels      int32 `json:",omitempty"`
	SortBy              TickHistorySort
	MessageTimeStampIn  TickHistoryTimeOptions
	ReportDateRangeType ReportDateRangeType
	//QueryStartDate is defined as pointer because it is optional
	QueryStartDate *time.Time `json:",omitempty"`
	//QueryEndDate is defined as pointer because it is optional
	QueryEndDate     *time.Time `json:",omitempty"`
	DaysAgo          int32      `json:",omitempty"`
	Preview          PreviewMode
	ExtractBy        TickHistoryExtractByMode
	DisplaySourceRIC bool
}

//InstrumentIdentifierList : defined type for InstrumentIdentifierList used in TickHistoryMarketDepthExtractionRequest. This type will be encoded to Json by Marshaller
type InstrumentIdentifierList struct {
	//It uses 'json' metadata to change the fieldname from Metadata to @data.type
	//It uses user-defined 'odata' metadata to define the default value
	Metadata              string                 `json:"@odata.type" odata:"#ThomsonReuters.Dss.Api.Extractions.ExtractionRequests.InstrumentIdentifierList"`
	InstrumentIdentifiers []InstrumentIdentifier `json:",omitempty"`
	//ValidationOptions is defined as pointer because it is optional
	ValidationOptions                      *InstrumentValidationOptions `json:",omitempty"`
	UseUserPreferencesForValidationOptions bool                         `json:",omitempty"`
}

//TickHistoryMarketDepthExtractionRequest : defined type for TickHistoryMarketDepthExtractionRequest. This type is used to request TRTH Market Depth data
//This type will be encoded to Json by Marshaller
type TickHistoryMarketDepthExtractionRequest struct {
	//It uses 'json' metadata to change the fieldname from Metadata to @data.type
	//It uses user-defined 'odata' metadata to define the default value
	Metadata          string                          `json:"@odata.type" odata:"#ThomsonReuters.Dss.Api.Extractions.ExtractionRequests.TickHistoryMarketDepthExtractionRequest"`
	ContentFieldNames []string                        `json:",omitempty"`
	IdentifierList    InstrumentIdentifierList        `json:",omitempty"`
	Condition         TickHistoryMarketDepthCondition `json:",omitempty"`
}

//RequestTokenResponse : The HTTP response from Authentication/RequestToken request will be decoded to this type by json.Unmarshal
type RequestTokenResponse struct {
	//The value in @odata.content field will be decoded to Metadata field
	Metadata string `json:"@odata.context,omitempty"`
	Value    string
}

//IdentifierValidationError : If the indentifier in the request is invalid, the repsonse from extraction (RawExtractionResult) will contain this information.
type IdentifierValidationError struct {
	//The value in '@odata.content' field will be decoded to 'Metadata' field
	Metadata   string `json:"@odata.context,omitempty"`
	Identifier InstrumentIdentifier
	Message    string
}

//RawExtractionResult : The HTTP response from the completed Extractions/ExtractRaw request will be decoded to this type by json.Unmarshal
type RawExtractionResult struct {
	//The value in '@odata.content' field will be decoded to this 'Metadata' field
	Metadata string `json:"@odata.context,omitempty"`
	//The value in 'JobId' field will be decoded to this 'JobID field
	JobID                      string `json:"JobId"`
	Notes                      []string
	IdentifierValidationErrors []IdentifierValidationError
}

//Credential : The type is used Authentication/RequestToken request. It will be encoded to JSON by Marshaller
type Credential struct {
	Username string
	Password string
}

//MarshalJSON : The custom JSON Marshaller for InstrumentIdentifierList. It uses reflection to set the value for 'Metadata' field.
//The default value is from 'odata" metadata
func (r InstrumentIdentifierList) MarshalJSON() ([]byte, error) {
	//This type is defined to avoid recursive while marshaling modified InstrumentIdentifierList
	//The modified InstrumentIdentifierList will be copied to this type
	//Therefore, json.Marshal can encode it to JSON with the value in 'Metatdata' field
	type _InstrumentIdentifierList InstrumentIdentifierList
	if r.Metadata == "" {
		st := reflect.TypeOf(r)
		field, _ := st.FieldByName("Metadata")
		r.Metadata = field.Tag.Get("odata")
	}
	return json.Marshal(_InstrumentIdentifierList(r))

}

//MarshalJSON : The custom JSON Marshaller for TickHistoryMarketDepthExtractionRequest. It uses reflection to set the value for 'Metadata' field.
//The default value is from 'odata" metadata
func (r TickHistoryMarketDepthExtractionRequest) MarshalJSON() ([]byte, error) {
	//This type is defined to avoid recursive while marshaling modified _TickHistoryMarketDepthExtractionRequest
	//The modified _TickHistoryMarketDepthExtractionRequest will be copied to this type.
	//Therefore, json.Marshal can encode it to JSON with the value in 'Metatdata' field
	type _TickHistoryMarketDepthExtractionRequest TickHistoryMarketDepthExtractionRequest
	if r.Metadata == "" {
		st := reflect.TypeOf(r)
		field, _ := st.FieldByName("Metadata")
		r.Metadata = field.Tag.Get("odata")
	}
	return json.Marshal(_TickHistoryMarketDepthExtractionRequest(r))
}
