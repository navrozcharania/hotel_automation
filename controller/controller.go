package controller

import (
	"fmt"
	"hotelElectricsSystem/common/constants"
	"hotelElectricsSystem/common/helpers"
	"hotelElectricsSystem/interfaces"
	"log"
)

type Controller struct {
	Algorithm interfaces.Algorithm
}

func (cntrlr *Controller) RaiseSensorEvent(h interfaces.Layout ,floor int64, crdType string, crdNum int64, sensorInp string) error {
	log.Printf("Received event from Floor: %d  Corridor: %s %d SensorInput: %s",floor, crdType, crdNum, sensorInp)
	return cntrlr.Algorithm.Process(h, floor, crdType, crdNum, sensorInp)
}

func NewController(algorithm interfaces.Algorithm) interfaces.Controller {
	return &Controller{
		Algorithm: algorithm,
	}
}

func (cntrlr *Controller) InitializeAppliances(h interfaces.Layout) {
	//log.Print("Switching on all the LIGHTS and ACs except lights from sub corridors.")
	for _, floor := range h.GetFloors() {
		for _, corridor := range floor.GetCorridors() {
			if corridor.Name() == constants.SUB {
				for _,appliance := range corridor.GetAppliances() {
					if appliance.Name() == constants.LIGHT {
						appliance.SwitchOff()
					}
				}
			}
		}
	}
}

func (cntrlr *Controller) PrintHotelUsage(h interfaces.Layout) {
	for _, floor := range h.GetFloors() {
		helpers.PrintUsageOnThisFloor(floor)
	}
}

func (cntrlr *Controller) PrintCostPerFloor(h interfaces.Layout) int64 {
	hotelTotal := int64(0)
	for _, floor := range h.GetFloors() {
		floorTotal := helpers.PrintCostOfThisFloor(floor)
		hotelTotal += floorTotal
	}
	fmt.Println("Hotel's current number of units: ", hotelTotal)
	return hotelTotal
}

