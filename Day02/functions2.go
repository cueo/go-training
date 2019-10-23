package main

import (
	"errors"
	"fmt"
)

func main() {
	//fmt.Println(divide(12, 0))
	div, err := divide(12, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(div)
	}

	//var str string
	//fmt.Println("str", str, len(str))
}

func divide(num int, den int) (div int, err error) {
	if den == 0 {
		err = errors.New("Den cannot be zero")
		return num, err
	} else {
		div = num / den
	}
	return div, nil
}