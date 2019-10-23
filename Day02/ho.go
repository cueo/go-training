package main

import "fmt"

type Greeter func() string

//Higher order function
func doSomething(greeter Greeter) {
	fmt.Println(greeter())
}

//Higher order function
func doSomethingElse(fn func() string) {

}

//func playWithSomething() func() string {}
func playWithSomething() Greeter {
	//closures
	welcome := "Namaskara"
	var localGreeting Greeter = func() string {
		return welcome
	}
	return localGreeting
}

func main() {
	greetingsInKannada := playWithSomething()
	fmt.Println(greetingsInKannada())

	var hi Greeter = func() string {
		return "Hi"
	}
	doSomething(hi)

	var hello Greeter = func() string {
		return "Hello"
	}
	hello()




	var greet = func() {
		fmt.Println("How are you?")
	}
	greet()

	var print = func(name string)  {
		fmt.Println(name)
	}
	print("Sam")
}
