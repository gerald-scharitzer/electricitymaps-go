# Electromap Go ⚡

is a Go client for the [Electricity Maps API](https://static.electricitymaps.com/api/docs/index.html)
as module [`gopkg.in/gerald-scharitzer/electromap.v0`](https://pkg.go.dev/gopkg.in/gerald-scharitzer/electromap.v0).

[![Go Reference](https://pkg.go.dev/badge/gopkg.in/gerald-scharitzer/electromap.v0.svg)](https://pkg.go.dev/gopkg.in/gerald-scharitzer/electromap.v0)

# Use 🔌

the binary as command line interface with `electromap`. Get help for that with `electromap -h`.

Use the Go API by importing the module.

```
import (
    em "gopkg.in/gerald-scharitzer/electromap.v0"
)
```

# Develop 🚀

 1. Get with `git clone https://github.com/gerald-scharitzer/electromap.git`
 2. Enter with `cd electromap`
 3. Increase version in [doc.go#Version](doc.go#Version)
 4. Test with `go test ./...`
 5. Run with `go run ./electromap`
 6. Build with `go build ./...`
 7. Install with `go install ./...`
 8. Run with `GOBIN/electromap` where `GOBIN` is the path to your installed Go binaries
 9. Tag with `git tag semver`
10. Push with `git push origin semver`
11. Publish with `GOPROXY=proxy.golang.org go list -m gopkg.in/gerald-scharitzer/electromap.vn@semver`
12. Uninstall with `go clean -i ./...`

where `semver` is the [semantic version](https://semver.org/spec/v2.0.0.html) (e.g. v0.0.0)
and `vn` is the major version number (e.g. v0)

# Todo 🚨

- Replace `print` and `println`
