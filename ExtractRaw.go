package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

/*
type Service struct {
	ServiceNo string
	Status    string
	Operator  string
}

type DataSet struct {
	Metadata  string `json:"odata.metadata"`
	BusStopID string
	Services  []*Service
}*/
type TickHistoryMarketDepthViewOptions int
type TickHistorySort int
type TickHistoryTimeOptions int
type ReportDateRangeType int
type PreviewMode int
type TickHistoryExtractByMode int

const (
	Ric TickHistoryExtractByMode = iota
	Entity
)
const (
	None PreviewMode = iota
	Content
	Instrument
)
const (
	RawMarketByPrice TickHistoryMarketDepthViewOptions = iota
	RawMarketByOrder
	RawMarketMaker
	LegacyLevel2
	NormalizedLL2
)

const (
	SingleByRic TickHistorySort = iota
	SingleByTimestamp
)

const (
	LocalExchangeTime TickHistoryTimeOptions = iota
	GmtUtc
)

const (
	NoRange ReportDateRangeType = iota
	Init
	Range
	Delta
	Last
)

var tickHistoryExtractByMode = [...]string{"Ric", "Entity"}
var previewMode = [...]string{"None", "Content", "Instrument"}
var reportDateRangeType = [...]string{
	"NoRange",
	"Init",
	"Range",
	"Delta",
	"Last",
}
var tickHistoryMarketDepthViewOptions = [...]string{
	"RawMarketByPrice",
	"RawMarketByOrder",
	"RawMarketMaker",
	"LegacyLevel2",
	"NormalizedLL2",
}

var tickHistorySort = [...]string{
	"SingleByRic",
	"SingleByTimestamp",
}

var tickHistoryTimeOptions = [...]string{
	"LocalExchangeTime",
	"GmtUtc",
}

func (d TickHistoryExtractByMode) MarshalText() ([]byte, error) {
	return []byte(tickHistoryExtractByMode[d]), nil
}

func (d PreviewMode) MarshalText() ([]byte, error) {
	return []byte(previewMode[d]), nil
}

func (d ReportDateRangeType) MarshalText() ([]byte, error) {
	return []byte(reportDateRangeType[d]), nil
}

func (d TickHistoryMarketDepthViewOptions) MarshalText() ([]byte, error) {
	return []byte(tickHistoryMarketDepthViewOptions[d]), nil
}

func (d TickHistorySort) MarshalText() ([]byte, error) {
	return []byte(tickHistorySort[d]), nil
}

func (d TickHistoryTimeOptions) MarshalText() ([]byte, error) {
	return []byte(tickHistoryTimeOptions[d]), nil
}

type InstrumentIdentifier struct {
	Identifier     string
	IdentifierType string
}
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
type TickHistoryMarketDepthCondition struct {
	View                TickHistoryMarketDepthViewOptions
	NumberOfLevels      int32 `json:",omitempty"`
	SortBy              TickHistorySort
	MessageTimeStampIn  TickHistoryTimeOptions
	ReportDateRangeType ReportDateRangeType
	QueryStartDate      *time.Time `json:",omitempty"`
	QueryEndDate        *time.Time `json:",omitempty"`
	DaysAgo             int32      `json:",omitempty"`
	Preview             PreviewMode
	ExtractBy           TickHistoryExtractByMode
	DisplaySourceRIC    bool
}
type InstrumentIdentifierList struct {
	Metadata                               string                       `json:"@odata.type" odata:"#ThomsonReuters.Dss.Api.Extractions.ExtractionRequests.InstrumentIdentifierList"`
	InstrumentIdentifiers                  []InstrumentIdentifier       `json:",omitempty"`
	ValidationOptions                      *InstrumentValidationOptions `json:",omitempty"`
	UseUserPreferencesForValidationOptions bool                         `json:",omitempty"`
}
type TickHistoryMarketDepthExtractionRequest struct {
	Metadata          string                          `json:"@odata.type" odata:"#ThomsonReuters.Dss.Api.Extractions.ExtractionRequests.TickHistoryMarketDepthExtractionRequest"`
	ContentFieldNames []string                        `json:",omitempty"`
	IdentifierList    InstrumentIdentifierList        `json:",omitempty"`
	Condition         TickHistoryMarketDepthCondition `json:",omitempty"`
}
type RequestTokenResponse struct {
	Metadata string `json:"@odata.context,omitempty"`
	Value    string
}
type IdentifierValidationError struct {
	Metadata   string `json:"@odata.context,omitempty"`
	Identifier InstrumentIdentifier
	Message    string
}
type RawExtractionResult struct {
	Metadata                   string `json:"@odata.context,omitempty"`
	JobId                      string
	Notes                      []string
	IdentifierValidationErrors []IdentifierValidationError
}

