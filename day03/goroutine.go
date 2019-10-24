package main

import (
	"fmt"
	"time"
)

func main() {
	// nogo()
	gogo()
}

func gogo() {
	go goroutine()
	fmt.Println("End of course")
	time.Sleep(3 * time.Second)
}

func nogo() {
	goroutine()
	fmt.Println("End of course")
	time.Sleep(3 * time.Second)
}

func goroutine() {
	time.Sleep(1 * time.Second)
	fmt.Println("I am executing in a different thread")
}
