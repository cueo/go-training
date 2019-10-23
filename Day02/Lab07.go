package main

import "fmt"

var numbers = []int{101, 22, 43, 14, 5, 906, 310, 561, 84, 23, 100}

type NumberCheck func(int) bool
type NumberOperation func(int) int

func computeTotal(operation NumberOperation) (total int) {
	for _,value := range numbers {
		total += operation(value)
	}
	return
}

func find(check NumberCheck) (result []int) {
	for _, value := range numbers {
		if check(value) {
			result = append(result, value)
		}
	}
	return
}

func main() {
	squares := computeTotal(func(num int) int { return num * num });
	cubes := computeTotal(func(num int) int { return num * num * num });
	fmt.Println(squares, cubes)

	fmt.Println("Even list", find(func(num int) bool { return num % 2 == 0 }))

	odd := func(num int) bool {
		return num % 2 == 1
	}
	fmt.Println("Odd list", find(odd))

	divBy5 := func(num int) bool {
		return num % 5 == 0
	}
	fmt.Println("Div By 5 list", find(divBy5))
}


