package appliance

import (
	"hotelElectricsSystem/common/constants"
	"hotelElectricsSystem/interfaces"
	"log"
)

// defined ac and lights as separate structs because then they might want to have different implementations at a later stage
// for example for light one might want to add special logic for allowing to switch off only in the day time
type light struct {
	powerOn bool
	cost    int64
	name    string
}

func (a *light) SwitchOn() {
	log.Print("LIGHTS ON.")
	a.powerOn = true
}

func (a *light) SwitchOff() {
	log.Print("LIGHTS OFF.")
	a.powerOn = false
}

func (a *light) CurrentStatus() bool {
	return a.powerOn
}

func (a *light) Cost() int64 {
	return a.cost
}

func (a *light) Name() string {
	return a.name
}

func NewLight() interfaces.Appliance {
	return &light{
		powerOn: true,
		cost:    constants.COST_LIGHT,
		name:    constants.LIGHT,
	}
}
