package main

import "fmt"

var x int = 9
var name string = "momes"

func main()  {
	fmt.Println(x)
	// println is a builtin method
	println("name is " + name)
	local()
}

func local() {
	var city string
	println("city is" + city)

	var i int = 9
	fmt.Println(i)

	// !! gives an error if not used !!
	//var isCool bool = true
}
