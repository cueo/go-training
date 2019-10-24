package ford

import (
	"errors"
	"fmt"
	"time"
)

type ford struct {
	model string
	year uint16
	seatingCapacity int
	kind string
}

func (fordObj ford) Model() string{
	return fordObj.model
}

func (fordObj ford) Year() uint16{
	return fordObj.year
}

func (fordObj ford) SeatingCapacity() int{
	return fordObj.seatingCapacity
}

func (fordObj ford) Type() string{
	return fordObj.kind
}


func New(year uint16) (ford, error) {
	if int(year) > time.Now().Year() {
		er := errors.New("Year cannot be greater than 2019")
		return ford{}, er
	}
	return ford{ model: "FORD", seatingCapacity: 6, kind: "Diesel", year: year }, nil
}

func (fordObj ford) String() string {
	message := fmt.Sprintf("FORD: (%v), Seating Capacity - %v", fordObj.year, fordObj.seatingCapacity)
	return message
}