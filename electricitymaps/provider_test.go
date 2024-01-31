package electricitymaps

import (
	"testing"

	"gotest.tools/assert"
)

func TestNew(t *testing.T) {
	want := ElectricityMaps{name: "Electricity Maps"}
	got := New()
	assert.Equal(t, *got, want)
}

func TestName(t *testing.T) {
	want := "Electricity Maps"
	got := New().Name()
	assert.Equal(t, got, want)
}
