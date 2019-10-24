package main

import "fmt"

func main() {
	fmt.Println("Start")
	n, d := 12, 0
	if d == 0 {
		panic("NOOOOOOOO!!!")
	} else {
		div := n / d
		fmt.Println(div)
	}
	fmt.Println("Done")
}
