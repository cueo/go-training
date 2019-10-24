package vehicle

type Vehicle interface {
	Model() string
	Year() uint16
	SeatingCapacity() int
	Type() string  // Electric or Self-Driving or Petrol or Diesel
}
