package main

import "fmt"

func main() {
	x := increment(1)
	fmt.Println(x)

	square, cube := compute(2)
	fmt.Println(square, cube)

	_, c := compute(3)
	fmt.Println(c)
}

//function with Named Return values
func operate(a int, b int) (product int, difference int) {
	product =  a * b
	difference = a - b
	return
}


func add(a int, b int) (sum int) {
	sum =  a + b
	return
}

func addNormal(a int, b int) int {
	sum :=  a + b
	return sum
}

func compute(number int) (int, int) {
	return number * number, number * number * number
}


func increment(number int) int {
	return number + 1
}