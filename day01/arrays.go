package main

import "fmt"

func main() {
	// int[] n => Java
	var numbers [10]int
	fmt.Println(numbers)
	// println(numbers)  // does not work ??

	numbers[0] = 100
	fmt.Println(numbers)

	// numbers[11] = 100  // compile time error

	langs := []string {"Go", "C", "Java", "Python", "Kotlin"}
	fmt.Println(langs)

	var arr = make([]int, 10)
	println(arr)
	fmt.Println(arr)

	// for each
	for index, value := range langs {
		println(index, value)
	}
}
