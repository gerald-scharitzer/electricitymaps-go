package electromap

// Get the address of a value as expression
//
//	a := addr(value)
//
// instead of multiple statements:
//
//	v := value
//	a := &v
//
// The argument is passed by value, so the address of the copy is returned.
// The memory is allocated implicitly through pass by value.
func addr[T any](x T) *T {
	return &x
}
