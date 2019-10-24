package main

import "fmt"

type Rectangle struct {
	length int
	breadth int
}
type Square struct {
	side int
}
type Circle struct {
	radius int
}
type Shape interface {
	area() int
	perimeter() int
}




func (c Circle) area() int {
	return c.radius * c.radius * 3 //deliberately ignored 3.14
}
func (s Square) area() int {
	return s.side * s.side
}
func (r Rectangle) area() int {
	return r.length * r.breadth
}

func (c Circle) perimeter() int {
	return c.radius * 2 * 3 //deliberately ignored 3.14
}
func (s Square) perimeter() int {
	return s.side * 4
}
func (r Rectangle) perimeter() int {
	return 2 * r.length * r.breadth
}

func calculate(shape Shape) {
	fmt.Println("Area: ", shape.area())
}

func doSomething(a interface{ read() string }) {

}
func (c Circle) read() string {
	return "circle"
}

func myPrintln(a ... interface{}) {

}

//instanceof keyword in Java
//interface{} is like a marker interface in Java; no methods at all;
//example is java.io.Serializable
func crazy(a interface{}) {
	switch a.(type) {
		case int :
			fmt.Println(a)
	case Square:
		//var sq Square = Square(a)
		var sq Square = a.(Square)
		fmt.Println(sq.side)
		fmt.Println(a.(Square).side)
	case Circle:
		fmt.Println(a.(Circle).radius)
	case string:
		fmt.Println("string", a)
	}
}

func main() {
	myPrintln(12, 12, 2.34)
	crazy(12)
	crazy("cool")
	crazy(Square{12})
	crazy([]int{})

	ci := Circle { 2 }
	sq := Square { 2 }
	re := Rectangle{ 2, 2}
	calculate(ci)
	calculate(sq)
	calculate(re)
	doSomething(ci)
}
