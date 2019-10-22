package main

import "fmt"

type Currency float64

const PI float64 = 3.14

var four = 4

func main() {
	types1()
	types2()
}

func types1() {
	println("=====")
	var dollar Currency = 60.9
	fmt.Printf("$%v\n", dollar)

	println("Value of pi =", PI)

	var x = 10
	fmt.Printf("%T\n", x) // print data type of x
	println("=====")
}

func types2() {
	println("=====")
	var one int = 1
	var two = 2
	three := 3

	var five, six = 5, 6
	seven, eight, nine := 7, 8, 9

	println(one, two, three, four, five, six, seven, eight, nine)
	println("=====")
}
