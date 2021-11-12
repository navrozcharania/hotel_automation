package main

import (
	"hotelElectricsSystem/algorithms"
	"hotelElectricsSystem/controller"
	"hotelElectricsSystem/hotelLayout"
	"hotelElectricsSystem/input_files"
	"hotelElectricsSystem/interfaces"
	"log"
)

const INPUT_FILE_PATH = "C:\\Users\\Navroz\\go\\src\\hotelElectricsSystem\\input_files\\"

func main() {
	// Initializing the layout and other things happen here
	floors, mainCrdList, subCrdList, _ := input_files.ReadLayout(INPUT_FILE_PATH + "layout.txt")
	var hotel interfaces.Layout
	hotel,_  = hotelLayout.BuildHotelLayout(floors, mainCrdList, subCrdList)
	var controllerInstance interfaces.Controller
	controllerInstance = controller.NewController(algorithms.GetBasicAlgorithm())
	controllerInstance.InitializeAppliances(hotel)
	controllerInstance.PrintHotelUsage(hotel)
	controllerInstance.PrintCostPerFloor(hotel)

	// Events portion starts here
	queryFloor, queryCrdType, queryCrdNum, querySensorInput, _ := input_files.ReadEvents(INPUT_FILE_PATH + "events.txt")

	for idx,_ := range queryFloor {
		err := controllerInstance.RaiseSensorEvent(hotel, queryFloor[idx], queryCrdType[idx], queryCrdNum[idx], querySensorInput[idx])
		if err != nil {
			log.Fatal(err)
		}
	}
	controllerInstance.PrintCostPerFloor(hotel)
}