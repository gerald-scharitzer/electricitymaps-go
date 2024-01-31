package electromap

// serves data on the carbon efficiency of electricity in time and space
type Provider interface {
	Name() string
}
