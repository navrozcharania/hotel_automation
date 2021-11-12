package hotelLayout

import (
	"hotelElectricsSystem/common/constants"
	"hotelElectricsSystem/hotelLayout/appliance"
	"hotelElectricsSystem/interfaces"
)

type corridor struct {
	ApplianceSetup []interfaces.Appliance
	name string
	number int64
}

func (crd *corridor) CorridorNumber() int64 {
	return crd.number
}

func (crd *corridor) GetAppliances() []interfaces.Appliance {
	return crd.ApplianceSetup
}

func (crd *corridor) AddAppliance(appliance interfaces.Appliance) {
	crd.ApplianceSetup = append(crd.ApplianceSetup, appliance)
}

func (crd *corridor) Name() string {
	return crd.name
}

func NewMainCorridor(number int64) interfaces.Corridor {
	corridor := &corridor{
		ApplianceSetup: make([]interfaces.Appliance, 0),
		name:           constants.MAIN,
		number: number,
	}
	corridor.AddAppliance(appliance.NewAc())
	corridor.AddAppliance(appliance.NewLight())
	return corridor
}

// Having different functions for init corridor and sub corridor as we may want to add some extra processing while initializing
// one thing I could do is switch light/ac in off mode while adding to corridor
func NewSubCorridor(number int64) interfaces.Corridor {
	corridor := &corridor{
		ApplianceSetup: make([]interfaces.Appliance, 0),
		name:           constants.SUB,
		number: number,
	}
	corridor.AddAppliance(appliance.NewAc())
	corridor.AddAppliance(appliance.NewLight())
	return corridor
}