// Get carbon efficiency data from electricity maps
// https://api-portal.electricitymaps.com/
// https://static.electricitymaps.com/api/docs/index.html
// https://www.electricitymaps.com/data-portal
package electricitymaps

import (
	"bufio"
	"fmt"
	"net/http"
)

type Monitor struct {
	State string `json:"state"`
}

// response to "health"
// https://static.electricitymaps.com/api/docs/index.html#health
type Health struct {
	Monitors []Monitor
	Status   string `json:"status"`
}

func GetHealth() (int, error) {

	const HOST = "https://api.electricitymap.org/"
	const HEALTH = "health"

	resp, err := http.Get(HOST + HEALTH)
	if err != nil {
		return 0, err
	}

	content_type := resp.Header.Get("content-type")
	if content_type != "application/json; charset=utf-8" {
		return 0, fmt.Errorf("Got content-type %q instead of \"application/json; charset=utf-8\"", content_type)
	}

	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return resp.StatusCode, nil

}