/*
type RequestTokenMsg struct {
	Credentials Credential
}
*/
type Credential struct {
	Username string
	Password string
}

/*
type ExtractRequest struct {
	ExtractRequest interface{}
}
*/
/*
var sampleData = []byte(`
{
	"@odata.context":"https://hosted.datascopeapi.reuters.com/RestApi/v1/$metadata#Edm.String",
	"value":"_071PUPbXGXUR_Uck5Rc4iPZWdpm6tF923YXbNAvXAoFokXIbxYpd4f7HVRS-EPAlAYUkkVLM9_BGmOr4YwRdi932tgySaTuLyD6HiNofaq72E1MCSUHWiG-xUfgmayhrd7e26amsEqaC8NQqn7sgcsTJcIgpRKuREMZuJrUIM6Q5sBGIW4nlAF5NMTZAH4E3B1kos_2U6rmBAQtRIzRQdWd4OuC-uc0u_n-p0KnG2XU32Svphesk2j0gN-UVCMV9HJHN-6oXgRb__v0m1xxzyG8zakvMyQ-XY4r5Y9r0Wsw"
}`)*/

func (r InstrumentIdentifierList) MarshalJSON() ([]byte, error) {

	type _InstrumentIdentifierList InstrumentIdentifierList
	if r.Metadata == "" {
		st := reflect.TypeOf(r)
		field, _ := st.FieldByName("Metadata")
		r.Metadata = field.Tag.Get("odata")
	}
	return json.Marshal(_InstrumentIdentifierList(r))
	/*
		return json.Marshal(struct {
			Type string `json:"@odata.type"`
			_InstrumentIdentifierList
		}{
			Type: "#ThomsonReuters.Dss.Api.Extractions.ExtractionRequests.InstrumentIdentifierList",
			_InstrumentIdentifierList: _InstrumentIdentifierList(r),
		})
	*/
}

func (r TickHistoryMarketDepthExtractionRequest) MarshalJSON() ([]byte, error) {
	type _TickHistoryMarketDepthExtractionRequest TickHistoryMarketDepthExtractionRequest
	if r.Metadata == "" {
		st := reflect.TypeOf(r)
		field, _ := st.FieldByName("Metadata")
		r.Metadata = field.Tag.Get("odata")
	}

	return json.Marshal(_TickHistoryMarketDepthExtractionRequest(r))
	/*
		type _TickHistoryMarketDepthExtractionRequest TickHistoryMarketDepthExtractionRequest

		return json.Marshal(struct {
			Type string `json:"@odata.type"`
			_TickHistoryMarketDepthExtractionRequest
		}{
			Type: "#ThomsonReuters.Dss.Api.Extractions.ExtractionRequests.TickHistoryMarketDepthExtractionRequest",
			_TickHistoryMarketDepthExtractionRequest: _TickHistoryMarketDepthExtractionRequest(r),
		})*/
}

