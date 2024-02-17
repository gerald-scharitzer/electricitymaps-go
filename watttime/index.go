package watttime

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	tr "gopkg.in/gerald-scharitzer/tecnic.v0/reflect"
)

const (
	DefaultRegion             = "CAISO_NORTH"
	DefaultSignalType         = "co2_moer"
	DefaultUnits              = "CO2 lbs/MWh"
	ModelTypeAverage          = "average"
	ModelTypeBinnedRegression = "binned_regression"
	ModelTypeHeatrate         = "heatrate"
	ModelTypeProxy            = "proxy"
	ModelTypeSyntheticProxy   = "synthetic_proxy"
)

// API time series
type Data struct {
	// datetime
	PointTime string `json:"point_time"`
	Value     float64
}

// API warning
type Warning struct {
	Type    string
	Message string
}

// API model
type Model struct {
	Date string
	Type string // TODO enum
}

// API metadata
type Meta struct {
	// duration in seconds from the start specified by `Data.PointTime`
	DataPointPeriodSeconds int `json:"data_point_period_seconds"`
	Region                 string
	SignalType             string `json:"signal_type"`
	Units                  string
	Warnings               []Warning
	Model                  Model
}

// The API response is a structure of data and its metadata
type SignalIndex struct {
	Data []Data // array of Data
	Meta Meta
}

// Get the current CO2 MOER index
//
// `apiRoot` points to the target of the API call.
// `nil` calls the API pointed to by `ApiRootDefault`.
//
// `region` is the grid region. `nil` defaults to `DefaultRegion`.
//
// `signalType` is the signal type. `nil` defaults to `DefaultSignalType`.
//
// # TODO Implement parameter types as stringers
//
// # TODO add parameter authToken
//
// # Returns the API signal index or an error
//
// https://docs.watttime.org/#tag/GET-Index/operation/get_signal_index_v3_signal_index_get
func GetSignalIndex(apiRoot *string, region *string, signalType *string) (*SignalIndex, error) {

	if apiRoot == nil { // get default
		apiRoot = tr.Addr(ApiRootDefault)
	}

	if region == nil {
		region = tr.Addr(DefaultRegion)
	}

	if signalType == nil {
		signalType = tr.Addr(DefaultSignalType)
	}

	req, err := http.NewRequest(http.MethodGet, *apiRoot+"v3/signal-index", nil) // TODO JoinPath
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Add("region", *region)
	query.Add("signal_type", *signalType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	content_type := resp.Header.Get("content-type") // TODO HTTP constants
	if content_type != "application/json; charset=utf-8" {
		return nil, fmt.Errorf("got content-type %q instead of \"application/json; charset=utf-8\"", content_type)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	health := SignalIndex{}
	err = json.Unmarshal(body, &health)
	if err != nil {
		return nil, err
	}

	return &health, nil

}

// Get comma-separated values
func (signalIndex *SignalIndex) Csv(writer io.Writer) error {
	csvWriter := csv.NewWriter(writer)                      // TODO reuse csvWriter
	err := csvWriter.Write([]string{"Point Time", "Value"}) // TODO make optional
	if err != nil {
		return err
	}
	for _, data := range signalIndex.Data {
		err = csvWriter.Write([]string{data.PointTime, fmt.Sprint(data.Value)})
		if err != nil {
			return err
		}
	}
	csvWriter.Flush()
	err = csvWriter.Error()
	if err != nil {
		return err
	}
	return nil
}
