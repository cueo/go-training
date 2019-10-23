package main

import "fmt"

func main() {
	x := increment(1)
	fmt.Println(x)

	square, cube := compute(2)
	fmt.Println(square, cube)

	_, c := compute(3)
	fmt.Println(c)
	fmt.Println(sum("total", 1, 2, 3, 4, 5))
}

//variadic functions
//... ellipsis argument should be the last in the list
func sum(message string, numbers ...int) (total int) {
	for _, value := range numbers  {
		total += value
	}
	return
	//fmt.Println("cool")
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