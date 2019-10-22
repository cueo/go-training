package main

import "fmt"

var x int = 10
var name string = "WTH is going on?"

func local() {
	var city string
	fmt.Println("city is " + city)

	var i int = 20
	fmt.Println(i)

	//Variables have to be used
	var isCool bool = true
	fmt.Println(isCool)


}

func main() {
	fmt.Println(x)
	fmt.Println(name)
	local()
}
