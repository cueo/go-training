package main

import (
	"fmt"
	"strings"
)

type Currency float64
type MyString string
func main() {
	var usd Currency = 71.10
	//print(usd)
	usd.print()
	var name MyString = "RAM"
	//fmt.Println(toL(name))
	fmt.Println(name.toL())
	name.greet("Hello")
}

func (mystr MyString) greet(message string) {
	fmt.Println(message + " " + string(mystr))
}


func (mystr MyString) toL() string {
	return strings.ToLower(string(mystr))
}



//func print(Currency currency) {}
func (currency Currency) print() {
	fmt.Printf("$%v\n", currency)
}
