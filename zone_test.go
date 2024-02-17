package electromap

import (
	"testing"

	tr "gopkg.in/gerald-scharitzer/tecnic.v0/reflect"
	"gotest.tools/v3/assert"
)

func TestGetZones(t *testing.T) {
	apiRoots := []*string{nil, tr.Addr(ApiRootDefault)}
	for _, apiRoot := range apiRoots {
		zones, err := GetZones(apiRoot)
		assert.NilError(t, err)
		assert.Assert(t, len(*zones) > 0)
		if apiRoot == nil {
			t.Log(apiRoot)
		} else {
			t.Log(*apiRoot)
		}
		for id, zone := range *zones {
			assert.Assert(t, len(id) > 0)
			assert.Assert(t, len(zone.Name) > 0)
			t.Log(id, zone.Country, zone.Name)
		}
	}
}
