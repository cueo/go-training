package main

import "fmt"

type Currency float64
const PI float64 = 3.14

func main() {

	var dollar Currency = 71.15
	fmt.Printf("$%v \n", dollar)

	//int, bool, string
	println("builtin package")

	var unsignedInt uint32

	var b1 byte
	var javaInt int
	var javaBool bool
	var long int64
	var float float32
	var double float64

	var char rune
	println(unsignedInt, b1, javaInt, javaBool, long, float, double, char)
	//0 0 false 0 +0.000000e+000 +0.000000e+000 0
}
