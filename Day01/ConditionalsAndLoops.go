package main

import "fmt"


//Remove your Java glasses and view this code
func main() {
	var x int = 10
	if x % 2 == 0 {
		fmt.Printf("%v is even\n", x)
		fmt.Println("x is even")
	} else {
		fmt.Println("x is odd")
	}

	var count int = 0
	for count < 5 {
		fmt.Println("this is a loop")
		//count
		count++
	}
}