/*
type _TickHistoryMarketDepthExtractionRequest1 _TickHistoryMarketDepthExtractionRequest

func (r _TickHistoryMarketDepthExtractionRequest) MarshalJson() ([]byte, error) {
	return json.Marshal(struct {
		Type string `json:"@odata.type"`
		_TickHistoryMarketDepthExtractionRequest1
	}{
		Type: "#ThomsonReuters.Dss.Api.Extractions.ExtractionRequests.TickHistoryMarketDepthExtractionRequest",
		_TickHistoryMarketDepthExtractionRequest1: _TickHistoryMarketDepthExtractionRequest1(r),
	})
}
*/
/*
func (r TickHistoryMarketDepthExtractionRequest) MarshalJSON() ([]byte, error) {

			temp := struct {
				ExtractRequest struct {
					Type string `json:"@odata.type"`
					_TickHistoryMarketDepthExtractionRequest
				}{
			 	Type: "#ThomsonReuters.Dss.Api.Extractions.ExtractionRequests.TickHistoryMarketDepthExtractionRequest",
				_TickHistoryMarketDepthExtractionRequest: _TickHistoryMarketDepthExtractionRequest(r),
			}
		}
			}{
				ExtractRequest: _TickHistoryMarketDepthExtractionRequest(r),
			}
			return json.Marshal(temp)


				return json.Marshal(struct {
					ExtractRequest struct {
						Type string `json:"@odata.type"`
						_TickHistoryMarketDepthExtractionRequest
					}{
						Type: "#ThomsonReuters.Dss.Api.Extractions.ExtractionRequests.TickHistoryMarketDepthExtractionRequest",
						_TickHistoryMarketDepthExtractionRequest: _TickHistoryMarketDepthExtractionRequest(r),
					}{}
					}
				)
}
*/

