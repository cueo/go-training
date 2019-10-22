package main

import "strings"

func main() {
	city := "Patna"
	newCity := strings.Replace(city, "P", "p", -1)
	println(newCity)
	println(len(city))

	for i := 0; i < len(city); i++ {
		println(city[i], string(city[i]))
	}
}
