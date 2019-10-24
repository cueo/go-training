package main

import "fmt"

//defer is like a finally block where you can use recover to handle the errors

func catchError() {
	fmt.Println("inside catchError")
	r := recover()
	if r != nil {
		fmt.Println("Caught", r)
	}
}
func main() {
	//defer catchError()
	fmt.Println("In main")
	f1()
	fmt.Println("***************** End of main")
}

func f1() {
	defer catchError()
	f2()
}

func f2() {
	fmt.Println("inside f2")
	num  := 10
	den := 0
	div := num/den
	fmt.Println("End of F2", div)
}