package appliance

import (
	"hotelElectricsSystem/common/constants"
	"hotelElectricsSystem/interfaces"
	"log"
)

// defined ac and lights as separate structs because then they might want to have different implementations at a later stage
type ac struct {
	powerOn bool
	cost int64
	name string
}

func (a *ac) SwitchOn() {
	log.Print("AC Switched on.")
	a.powerOn = true
}

func (a *ac) SwitchOff() {
	log.Print("AC Switched off.")
	a.powerOn = false
}

func (a *ac) CurrentStatus() bool {
	return a.powerOn
}

func (a *ac) Cost() int64 {
	return a.cost
}

func (a *ac) Name() string {
	return a.name
}

func NewAc() interfaces.Appliance {
	return &ac{
		powerOn: true,
		cost: constants.COST_AC,
		name:constants.AC,
	}
}