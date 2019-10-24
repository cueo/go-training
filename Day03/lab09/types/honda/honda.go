package honda

import (
	"errors"
	"fmt"
	"time"
)

type honda struct {
	model string
	year uint16
	seatingCapacity int
	kind string
}

func (hondaObj honda) Model() string{
	return hondaObj.model
}

func (hondaObj honda) Year() uint16{
	return hondaObj.year
}

func (hondaObj honda) SeatingCapacity() int{
	return hondaObj.seatingCapacity
}

func (hondaObj honda) Type() string{
	return hondaObj.kind
}


func New(year uint16) (honda, error) {
	if int(year) > time.Now().Year() {
		er := errors.New("Year cannot be greater than 2019")
		return honda{}, er
	}
	return honda{ model: "Honda", seatingCapacity: 4, kind: "Electric", year: year }, nil
}

func (hondaObj honda) String() string {
	message := fmt.Sprintf("HONDA: (%v), Seating Capacity - %v", hondaObj.year, hondaObj.seatingCapacity)
	return message
}