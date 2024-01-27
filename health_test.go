package electromap

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestGetHealth(t *testing.T) {
	want := Health{Monitors: Monitors{State: "ok"}, Status: "ok"}
	apiRootDefault := ApiRootDefault
	apiRoots := []*string{nil, &apiRootDefault}
	for _, apiRoot := range apiRoots {
		got, err := GetHealth(apiRoot)
		assert.NilError(t, err)
		assert.Equal(t, *got, want)
		var name string
		if apiRoot == nil {
			name = "nil"
		} else {
			name = *apiRoot
		}
		t.Log(name, got.Monitors.State, got.Status)
	}
}
