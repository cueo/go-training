package vehicle

type VehicleProperties struct {
	Model string
	Year uint16
	SeatingCapacity int
	Kind string
}

type Vehicle interface {
	Model() string
	Year() uint16
	SeatingCapacity() int
	Type() string //Electric or Self-Driving or Petrol or Diesel
}

func (vp VehicleProperties) String() string {
	return "VP string() method"
}