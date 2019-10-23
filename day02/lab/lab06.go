package main

import (
	"fmt"
	"math"
)

var numbers = []int{101, 22, 43, 14, 5, 906, 310, 561, 84, 23, 100}

type DivFunc func([]int) []int
type SumFunc func([]int) int64

type NumCheck func(int) bool
type NumOperation func(int) int

func main() {
	fmt.Println("all:", numbers)

	sep := "\n===========\n"

	fmt.Println(sep)

	fmt.Println("even:", find(numbers, even))
	fmt.Println("odd:", find(numbers, odd))
	fmt.Println("five:", find(numbers, five))

	fmt.Println("square:", calculate(numbers, squareSum))
	fmt.Println("cube:", calculate(numbers, cubeSum))

	fmt.Println(sep)

	fmt.Println("even:", find2(numbers, func(a int) bool { return a%2 == 0 }))
	fmt.Println("odd:", find2(numbers, func(a int) bool { return a%2 != 0 }))
	f := func(a int) bool {
		return a%5 == 0
	}
	fmt.Println("five:", find2(numbers, f))

	fmt.Println("square:", calculate2(numbers, func(a int) int { return a * a }))
	fmt.Println("cube:", calculate2(numbers, func(a int) int { return a * a * a }))

	fmt.Println(sep)
}

func find2(a []int, f NumCheck) (n []int) {
	for _, v := range a {
		if f(v) {
			n = append(n, v)
		}
	}
	return
}

func calculate2(a []int, f NumOperation) int64 {
	s := 0
	for _, v := range a {
		s += f(v)
	}
	return int64(s)
}

func find(a []int, f DivFunc) []int {
	return f(a)
}

func calculate(a []int, f SumFunc) int64 {
	return f(a)
}

func powerSum(a []int, pow int) (sum int64) {
	for _, v := range a {
		sum += int64(math.Pow(float64(v), float64(pow)))
	}
	return
}

func squareSum(a []int) int64 {
	return powerSum(a, 2)
}

func cubeSum(a []int) int64 {
	return powerSum(a, 3)
}

func divisible(a []int, d int, isDiv bool) (e []int) {
	for _, v := range a {
		if (v%d == 0) == isDiv {
			e = append(e, v)
		}
	}
	return
}

func even(a []int) []int {
	return divisible(a, 2, true)
}

func odd(a []int) []int {
	return divisible(a, 2, false)
}

func five(a []int) []int {
	return divisible(a, 5, true)
}
