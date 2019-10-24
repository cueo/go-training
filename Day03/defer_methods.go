package main

import (
	"fmt"
)

//setTimeout

func eat() {
	fmt.Println("eat")
}

func main() {
	defer eat()
	fmt.Println("End of main")
}








