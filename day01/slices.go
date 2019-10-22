package main

import "fmt"

func main() {
	numbers := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < 10; i++ {
		numbers[i] = i+1
	}
	fmt.Println(numbers)

	// slicing
	arr1 := numbers[:2]
	fmt.Println(arr1)

	arr2 := numbers[4:7]
	fmt.Println(arr2)

	arr3 := numbers[0:2]
	arr3[0] = 6
	arr3[1] = 9
	fmt.Println(arr3, numbers)
	arr3 = append(arr3, 100, 200)
	fmt.Println(arr3, numbers)

	deepCopy := make([]int, 3)
	copy(deepCopy, numbers[0:3])
	deepCopy[0] = 9
	deepCopy[1] = 6
	fmt.Println(deepCopy, numbers)
	deepCopy = append(deepCopy, numbers...)
	fmt.Println(deepCopy)
}
