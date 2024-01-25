// Get carbon efficiency data from electricity maps
// https://api-portal.electricitymaps.com/
// https://static.electricitymaps.com/api/docs/index.html
// https://www.electricitymaps.com/data-portal
package electromap

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Monitors struct {
	State string
}

// response
type Health struct {
	Monitors Monitors
	Status   string
}

// https://static.electricitymaps.com/api/docs/index.html#health
func GetHealth() (*Health, error) {

	const host = "https://api.electricitymap.org/"

	resp, err := http.Get(host + "health")
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

	health := Health{}
	err = json.Unmarshal(body, &health)
	return &health, nil

}
