package electromap

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// API health monitors
type Monitors struct {
	// "ok" if the API is fine
	State string
}

// The API response is a structure of states
type Health struct {
	Monitors Monitors
	// "ok" if the API is fine
	Status string
}

// Get the health status of the API
//
// `apiRoot` points to the target of the API call.
// `nil` calls the API pointed to by `ApiRootDefault`.
//
// # TODO add parameter authToken
//
// # Returns the API health or an error
//
// https://static.electricitymaps.com/api/docs/index.html#health
func GetHealth(apiRoot *string) (*Health, error) {

	if apiRoot == nil { // get default
		apiRootDefault := ApiRootDefault // get addressable variable
		apiRoot = &apiRootDefault
	}

	resp, err := http.Get(*apiRoot + "health")
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