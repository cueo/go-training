package main

import (
	"./intuit"
	"fmt"
)

func main() {
	fmt.Println(intuit.PI)
	intuit.Add(12, 23)

	data := intuit.Data{ "somevalue" }
	fmt.Println(data.Value)

	//intuit.subtract(12, 12)
}
