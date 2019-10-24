package bmw

import (
	"../vehicle"
	"errors"
	"time"
)

type bmw struct {
	vehicle.VehicleProperties
}

func (bmwObj bmw) Model() string{
	return bmwObj.VehicleProperties.Model
}

func (bmwObj bmw) Year() uint16{
	return bmwObj.VehicleProperties.Year
}

func (bmwObj bmw) SeatingCapacity() int{
	return bmwObj.VehicleProperties.SeatingCapacity
}

func (bmwObj bmw) Type() string{
	return bmwObj.VehicleProperties.Kind
}


func New(year uint16) (bmw, error) {
	if int(year) > time.Now().Year() {
		er := errors.New("Year cannot be greater than 2019")
		return bmw{}, er
	}

	return bmw{ vehicle.VehicleProperties{
		Model:           "BMW",
		Year:            year,
		SeatingCapacity: 5,
		Kind:            "Electric",
	}}, nil
}

func (bmwObj bmw) String() string {
	//message := fmt.Sprintf("BMW: (%v), Seating Capacity - %v", bmwObj.VehicleProperties.Year, bmwObj.VehicleProperties.SeatingCapacity)
	return bmwObj.VehicleProperties.String()
}

