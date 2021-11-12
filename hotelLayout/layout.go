package hotelLayout

import (
	"hotelElectricsSystem/interfaces"
)

type HotelLayout struct {
	Floors []interfaces.Floor
}

func (layout *HotelLayout) GetFloors() []interfaces.Floor {
	return layout.Floors
}

func (layout *HotelLayout) AddFloor(floor interfaces.Floor) {
	layout.Floors = append(layout.Floors, floor)
}

func BuildHotelLayout(floors int64, mainCorridors, subCorridors []int64) (interfaces.Layout,error) {
	hotelLayout := &HotelLayout{Floors: make([]interfaces.Floor, floors)}
	for index, _ := range mainCorridors {
		hotelLayout.Floors[index] = NewFloor(mainCorridors[index], subCorridors[index], int64(index+1))
	}
	return hotelLayout, nil
}
