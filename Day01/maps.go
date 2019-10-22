package main

import "fmt"

func main() {
	//var numbers [10]int
	//delete(numbers, 0)

	var dictionary map [int]string = make(map[int]string)
	fmt.Println(dictionary)

	//unordered
	var cities map [string] string = make(map[string]string, 10)
	cities["KA"] = "Bangalore"
	cities["TN"] = "Chennai"
	cities["KL"] = "Kochi"
	//cities["AP"] = 2
	fmt.Println(cities["KL"])
	delete(cities, "TN")

	for key, value := range cities  {
		fmt.Println(key, value)
	}
}
