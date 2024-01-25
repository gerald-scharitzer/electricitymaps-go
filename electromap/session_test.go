package electromap

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

const sessionYaml = `
apiRoot: https://api.electricitymap.org/
authToken: test
`

func TestGetSessionFromYaml(t *testing.T) {
	want := Session{ApiRoot: "https://api.electricitymap.org/", AuthToken: "test"}
	got, err := GetSessionFromYaml([]byte(sessionYaml))
	assert.NilError(t, err)
	assert.Equal(t, *got, want)
}

func TestGetSessionFromFile(t *testing.T) {
	want := Session{ApiRoot: "https://api.electricitymap.org/", AuthToken: "test"}
	file, err := os.Open("../session.yaml")
	assert.NilError(t, err)
	got, err := GetSessionFromFile(file)
	assert.NilError(t, err)
	assert.Equal(t, *got, want)
}
