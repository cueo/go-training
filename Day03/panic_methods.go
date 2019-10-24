package main

import "fmt"

func catch(functionName string) {
	r := recover()
	if r != nil {
		fmt.Println("Something has happened in ", functionName, r)
	} else {
		fmt.Println("everything is fine. happy diwali")
	}
}

func main() {
	fmt.Println("Inside main")
	divide()
	fmt.Println("End of main")
}

func divide() {
	//LIFO
	defer catch("divide")
	defer func() { fmt.Println("second defer") }()

	num := 10
	den := 0
	div := num/den
	fmt.Println("End of divide", div)
}












func main2() {
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