/*
func NewTickHistoryMarketDepthExtractionRequest() *TickHistoryMarketDepthExtractionRequest {
	obj := new(TickHistoryMarketDepthExtractionRequest)
	obj.Type = "#ThomsonReuters.Dss.Api.Extractions.ExtractionRequests.TickHistoryMarketDepthExtractionRequest"
	obj.IdentifierList.Type = "#ThomsonReuters.Dss.Api.Extractions.ExtractionRequests.InstrumentIdentifierList"
	return obj
}
*/
/*
func HttpGetRange(client *http.Client, url string, token *string, start int, end int, trace bool) (*http.Response, error) {
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Prefer", "respond-async")
	if end == -1 {
		req.Header.Add("Range", "bytes="+strconv.Itoa(start)+"-")
	} else {
		req.Header.Add("Range", "bytes="+strconv.Itoa(start)+"-"+strconv.Itoa(end))
	}
	if token != nil {
		req.Header.Add("Authorization", "Token "+*token)
	}
	if trace == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		fmt.Println(string(dump))
	}

	resp, err := client.Do(req)

	if trace == true {
		dumpBody := true
		contentLength, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
		if contentLength > 100 {
			dumpBody = false
		}

		dump, _ := httputil.DumpResponse(resp, dumpBody)
		log.Println(string(dump))
	}

	return resp, err
}*/
func HttpGet(client *http.Client, url string, token *string, trace bool) (*http.Response, error) {
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Prefer", "respond-async")
	if token != nil {
		req.Header.Add("Authorization", "Token "+*token)
	}
	if trace == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	resp, err := client.Do(req)

	if trace == true {
		dumpBody := true
		contentLength, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
		if contentLength > 100 {
			dumpBody = false
		}

		dump, _ := httputil.DumpResponse(resp, dumpBody)
		fmt.Println(string(dump))
	}

	return resp, err

}
func PrintDownloadPercent(done chan int64, path string, total int64) {

	var stop bool = false

	for {
		select {
		case <-done:
			stop = true
		default:

			file, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}

			fi, err := file.Stat()
			if err != nil {
				log.Fatal(err)
			}

			size := fi.Size()

			if size == 0 {
				size = 1
			}

			var percent float64
			percent = float64(size) / float64(total) * 100

			log.Printf("%s, Bytes: %d/Total: %d (%.0f%%)", path, size, total, percent)

		}

		if stop {
			break
		}

		time.Sleep(time.Second * 5)
	}
}
func HttpPost(client *http.Client, url string, token *string, body *bytes.Buffer, trace bool) (*http.Response, error) {
	//req, err := http.NewRequest("POST", trthURL+"Authentication/RequestToken", bytes.NewBuffer(b))
	req, _ := http.NewRequest("POST", url, body)
	//req, err := http.NewRequest("POST", trthURL+"Authentication/RequestToken", bytes.NewBuffer(jsonStr))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Prefer", "respond-async")
	if token != nil {
		req.Header.Add("Authorization", "Token "+*token)
	}
	if trace == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		fmt.Println(string(dump))
	}

	resp, err := client.Do(req)

	if trace == true {

		dumpBody := true
		contentLength, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
		if contentLength > 100 {
			dumpBody = false
		}

		dump, _ := httputil.DumpResponse(resp, dumpBody)
		fmt.Println(string(dump))
	}

	return resp, err
}
func main() {
	trthURL := "https://hosted.datascopeapi.reuters.com/RestApi/v1/"
	//var jsonStr = []byte(`{"Credentials":{"Username":"9008895", "Password":"Reuters123"}}`)
	request := new(TickHistoryMarketDepthExtractionRequest)

	request.Condition.View = NormalizedLL2
	request.Condition.SortBy = SingleByRic
	request.Condition.NumberOfLevels = 10
	request.Condition.MessageTimeStampIn = GmtUtc
	request.Condition.DisplaySourceRIC = true
	request.Condition.ReportDateRangeType = Range
	startdate := time.Date(2016, 8, 29, 9, 0, 0, 0, time.UTC)
	request.Condition.QueryStartDate = &startdate
	enddate := time.Date(2016, 9, 29, 12, 0, 0, 0, time.UTC)
	request.Condition.QueryEndDate = &enddate
	//request.Condition.QueryEndDate = nil
	request.ContentFieldNames = []string{
		"Ask Price",
		"Ask Size",
		"Bid Price",
		"Bid Size",
		"Domain",
		"History End",
		"History Start",
		"Instrument ID",
		"Instrument ID Type",
		"Number of Buyers",
		"Number of Sellers",
		"Sample Data",
	}

	//request.ContentFieldNames = append(request.ContentFieldNames, "Ask Size")

	request.IdentifierList.InstrumentIdentifiers = append(request.IdentifierList.InstrumentIdentifiers, InstrumentIdentifier{Identifier: "IBM.N", IdentifierType: "Ric"})
	request.IdentifierList.ValidationOptions = &InstrumentValidationOptions{AllowHistoricalInstruments: true}
	//request.IdentifierList.ValidationOptions.AllowHistoricalInstruments = true
	/*
		reqxx := struct {
			ExtractRequest *TickHistoryMarketDepthExtractionRequest
		}{
			ExtractRequest: request,
		}*/

	//req1, _ := json.Marshal(reqxx)

	//req1, _ := json.Marshal(ExtractRequest{ExtractRequest: request})

	//fmt.Println(string(req1))
	/*
		map1 := map[string]interface{}{
			"ExtractRequest": map[string]interface{}{
				"@odata.type":       "#ThomsonReuters.Dss.Api.Extractions.ExtractionRequests.TickHistoryMarketDepthExtractionRequest",
				"ContentFieldNames": [2]string{"BID", "ASK"},
				"IdentifierList": map[string]interface{}{
					"@odata.type": "#ThomsonReuters.Dss.Api.Extractions.ExtractionRequests.InstrumentIdentifierList",
					"InstrumentIdentifiers": [2]map[string]string{
						{
							"Identifier":     "IBM.N",
							"IdentifierType": "Ric",
						}, {
							"Identifier":     "PTT.BK",
							"IdentifierType": "Ric",
						},
					},
					"ValidationOptions": map[string]interface{}{
						"AllowHistoricalInstruments": true,
					},
				},
				"Condition": map[string]interface{}{
					"View":             "RawMarketPricePrice",
					"SortBy":           "SingleByRic",
					"QueryStartDate":   time.Date(2016, 5, 11, 19, 0, 0, 0, time.UTC),
					"DisplaySourceRic": false,
				},
			},
		}
	*/
	//map1 = make(map[string]interface{})
	//req2, _ := json.Marshal(map1)
	//fmt.Println(string(req2))
	tr := &http.Transport{
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	/*
		message := &RequestTokenMsg{
			Credentials: Credential{
				"9008895",
				"Reuters123",
			},
		}
	*/
	//b, err := json.Marshal(message)
	b, err := json.Marshal(struct {
		Credentials Credential
	}{
		Credentials: Credential{
			"9008895",
			"Reuters123",
		},
	})
	resp, err := HttpPost(client, trthURL+"Authentication/RequestToken", nil, bytes.NewBuffer(b), false)

	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		panic(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		//var objmap map[string]interface{}
		//err = json.Unmarshal(body, &objmap)
		//errorMessage := objmap["error"].(map[string]interface{})
		log.Fatalf("Status Code: %s\n%s ", resp.Status, string(body))
		//panic(errorMessage["message"])
	}

	var respMsg = &RequestTokenResponse{}
	//fmt.Println("response body:", string(sampleData))
	err = json.Unmarshal(body, respMsg)
	resp.Body.Close()
	//err = json.Unmarshal(sampleData, respMsg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Context: %s\n", respMsg.Metadata)
	fmt.Printf("Token: %s\n", respMsg.Value)
	token := respMsg.Value

	req1, _ := json.Marshal(struct {
		ExtractionRequest *TickHistoryMarketDepthExtractionRequest
	}{
		ExtractionRequest: request,
	})

	resp, err = HttpPost(client, trthURL+"Extractions/ExtractRaw", &token, bytes.NewBuffer(req1), true)

	if err != nil {
		log.Fatal(err)
	}

	for resp.StatusCode == 202 {
		time.Sleep(3000 * time.Millisecond)
		location := resp.Header.Get("Location")
		location = strings.Replace(location, "http:", "https:", 1)
		resp, err = HttpGet(client, location, &token, true)
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("Status Code: %s\n%s ", resp.Status, string(body))
	}

	extractRawResult := &RawExtractionResult{}
	err = json.Unmarshal(body, extractRawResult)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(extractRawResult.Metadata)
	fmt.Println(extractRawResult.JobId)
	fmt.Println(extractRawResult.Notes)
	resp.Body.Close()

	jobIDURL := trthURL + "Extractions/RawExtractionResults('" + extractRawResult.JobId + "')" + "/$value"
	//jobIDURL := trthURL + "StandardExtractions/UserPackageDeliveries('0x05d4d06c151b2f86')/$value"
	resp, err = HttpGet(client, jobIDURL, &token, true)

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(body))
	if resp.StatusCode != 200 {
		body, _ = ioutil.ReadAll(resp.Body)
		log.Fatalf("Status Code: %s\n%s ", resp.Status, string(body))
	}

	size, err := strconv.Atoi(resp.Header.Get("Content-Length"))

	if err != nil {
		log.Fatal(err)
	}

	done := make(chan int64)
	outputFileName := "output_" + strconv.Itoa(os.Getpid()) + ".csv.gz"

	out, err := os.Create(outputFileName)
	if err != nil {
		log.Fatal(err)
	}

	go PrintDownloadPercent(done, outputFileName, int64(size))

	n, err := io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	done <- n
	resp.Body.Close()
	log.Println("Download Completed!")

}
