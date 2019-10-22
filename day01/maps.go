package main

import "fmt"

func main() {
	var dictionary = make(map[int]string)
	fmt.Println(dictionary)

	var cities = make(map[string]string)
	cities["KA"] = "Bangalore"
	cities["BR"] = "Patna"
	fmt.Println(cities)
	for key, value := range cities {
		fmt.Println(key, value)
	}
	delete(cities, "KA")
	fmt.Println(cities)
}
