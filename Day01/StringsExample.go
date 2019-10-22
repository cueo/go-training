package main

import (
	"fmt"
	"strings"
)

func main() {
	city := "Bangalore"
	newCity := strings.Replace(city, "a", "A", -1)
	fmt.Println(len(city), newCity)
	for i := 0; i < len(city) ; i++  {
		fmt.Println(string(city[i]))
	}
}
