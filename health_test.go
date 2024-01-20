package electricitymaps

import (
	"net/http"
	"testing"

	"gotest.tools/v3/assert"
)

func TestGetHealth(t *testing.T) {
	want := http.StatusOK
	got, err := GetHealth()
	assert.NilError(t, err)
	assert.Equal(t, got, want)
}
