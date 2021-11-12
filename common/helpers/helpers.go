package helpers

import (
	"fmt"
	"hotelElectricsSystem/common/constants"
	"hotelElectricsSystem/interfaces"
)

func GetCostOfGivenFloor(floor interfaces.Floor) (int64, error) {
	floorTotal := int64(0)
	for _, corridor := range floor.GetCorridors() {
		for _, appliance := range corridor.GetAppliances() {
			if appliance.CurrentStatus() {
				floorTotal += appliance.Cost()
			}
		}
	}
	return floorTotal, nil
}

func PrintCostOfThisFloor(floor interfaces.Floor) int64 {
	floorTotal := int64(0)
	for _, corridor := range floor.GetCorridors() {
		for _,appliance := range corridor.GetAppliances() {
			if appliance.CurrentStatus() {
				floorTotal += appliance.Cost()
			}
		}
	}
	fmt.Println("Floor number:", floor.StoreyNumber()," Current Units: ",floorTotal)
	fmt.Println("\n---------------------------------------------------")
	return floorTotal
}


func PrintUsageOnThisFloor(floor interfaces.Floor) {
	fmt.Println("Floor number:", floor.StoreyNumber())
	for idx, corridor := range floor.GetCorridors() {
		fmt.Print("Corridor ", corridor.Name(), " ", idx)
		for _, appliance := range corridor.GetAppliances() {
			applianceStatus := "OFF"
			if appliance.CurrentStatus() {
				applianceStatus = "ON"
			}
			fmt.Print("  ", appliance.Name(), ":", applianceStatus)
		}
		fmt.Println()
	}
	fmt.Println("\n---------------------------------------------------")
}

func NumberOfEachCorridors(floor interfaces.Floor) map[string]int64 {
	countMap := map[string]int64{constants.MAIN: 0, constants.SUB: 0}
	for _, corridor := range floor.GetCorridors() {
		switch corridor.Name() {
		case constants.MAIN:
			countMap[constants.MAIN] += 1
		case constants.SUB:
			countMap[constants.SUB] += 1
		}
	}
	return countMap
}