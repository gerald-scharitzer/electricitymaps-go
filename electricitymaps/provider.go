package electricitymaps

// implements electromap.Provider
type ElectricityMaps struct {
	name string
}

const defaultName = "Electricity Maps"

func New() *ElectricityMaps {
	return &ElectricityMaps{name: defaultName}
}

func (o *ElectricityMaps) Name() string {
	return o.name
}
