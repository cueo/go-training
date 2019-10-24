package main

import "fmt"

func main() {
	deferred()
	defer deferred()
	fmt.Println("I am main!")
}

func deferred() {
	fmt.Println("I am deferred!")
}
