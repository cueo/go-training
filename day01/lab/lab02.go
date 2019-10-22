package main

import (
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {
	input := "madamz"
	// var input string
	// fmt.Scanf("%s", &input)
	isPalindrome := checkPalindrome(input)
	if isPalindrome {
		println("String is palindrome!")
	} else {
		println("String is NOT palindrome!")
	}
	anagram := createAnagram2(input)
	println("Anagram:", anagram)
	encodedInput := encode(input)
	println("Encoded:", encodedInput)
}

func checkPalindrome(input string) bool {
	start, end := 0, len(input)-1
	for start < end {
		if input[start] != input[end] && math.Abs(float64(input[start]-input[end])) != 32 {
			return false
		}
		start++
		end--
	}
	return true
}

func encode(input string) string {
	for i := 0; i < len(input); i++ {
		input = input[:i] + string(input[i]+1) + input[i+1:]
	}
	return input
}

func createAnagram(input string) string {
	rand.Seed(time.Now().UTC().UnixNano())
	for count := 1; count < len(input); count++ {
		r := rand.Intn(len(input))
		rChar, countChar := string(input[r]), string(input[count])
		input = input[:count] + rChar + input[count+1:]
		input = input[:r] + countChar + input[r+1:]
	}
	return input
}

func createAnagram2(input string) string {
	rand.Seed(time.Now().UTC().UnixNano())
	var anagram string
	for count := 0; count < len(input); {
		r := rand.Intn(len(input))
		if input[r] != '_' {
			letter := string(input[r])
			anagram += letter
			strings.Replace(input, letter, "_", 1)
			count++
		}
	}
	return anagram
}
