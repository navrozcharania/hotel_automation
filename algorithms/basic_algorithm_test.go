package algorithms

import (
	"hotelElectricsSystem/common"
	"hotelElectricsSystem/common/constants"
	"hotelElectricsSystem/hotelLayout"
	"testing"
)

func TestAlgorithm_Process(t *testing.T) {
	layout, err := hotelLayout.BuildHotelLayout(2, []int64{1,1}, []int64{1,1})
	common.AssertNil(t, err)
	floor := layout.GetFloors()[0]
	corridorMain, corridorSub := floor.GetCorridors()[0], floor.GetCorridors()[1]
	for _, appliance := range corridorSub.GetAppliances() {
		if appliance.Name() == constants.AC {
			common.AssertTrue(t, appliance.CurrentStatus())
		}
		if appliance.Name() == constants.LIGHT {
			common.AssertTrue(t, appliance.CurrentStatus())
		}
	}
	err = GetBasicAlgorithm().Process(layout, 1, "SUB", 1 , "MOVEMENT")
	common.AssertNil(t, err)
	for _, appliance := range corridorSub.GetAppliances() {
		if appliance.Name() == constants.AC {
			common.AssertFalse(t, appliance.CurrentStatus())
		}
		if appliance.Name() == constants.LIGHT {
			common.AssertTrue(t, appliance.CurrentStatus())
		}
	}

	for _, appliance := range corridorMain.GetAppliances() {
		if appliance.Name() == constants.AC {
			common.AssertTrue(t, appliance.CurrentStatus())
		}
		if appliance.Name() == constants.LIGHT {
			common.AssertTrue(t, appliance.CurrentStatus())
		}
	}
}
