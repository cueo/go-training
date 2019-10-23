package main

import (
	"errors"
	"fmt"
)

func main() {
	// fmt.Println(divide(12, 0))  // panic: runtime error
	fmt.Println(divideHandle(12, 0))
	a, err := divideWithError(12, 0)
	a, err = divideWithError(12, 0)
	if err == nil {
		fmt.Println(a)
	} else {
		fmt.Println(err.Error())
	}
}

func divide(n int, d int) int {
	return n / d
}

func divideHandle(n int, d int) (a int) {
	if d == 0 {
		fmt.Println("Denominator cannot be zero!")
	} else {
		a = n / d
	}
	return
}

func divideWithError(n int, d int) (a int, err error) {
	if d == 0 {
		err = errors.New("denominator cannot be zero")
	} else {
		a, err = n / d, nil
	}
	return
}
