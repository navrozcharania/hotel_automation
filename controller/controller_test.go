package controller

import (
	"hotelElectricsSystem/common"
	"hotelElectricsSystem/common/constants"
	"hotelElectricsSystem/hotelLayout"
	"hotelElectricsSystem/interfaces"
	"testing"
)

// Tests around interfaces of controller

func TestController_RaiseSensorEvent(t *testing.T) {
	newController := NewController(getShortCircuitAlgorithm())
	layout, err := hotelLayout.BuildHotelLayout(2, []int64{1,1}, []int64{1,1})
	newController.InitializeAppliances(layout)
	floor := layout.GetFloors()[0]
	corridorMain, corridorSub := floor.GetCorridors()[0], floor.GetCorridors()[1]
	common.AssertNil(t, err)
	for _, appliance := range corridorSub.GetAppliances() {
		if appliance.Name() == constants.AC {
			common.AssertTrue(t, appliance.CurrentStatus())
		}
		if appliance.Name() == constants.LIGHT {
			common.AssertFalse(t, appliance.CurrentStatus())
		}
	}
	err = newController.RaiseSensorEvent(layout, 1, "SUB", 1 , "MOVEMENT")
	common.AssertNil(t, err)
	for _, appliance := range corridorSub.GetAppliances() {
		if appliance.Name() == constants.AC {
			common.AssertFalse(t, appliance.CurrentStatus())
		}
		if appliance.Name() == constants.LIGHT {
			common.AssertFalse(t, appliance.CurrentStatus())
		}
	}

	for _, appliance := range corridorMain.GetAppliances() {
		if appliance.Name() == constants.AC {
			common.AssertFalse(t, appliance.CurrentStatus())
		}
		if appliance.Name() == constants.LIGHT {
			common.AssertFalse(t, appliance.CurrentStatus())
		}
	}

}

func TestController_InitializeAppliances(t *testing.T) {
	newController := NewController(nil)
	layout, err := hotelLayout.BuildHotelLayout(2, []int64{1,1}, []int64{1,1})
	common.AssertNil(t, err)
	floor := layout.GetFloors()[0]
	_, corridorSub := floor.GetCorridors()[0], floor.GetCorridors()[1]
	common.AssertNil(t, err)
	for _, appliance := range corridorSub.GetAppliances() {
		if appliance.Name() == constants.AC {
			common.AssertTrue(t, appliance.CurrentStatus())
		}
		if appliance.Name() == constants.LIGHT {
			common.AssertTrue(t, appliance.CurrentStatus())
		}
	}
	newController.InitializeAppliances(layout)
	for _, appliance := range corridorSub.GetAppliances() {
		if appliance.Name() == constants.AC {
			common.AssertTrue(t, appliance.CurrentStatus())
		}
		if appliance.Name() == constants.LIGHT {
			common.AssertFalse(t, appliance.CurrentStatus())
		}
	}
}

func TestController_PrintHotelUsage(t *testing.T) {
	newController := NewController(nil)
	layout, err := hotelLayout.BuildHotelLayout(2, []int64{1,1}, []int64{1,1})
	common.AssertNil(t, err)
	newController.PrintHotelUsage(layout)
}

func TestController_PrintCostPerFloor(t *testing.T) {
	newController := NewController(nil)
	layout, err := hotelLayout.BuildHotelLayout(2, []int64{1,1}, []int64{1,1})
	common.AssertNil(t, err)
	newController.PrintHotelUsage(layout)
}


func getShortCircuitAlgorithm() interfaces.Algorithm {
	return &shortCircuit{}
}
type shortCircuit struct {
}
func (s *shortCircuit)Process(layout interfaces.Layout, a int64, b string, c int64, d string) error {
	for _, floor := range layout.GetFloors() {
		for _, corridor := range floor.GetCorridors() {
			for _, appliance := range corridor.GetAppliances() {
				appliance.SwitchOff()
			}
		}
	}
	return nil
}