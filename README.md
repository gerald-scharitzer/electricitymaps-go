# Electromap Go

is a Go client for the [Electricity Maps API](https://static.electricitymaps.com/api/docs/index.html) as module `gopkg.in/gerald-scharitzer/electromap`.

# Develop

1. Get with `git clone https://github.com/gerald-scharitzer/electromap.git`
2. Enter with `cd electromap-go`
3. Test with `go test ./...`
4. Run with `go run ./main`
5. Tag with `git tag semver`
6. Push with `git push origin semver`
7. Publish with `GOPROXY=proxy.golang.org go list -m gopkg.in/gerald-scharitzer/electromap@semver`

where `semver` is the semantic version (e.g. v0.0.0)
