package honda

import (
	"errors"
	"fmt"
)

type Honda struct {
	model    string
	year     uint16
	capacity int
	kind     string
}

func New(year uint16) (Honda, error) {
	var err error
	if year > 2019 {
		err = errors.New("year not in valid range")
	}
	v := Honda{
		model: "City",
		year: year,
		capacity: 4,
		kind: "Diesel",
	}
	return v, err
}

func (honda Honda) Model() string {
	return honda.model
}

func (honda Honda) Year() uint16 {
	return honda.year
}

func (honda Honda) SeatingCapacity() int {
	return honda.capacity
}

func (honda Honda) Type() string {
	return honda.kind
}

// similar to `toString()` in Java
func (honda Honda) String() string {
	return fmt.Sprintf("Honda %s, made in year %d seats %d people and drives on %s. Enjoy your ride!", honda.model, honda.year, honda.capacity, honda.kind)
}
