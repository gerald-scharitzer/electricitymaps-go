# Electromap Go ⚡

is a Go client for the [Electricity Maps API](https://static.electricitymaps.com/api/docs/index.html)
as module [`gopkg.in/gerald-scharitzer/electromap.v0`](https://pkg.go.dev/gopkg.in/gerald-scharitzer/electromap.v0).

[![Go Reference](https://pkg.go.dev/badge/gopkg.in/gerald-scharitzer/electromap.v0.svg)](https://pkg.go.dev/gopkg.in/gerald-scharitzer/electromap.v0)

# Features ✨

Provide a command line interface (CLI) and a Go application programming interface (API) for the web API of Electricity Maps.

Format JSON responses as CSV tables.

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
 6. Check with `go vet ./...`
 7. Build with `go build ./...`
 8. Install with `go install ./...`
 9. Run with `GOBIN/electromap -v` where `GOBIN` is the path to your installed Go binaries to get the version string
10. Tag with `git tag semver`
11. Push with `git push origin semver`
12. Publish with `GOPROXY=proxy.golang.org go list -m gopkg.in/gerald-scharitzer/electromap.vn@semver`
13. Uninstall with `go clean -i ./...`

where `semver` is the [semantic version](https://semver.org/spec/v2.0.0.html) (e.g. v0.0.0)
and `vn` is the major version number (e.g. v0)

## Rules

- Use `fmt.Println` instead of `println`, because [`println`](https://pkg.go.dev/builtin@go1.21.6#println)
is implementation-specific and not guaranteed to stay in the language.
- Use `fmt.Print` instead of `print`, because [`print`](https://pkg.go.dev/builtin@go1.21.6#print)
is implementation-specific and not guaranteed to stay in the language.
