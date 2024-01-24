// Get carbon efficiency data from electricity maps
// https://api-portal.electricitymaps.com/
// https://static.electricitymaps.com/api/docs/index.html
// https://www.electricitymaps.com/data-portal
package electricitymaps

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Zone struct {
	Country string `json:"countryName"`
	Name    string `json:"zoneName"`
}

// response
type Zones map[string]Zone

// https://static.electricitymaps.com/api/docs/index.html#zones
func GetZones(apiRoot *string) (*Zones, error) {

	resp, err := http.Get(*apiRoot + "v3/zones")
	if err != nil {
		return nil, err
	}

	content_type := resp.Header.Get("content-type")
	if content_type != "application/json; charset=utf-8" {
		return nil, fmt.Errorf("Got content-type %q instead of \"application/json; charset=utf-8\"", content_type)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	zones := Zones{}
	err = json.Unmarshal(body, &zones)
	return &zones, nil

}
