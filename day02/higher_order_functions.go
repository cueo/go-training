package main

import "fmt"

type Greeter func(name string) string

func main() {
	var greeter Greeter = func(name string) string {
		return "Hello, " + name
	}
	fmt.Println(1, greeter("momes"))
	printGreetings(greeter, "momes")

	greet := getGreeter()
	fmt.Println(3, greet("momes"))
	fmt.Println(4, getGreeter()("momes"))
	printGreetings(getGreeter(), "momes")
}

func getGreeter() (greeter Greeter) {
	message := "Hello, "
	greeter = func(name string) string {
		message += name + "!"
		return message
	}
	return
}

// higher order function
func printGreetings(greeter Greeter, name string) {
	fmt.Println("in", greeter(name))
}

