package main

import (
	"fmt"
	"time"
)

func stepAside() {
	time.Sleep(1 * time.Second)
	fmt.Println("step Aside")
}

func main() {
	//launching a go routine
	go stepAside()
	fmt.Println("End of main")
	time.Sleep(5 * time.Second)
}
