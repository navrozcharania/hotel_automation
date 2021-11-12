package common

import (
	"errors"
	"fmt"
	"hotelElectricsSystem/common/constants"
)

var SensorInputToIntegerMap = map[string]int64{
	constants.MOVEMENT:1,
	constants.SILENCE: 0,
}

func GetActionFromSensorInput(sensorInp string) (int64, error) {
	var action int64
	var ok bool
	if action, ok = SensorInputToIntegerMap[sensorInp]; !ok {
		return 0, errors.New("no mapping found for given sensor input: " + fmt.Sprint(sensorInp))
	}
	return action, nil
}