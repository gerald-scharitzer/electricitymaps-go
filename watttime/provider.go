package watttime

// implements electromap.Provider
type WattTime struct {
	name string
}

const defaultName = "WattTime"

func New() *WattTime {
	return &WattTime{name: defaultName}
}

func (o *WattTime) Name() string {
	return o.name
}
