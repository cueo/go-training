package main

import "fmt"

func main() {
	recoverExample()
	println("====")
	panicExample()
}

func panicExample() {
	fmt.Println("Start")
	n, d := 12, 0
	if d == 0 {
		panic("NOOOOOOOO!!! Why you try divide by zero? :(")
	} else {
		div := n / d
		fmt.Println(div)
	}
	fmt.Println("Done")  // does not print if panic
}


func recoverExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from error:", r)
		}
	}()
	panicExample()
	fmt.Println("Done")  // does not print if panic
}
