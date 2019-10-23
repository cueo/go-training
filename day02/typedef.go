package main

import (
	"fmt"
	"strings"
)

type CoolString string

func main() {
	var cool CoolString = "hello"
	cool.funky()
}

func (str CoolString) funky() {
	var funkyString string
	for i := 0; i < len(str); i++  {
		// fmt.Println(strings.ToUpper(string(str[i])))
		if i%2 == 0 {
			funkyString += strings.ToUpper(string(str[i]))
		} else {
			funkyString += strings.ToLower(string(str[i]))
		}
	}
	fmt.Println(funkyString)
}
