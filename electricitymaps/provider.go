package electricitymaps

// implements electromap.Provider
type ElectricityMaps struct {
	name string
}

func New() *ElectricityMaps {
	return &ElectricityMaps{name: "Electricity Maps"}
}

func (o *ElectricityMaps) Name() string {
	return o.name
}
