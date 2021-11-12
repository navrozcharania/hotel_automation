package algorithms

import (
	"errors"
	"fmt"
	"hotelElectricsSystem/common/constants"
	"hotelElectricsSystem/common/helpers"
	"hotelElectricsSystem/interfaces"
	"log"
)

type Algorithm struct {
	PermissibleMulitplier map[string]int64
}

func GetBasicAlgorithm() interfaces.Algorithm {
	return &Algorithm{PermissibleMulitplier: map[string]int64{constants.MAIN: 15, constants.SUB: 10}}
}
func (algo *Algorithm) Process(layout interfaces.Layout, floor int64, crdType string, crdNum int64, sensorInp string) error {
	// figure floor to work on
	// switch things on or off
	// re-balance things if required
	var workingFloor interfaces.Floor
	for _, currFloor := range layout.GetFloors() {
		if currFloor.StoreyNumber() == floor {
			workingFloor = currFloor
			break
		}
	}
	if workingFloor == nil {
		return errors.New("given floor number not available: " + fmt.Sprint(floor))
	}

	var workingCorridor interfaces.Corridor
	switch crdType {
	case constants.MAIN:
		log.Print("nothing to do as corridor type is main")
		return nil
	case constants.SUB:
		for _, currCorridor := range workingFloor.GetCorridors() {
			if currCorridor.Name() == crdType && currCorridor.CorridorNumber() == crdNum {
				workingCorridor = currCorridor
				break
			}
		}
		if workingCorridor == nil {
			return errors.New("given corridor not available: " + fmt.Sprint(crdType, crdNum))
		}
		if err := algo.actuateOnSensorInput(workingCorridor, sensorInp); err != nil {
			return err
		}
	}

	if err := algo.balanceFloorUsage(workingFloor, workingCorridor, sensorInp); err != nil {
		return err
	}

	helpers.PrintUsageOnThisFloor(workingFloor)
	return nil
}

func (algo *Algorithm) actuateOnSensorInput(workingCorridor interfaces.Corridor, sensorInp string) error {
	switch sensorInp {
	case constants.MOVEMENT:
		// check if light is already on, if yes then do nothing
		// if no then switch on light and see balance of the floor in case any fixing is needed then do that fixing
		for _, appliance := range workingCorridor.GetAppliances() {
			if appliance.Name() == constants.LIGHT {
				if !appliance.CurrentStatus() {
					appliance.SwitchOn()
				} else {
					log.Print("LIGHTS are alerady on.")
				}
			}
		}
	case constants.SILENCE:
		// check ac and light
		// if light is on then switch it off
		// if ac is on then exit
		// if ac is off then switch it on and make sure this ac doesn't get switched off in re-balancing
		for _, appliance := range workingCorridor.GetAppliances() {
			if appliance.Name() == constants.LIGHT {
				if appliance.CurrentStatus() {
					appliance.SwitchOff()
				} else {
					log.Print("LIGHTS are already off.")
				}
			}
			if appliance.Name() == constants.AC {
				if appliance.CurrentStatus() == false {
					appliance.SwitchOn()
				} else {
					log.Print("AC is already on.")
				}
			}
		}
	default:
		return errors.New("no procedure found for given sensor input: " + fmt.Sprint(sensorInp))
	}
	return nil
}

func (algo *Algorithm) balanceFloorUsage(workingFloor interfaces.Floor, workingCorridor interfaces.Corridor, sensorInp string) error {
	cost, err := helpers.GetCostOfGivenFloor(workingFloor)
	if err != nil {
		return err
	}

	permissibleValue := algo.getPermissibleValue(workingFloor)

	// case where cost is equal to permissible Value do nothing
	if cost == permissibleValue {
		log.Print("cost is within limits not re-balancing")
		return nil
	}

	// case where cost smaller than permissible Value then try to switch somethings on
	if cost < permissibleValue {
		if err = algo.positiveBalanceFloorUsage(workingFloor, workingCorridor, permissibleValue, permissibleValue-cost); err != nil {
			return err
		}
		return nil
	}

	// case where cost greater than permissible Value then try to switch somethings off
	if err = algo.negativeBalanceFloorUsage(workingFloor, workingCorridor, permissibleValue, cost-permissibleValue, sensorInp); err != nil {
		return err
	}

	return nil
}

func (algo *Algorithm) positiveBalanceFloorUsage(floor interfaces.Floor, corridor interfaces.Corridor, permissible, balance int64) error {
	if balance < 10 {
		// if balance is smaller than required then no need for processing as ac cannot be swithced on
		return nil
	}
	// no acs to switch on
	// keep switching on acs till balance is covered
	// find any subCorridor which has an ac switched off and switch it on
	for _, currCorridor := range floor.GetCorridors() {
		if currCorridor.Name() == constants.SUB {
			for _, currAppliance := range currCorridor.GetAppliances() {
				if currAppliance.Name() == constants.AC {
					currAppliance.SwitchOn()
					break
				}
			}
			break
		}
	}

	cost, err := helpers.GetCostOfGivenFloor(floor)
	if err != nil {
		return err
	}

	// if cost and permissible still have a difference of an AC then try switching another ac on
	if permissible-cost >= 10 && balance != permissible-cost {
		return algo.positiveBalanceFloorUsage(floor, corridor, permissible, permissible-cost)
	}
	return nil
}

func (algo *Algorithm) negativeBalanceFloorUsage(floor interfaces.Floor, corridor interfaces.Corridor, permissible, balance int64, sensorInp string) error {
	/*
		1. one which was not switched on recently and lights are off too.
		2. if 1 exhausts then one where light is on we'll switch off ac.
		3. ac that was recently switched on.
		(case 3 should never arise as the constraint always covers having the ac or light switched on in one place)
	*/
	for _, currCorridor := range floor.GetCorridors() {
		if currCorridor.Name() == constants.MAIN {
			// do nothing in case it's a main corridor
			continue
		}

		// case of movement is on so we can switch ac off for anyone
		switch sensorInp {
		case constants.MOVEMENT:
			for _, currAppliance := range currCorridor.GetAppliances() {
				if currAppliance.Name() == constants.AC {
					currAppliance.SwitchOff()
					return nil
				}
			}
		case constants.SILENCE:
			// case of silence we can switch ac off for any except current
			if currCorridor.CorridorNumber() != corridor.CorridorNumber() {
				for _, currAppliance := range currCorridor.GetAppliances() {
					if currAppliance.Name() == constants.AC {
						currAppliance.SwitchOff()
						return nil
					}
				}
			}
		}
	}

	cost, err := helpers.GetCostOfGivenFloor(floor)
	if err != nil {
		return err
	}

	// if cost still exceeds then switch some more things off
	if cost > permissible {
		return algo.negativeBalanceFloorUsage(floor, corridor, permissible, cost-permissible, sensorInp)
	}

	return nil
}

func (algo *Algorithm) getPermissibleValue(floor interfaces.Floor) int64 {
	countCorridors := helpers.NumberOfEachCorridors(floor)
	permissibleValue := algo.PermissibleMulitplier[constants.MAIN]*countCorridors[constants.MAIN] +
		algo.PermissibleMulitplier[constants.SUB]*countCorridors[constants.SUB]
	return permissibleValue
}
