package ford

import "errors"

type Ford struct {
	model    string
	year     uint16
	capacity int
	kind     string
}

func New(year uint16) (Ford, error) {
	var err error
	if year > 2019 {
		err = errors.New("year not in valid range")
	}
	v := Ford{
		model: "Ford",
		year: year,
	}
	return v, err
}

func (ford Ford) Model() string {
	return ford.model
}

func (ford Ford) Year() uint16 {
	return ford.year
}

func (ford Ford) SeatingCapacity() int {
	return ford.capacity
}

func (ford Ford) Type() string {
	return ford.kind
}
