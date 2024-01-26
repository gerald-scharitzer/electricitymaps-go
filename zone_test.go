package electromap

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestGetZones(t *testing.T) {
	apiRoot := "https://api.electricitymap.org/"
	zones, err := GetZones(&apiRoot)
	assert.NilError(t, err)
	assert.Assert(t, len(*zones) > 0)
	for id, zone := range *zones {
		assert.Assert(t, len(id) > 0)
		assert.Assert(t, len(zone.Name) > 0)
		// TODO replace with log: println(id, zone.Country, zone.Name)
	}
}
