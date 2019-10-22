package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//Slicing creates a shallow copy
	arr1 := numbers[:2]
	fmt.Println(arr1) //1, 2

	arr2 := numbers[4:7]
	fmt.Println(arr2)

	arr3 := numbers[0:2]
	arr3[0] = 100
	arr3[1] = 200
	fmt.Println(arr3, numbers)

	deepCopy := make([]int, 3)
	copy(deepCopy, numbers[4:7]) //deep copy
	deepCopy[0] = 500
	deepCopy[1] = 600
	deepCopy[2] = 700

	deepCopy  = append(deepCopy, 10000, 20000)
	anotherCopy  := append(deepCopy, 10000, 20000)
	deepCopy[0] = 8098
	fmt.Println("*****", anotherCopy, deepCopy, " ----")
	deepCopy = append(deepCopy, numbers...)//like spread operator in ES6
	fmt.Println(deepCopy, numbers)







}
