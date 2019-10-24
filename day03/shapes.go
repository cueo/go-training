package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float32
}

type Square struct {
	side float32
}

type Rectangle struct {
	length  float32
	breadth float32
}

type Circle struct {
	radius float32
}

func (s Square) area() (a float32) {
	a = s.side * s.side
	return
}

func (r Rectangle) area() (a float32) {
	a = r.length * r.breadth
	return
}

func (c Circle) area() (a float32) {
	a = math.Pi * c.radius * c.radius
	return
}

func main() {
	c := Circle{radius: 10}
	s := Square{side: 10}
	r := Rectangle{
		length:  10,
		breadth: 5,
	}
	calculateAreas(c, s, r)
	calculateAreasWithSwitch(c, s, r)
}

func calculateAreas(shapes ...Shape) {
	for _, shape := range shapes {
		fmt.Printf("Area for %T = %f\n", shape, shape.area())
	}
}

func calculateAreasWithSwitch(shapes ...interface{}) {
	for _, shape := range shapes {
		switch shape.(type) {
		case Square:
			fmt.Printf("Area for square = %f\n", shape.(Square).area())
		case Rectangle:
			fmt.Printf("Area for rectangle = %f\n", shape.(Rectangle).area())
		case Circle:
			fmt.Printf("Area for circle = %f\n", shape.(Circle).area())
		}
	}
}

// equivalent to void* in C++
func anonymousInterface(a interface{ read() string }) {

}
