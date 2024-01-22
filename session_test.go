package electricitymaps

import (
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

func TestGetSession(t *testing.T) {
	want := Session{}
	got := GetSession()
	assert.Equal(t, got, want)
}

const YAML = `
authtoken: test
`

func TestGetSessionFromYaml(t *testing.T) {
	want := Session{AuthToken: "test"}
	got, err := GetSessionFromYaml([]byte(YAML))
	assert.NilError(t, err)
	assert.Equal(t, *got, want)
}

func TestGetSessionFromFile(t *testing.T) {
	want := Session{AuthToken: "test"}
	file, err := os.Open("session.yaml")
	assert.NilError(t, err)
	got, err := GetSessionFromFile(file)
	assert.NilError(t, err)
	assert.Equal(t, *got, want)
}
