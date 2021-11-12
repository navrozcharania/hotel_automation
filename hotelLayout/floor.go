package hotelLayout

import "hotelElectricsSystem/interfaces"

type Floor struct {
	Corridors []interfaces.Corridor
	Number int64
}

func (flr *Floor) GetCorridors() []interfaces.Corridor {
	return flr.Corridors
}

func (flr *Floor) AddCorridor(c interfaces.Corridor) {
	flr.Corridors = append(flr.Corridors, c)
}

func (flr *Floor) StoreyNumber() int64 {
	return flr.Number
}

func NewFloor(numMain, numSub, storeyNumber int64) interfaces.Floor {
	floor := &Floor{Corridors: make([]interfaces.Corridor, 0), Number: storeyNumber}
	for i := numMain; i>0;i-- {
		floor.AddCorridor(NewMainCorridor(i))
	}
	for i := numSub; i>0;i-- {
		floor.AddCorridor(NewSubCorridor(i))
	}
	return floor
}



