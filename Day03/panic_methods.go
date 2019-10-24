package main

import "fmt"

func main() {
	fmt.Println("Inside main")
	num := 10
	den := 0
	if den == 0 {
		panic("Oops! You're trying to divide by zero")
	} else {
		div := num / den
		fmt.Println(div)
	}
	fmt.Println("End of main")
}
