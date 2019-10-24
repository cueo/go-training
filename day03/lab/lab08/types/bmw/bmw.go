package bmw

import "errors"

type BMW struct {
	model    string
	year     uint16
	capacity int
	kind     string
}

func New(year uint16) (BMW, error) {
	var err error
	if year > 2019 {
		err = errors.New("year not in valid range")
	}
	v := BMW{
		model: "BMW X1",
		year: year,
		capacity: 4,
		kind: "Petrol",
	}
	return v, err
}

func (bmw BMW) Model() string {
	return bmw.model
}

func (bmw BMW) Year() uint16 {
	return bmw.year
}

func (bmw BMW) SeatingCapacity() int {
	return bmw.capacity
}

func (bmw BMW) Type() string {
	return bmw.kind
}
