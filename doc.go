// Get carbon efficiency data from Electricity Maps.
//
// https://api-portal.electricitymaps.com/
//
// https://static.electricitymaps.com/api/docs/index.html
//
// https://www.electricitymaps.com/data-portal
package electromap

import (
	"strconv"
	"strings"
)

// Semantic version https://semver.org/
type SemVer struct {
	Major      int
	Minor      int
	Patch      int
	PreRelease string
}

func Version() SemVer {
	return SemVer{Major: 0, Minor: 0, Patch: 2, PreRelease: "alpha.0"}
}

func VersionString() string {
	version := Version()
	var builder strings.Builder
	builder.WriteString(strconv.Itoa(version.Major))
	builder.WriteRune('.')
	builder.WriteString(strconv.Itoa(version.Minor))
	builder.WriteRune('.')
	builder.WriteString(strconv.Itoa(version.Patch))
	if len(version.PreRelease) > 0 {
		builder.WriteRune('-')
		builder.WriteString(version.PreRelease)
	}
	return builder.String()
}
