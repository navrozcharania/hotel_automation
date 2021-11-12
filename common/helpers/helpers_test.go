package helpers

import (
	"hotelElectricsSystem/common"
	"hotelElectricsSystem/common/constants"
	"hotelElectricsSystem/hotelLayout"
	"testing"
)

func Test_PrintCostOfThisFloor(t *testing.T) {
	layout, _ := hotelLayout.BuildHotelLayout(2, []int64{1,1}, []int64{1,1})
	PrintCostOfThisFloor(layout.GetFloors()[0])
}

func Test_PrintUsageOnThisFloor(t *testing.T) {

	layout, _ := hotelLayout.BuildHotelLayout(2, []int64{1,1}, []int64{1,1})

	PrintUsageOnThisFloor(layout.GetFloors()[0])
}

func Test_GetCostOfGivenFloor(t *testing.T) {

	layout, _ := hotelLayout.BuildHotelLayout(2, []int64{1,1}, []int64{1,1})

	val, err := GetCostOfGivenFloor(layout.GetFloors()[0])
	common.AssertNil(t, err)
	common.AssertEqualInt64(t, 30 , val)
}

func Test_NumberOfEachCorridors(t *testing.T) {

	layout, _ := hotelLayout.BuildHotelLayout(2, []int64{1,1}, []int64{1,1})

	val := NumberOfEachCorridors(layout.GetFloors()[0])
	common.AssertEqualInt64(t, val[constants.SUB], 1)
	common.AssertEqualInt64(t, val[constants.MAIN], 1)
}