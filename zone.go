package electromap

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	tr "gopkg.in/gerald-scharitzer/tecnic.v0/reflect"
)

// Zones that are countries have a zone name, but not country name.
// There are zones that belong to countries, but have no country name.
type Zone struct {
	// Not every zone is mapped to a country.
	Country string `json:"countryName"`
	Name    string `json:"zoneName"`
}

// The API response is a map of zone keys to zone values
type Zones map[string]Zone

// Get the map of zones
//
// `apiRoot` points to the target of the API call.
// `nil` calls the API pointed to by `ApiRootDefault`.
//
// # TODO add parameter authToken
//
// # Returns the map of zones or an error
//
// https://static.electricitymaps.com/api/docs/index.html#zones
func GetZones(apiRoot *string) (*Zones, error) {

	if apiRoot == nil { // get default
		apiRoot = tr.Addr(ApiRootDefault)
	}

	resp, err := http.Get(*apiRoot + "v3/zones")
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

	zones := Zones{}
	err = json.Unmarshal(body, &zones)
	if err != nil {
		return nil, err
	}

	return &zones, nil

}
