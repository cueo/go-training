package main

import (
	"strconv"
	"strings"
)

func main() {
	sum := part1()
	println("part 1:", sum)

	arr := []int {10, 19, 21, 39, 309, 431, 643, 942, 1209, 7981, 8888, 9910}
	m := initializeMap(arr)
	// sum = part2a(arr)
	sum = part2b(m)
	println("part 2:", sum)

	input := "+5 -1 +9 +5 -67 +5 -3 +2 -4 +6 +8 -13 +2 -4 +6 +18 -3 +2 -4 +6 +88 +15 -1 +9 +5 -67 +45 -3 +2 -4 +36 +8 -13 +2 -4 +6 +18 -3 +2 -74 +11 +109"
	sum3 := part3(input)
	println("part 3:", sum3)
}

func initializeMap(arr []int) map[int]bool {
	m := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		m[arr[i]] = true
	}
	return m
}

func part1() int {
	sum := 0
	for i := 1; i <= 10001; i++ {
		if i%2 != 0 {
			sum += i
		}
	}
	return sum
}

func part2a(arr []int) int {
	sum := 0
	for i := 1; i <= 10001; i++ {
		if !inArray(i, arr) {
			sum += i
		}
	}
	avg := sum / (10001 - len(arr))
	return avg
}

func part2b(m map[int]bool) int {
	sum := 0
	for i := 1; i <= 10001; i++ {
		// println(i, m[i])
		if !m[i] {
			sum += i
		}
	}
	avg := sum / (10001 - len(m))
	return avg
}

func inArray(num int, arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			return true
		}
	}
	return false
}

func part3(input string) int64 {
	stringArray := strings.Split(input, " ")
	var sum int64
	for i := 0; i < len(stringArray); i++ {
		// ignore error
		num, _ := strconv.ParseInt(stringArray[i], 10, 64)
		sum += num
	}
	return sum
}
