package electromap

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestAddr(t *testing.T) {
	// test with int and string TODO: add more types
	tests := []any{42, "42"}
	for _, test := range tests {
		want := &test
		got := addr(test)
		assert.Equal(t, *got, *want)
	}
}
