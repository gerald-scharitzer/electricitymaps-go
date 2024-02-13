package electromap

import (
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

func TestGetHealth(t *testing.T) {
	want := Health{Monitors: Monitors{State: "ok"}, Status: "ok"}
	apiRoots := []*string{nil, Addr(ApiRootDefault)}
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

func TestCsv(t *testing.T) {
	health := Health{Monitors: Monitors{State: "ok"}, Status: "ok"}
	err := health.Csv(os.Stdout) // TODO write to log
	assert.NilError(t, err)      // TODO test values
}
