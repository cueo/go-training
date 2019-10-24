package bmw

import (
	"errors"
	"fmt"
	"time"
)

type bmw struct {
	model string
	year uint16
	seatingCapacity int
	kind string
}

func (bmwObj bmw) Model() string{
	return bmwObj.model
}

func (bmwObj bmw) Year() uint16{
	return bmwObj.year
}

func (bmwObj bmw) SeatingCapacity() int{
	return bmwObj.seatingCapacity
}

func (bmwObj bmw) Type() string{
	return bmwObj.kind
}


func New(year uint16) (bmw, error) {
	if int(year) > time.Now().Year() {
		er := errors.New("Year cannot be greater than 2019")
		return bmw{}, er
	}
	return bmw{ model: "BMW", seatingCapacity: 5, kind: "Petrol", year: year }, nil
}

func (bmwObj bmw) String() string {
	message := fmt.Sprintf("BMW: (%v), Seating Capacity - %v", bmwObj.year, bmwObj.seatingCapacity)
	return message
}

