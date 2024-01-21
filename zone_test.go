package electricitymaps

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestGetZones(t *testing.T) {
	zones, err := GetZones()
	assert.NilError(t, err)
	assert.Assert(t, len(*zones) > 0)
	for id, zone := range *zones {
		assert.Assert(t, len(id) > 0)
		assert.Assert(t, len(zone.Name) > 0)
		println(id, zone.Country, zone.Name)
	}
}
