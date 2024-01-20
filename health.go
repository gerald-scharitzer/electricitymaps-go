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

type Monitors struct {
	State string
}

// response to "health"
// https://static.electricitymaps.com/api/docs/index.html#health
type Health struct {
	Monitors Monitors
	Status   string
}

func GetHealth() (Health, error) {

	const HOST = "https://api.electricitymap.org/"
	const HEALTH = "health"
	health := Health{}

	resp, err := http.Get(HOST + HEALTH)
	if err != nil {
		return health, err
	}

	content_type := resp.Header.Get("content-type")
	if content_type != "application/json; charset=utf-8" {
		return health, fmt.Errorf("Got content-type %q instead of \"application/json; charset=utf-8\"", content_type)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return health, err
	}

	err = json.Unmarshal(body, &health)
	return health, nil

}
