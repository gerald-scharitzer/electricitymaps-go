package electricitymaps

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestGetHealth(t *testing.T) {
	want := Health{Monitors: Monitors{State: "ok"}, Status: "ok"}
	got, err := GetHealth()
	assert.NilError(t, err)
	assert.Equal(t, got, want)
	println("health", got.Monitors.State, got.Status)
}
