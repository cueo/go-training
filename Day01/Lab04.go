package main

import (
	"fmt"
	"strconv"
	"strings"
)


var exclList = []int{ 10, 19, 21, 39, 309, 431, 643, 942, 1209, 7981, 8888, 9910 }

func main() {
	part1()
	part2()
	part3()
}

//input := "+5 -1 +9 +5 -67 +5 -3 +2 -4 +6 +8 -13 +2 -4 +6 +18 -3 +2 -4 +6 +88 +15 -1 +9 +5 -67 +45 -3 +2 -4 +36 +8 -13 +2 -4 +6 +18 -3 +2 -74 +11 +109"

func part3() {
	total, totalWithError := 0, 0
	input := "+5 -1 +9 +5 -67 +5 -a3 +2 -4 +6 +8 -13 +2 -4 +6 +18 -3 +2 -4 +6 +88 +15 -1 +9 +5 -67 +45 -3 +2 -4 +36 +8 -13 +2 -4 +6 +18 -3 +2 -74 +11 +109"
	arr := strings.Split(input, " ")
	for _, value := range arr {
		number, err := strconv.Atoi(strings.TrimSpace(value))
		if err == nil {
			total += number
		} else {
			fmt.Println("Error: ", err, "Number: ", number, "Actual value: ", value)
		}
		n, _ := strconv.Atoi(strings.TrimSpace(value))
		totalWithError += n

	}
	fmt.Println("Total: ", total)
	fmt.Println("TotalWithError: ", totalWithError)
}


//Calculate the 'integer' average of all numbers from 1 to 10001,
// excluding the numbers: 10, 19, 21, 39, 309, 431, 643, 942, 1209, 7981, 8888, 9910

func part2() {
	sum := 0
	for i := 1; i < 10001; i++ {
		if !presentInExclList(i) {
			sum += i
		}
	}
	fmt.Println("Average: ", int(sum/(10000 - len(exclList))))
}

func presentInExclList(number int) bool {
	found := false
	for _, value := range exclList {
		if value == number {
			found = true
			break
		}
	}
	return found
}


func presentInExclListOld(number int) bool {
	found := false
	for i := 0; i < len(exclList); i++ {
		if exclList[i] == number {
			found = true
			break
		}
	}
	return found
}
//Calculate the Sum of ODD numbers between 1 to 10001
func part1()  {
	sum := 0
	for i := 1; i < 10001; i++ {
		if i % 2 == 1 {
			sum += i
		}
	}
	fmt.Println("Sum: ", sum)
}