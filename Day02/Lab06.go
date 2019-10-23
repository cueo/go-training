package main

import "fmt"

type MyNumbers []int

func main() {
	var myCollection MyNumbers = []int{1, 20, 40, 4, 53, 87 }
	fmt.Println("Max", myCollection.max())
	fmt.Println("Min", myCollection.min())
	fmt.Println("Sum", myCollection.sum())
}

func (coll MyNumbers) max() int {
	largest := coll[0]
	for _, value := range coll {
		if value > largest {
			largest = value
		}
	}
	return largest
}

func (coll MyNumbers) sum() int {
	total := coll[0]
	for _, value := range coll {
		total += value
	}
	return total
}


func (coll MyNumbers) min() int {
	smallest := coll[0]
	for _, value := range coll {
		if value < smallest {
			smallest = value
		}
	}
	return smallest
}