package main

import "fmt"

func main() {
	sum := computeSum("please", 1, 2, 3, 4)
	fmt.Println(sum)
}

func computeSum(message string, numbers ...int) (sum int) {
	for _, value := range numbers {
		sum += value
	}
	return
}
