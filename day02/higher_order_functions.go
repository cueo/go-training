package main

import "fmt"

type Greeter func(name string) string

func main() {
	var greeter Greeter = func(name string) string {
		return "Hello, " + name
	}
	fmt.Println(greeter("momes"))
	printGreetings(greeter, "momes")

	fmt.Println(getGreeter()("momes"))
	printGreetings(getGreeter(), "momes")
}

func getGreeter() (greeter Greeter) {
	greeter = func(name string) string {
		return "Hello, " + name + "!"
	}
	return
}

// higher order function
func printGreetings(greeter Greeter, name string) {
	fmt.Println(greeter(name))
}

