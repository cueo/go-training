package main

import "fmt"

func main() {
	fmt.Println("Started main")
	divide()
	fmt.Println("Completed main")
}

func divide()  {
	defer catch("div")
	defer func() {fmt.Println("Second defer is called first! LIFO.")}()
	n, d := 69, 0
	div := n / d
	fmt.Println("Divided and conquered", div)
}

func catch(identifier string) {
	if r := recover(); r != nil {
		fmt.Printf("Error occurred in %s! Error=%s :(\n", identifier, r)
	}
}
