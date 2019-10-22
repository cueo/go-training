package main

import "fmt"

func main() {
	//int[] numbers
	const x  int = 10
	var numbers [x]int
	numbers[0] = 100
	//compile-time error
	//numbers[11] = 200
	fmt.Println(numbers, len(numbers))

	langs := []string{ "Java", "C#", "Ruby", "Go", "ES6" }
	//langs := [5]string{ "Java", "C#", "Ruby", "Go", "ES6" }
	fmt.Println(langs)
	for i := 0; i < len(langs) ; i++  {
		fmt.Println(langs[i])
	}
	size := 10
	var arr1 = make([]int, size)
	fmt.Println(arr1)
	println(arr1)
}
