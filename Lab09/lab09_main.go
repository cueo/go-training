package main

import (
	"fmt"

	"./types/bmw"
	"./types/ford"
	"./types/honda"
	"./types/vehicle"
)

func main() {
	garage := make([]vehicle.Vehicle, 3)

	garage[0], _ = ford.New(2001)
	garage[1], _ = honda.New(2010)
	garage[2], _ = bmw.New(2019)

	for _, v := range garage {
		//calls String() string; This is similar to our toString() method
		fmt.Println(v) //Prints model, year, type, seatingcapacity
	}
}
