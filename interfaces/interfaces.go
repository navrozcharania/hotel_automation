package interfaces

type Appliance interface {
	SwitchOn()
	SwitchOff()
	CurrentStatus() bool
	Cost() int64
	Name() string
}

type Corridor interface {
	GetAppliances() []Appliance
	AddAppliance(Appliance)
	Name() string
	CorridorNumber() int64
}

type Floor interface {
	GetCorridors() []Corridor
	AddCorridor(Corridor)
	StoreyNumber() int64
}

type Layout interface {
	GetFloors() []Floor
	AddFloor(Floor)
}

type Controller interface {
	PrintHotelUsage(Layout)
	InitializeAppliances(Layout)
	PrintCostPerFloor(Layout) int64
	RaiseSensorEvent(Layout, int64, string, int64, string) error
}

type Algorithm interface {
	Process(Layout, int64, string, int64, string) error
}