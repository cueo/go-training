package main

import "fmt"

func main() {
	multilineString()
}

func multilineString() {
	para := `Hi, there!
I am mutliline string.
Enclosed inside backticks.
   I preserve spaces too!`

	fmt.Println(para)
}
