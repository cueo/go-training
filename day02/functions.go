package main

import "fmt"

func main() {
	// multiple return values
	square, cube := compute(2)
	fmt.Println(square, cube)

	// named return values
	sum, diff := operate(5, 6)
	fmt.Println(sum, diff)
}

func operate(x int, y int) (sum int, diff int) {
	sum = x + y
	diff = x - y
	return
}

func compute(x int) (int, int) {
	return x*x, x*x*x
}
