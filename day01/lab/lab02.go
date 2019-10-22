package main

import (
	rand2 "math/rand"
	"time"
)

func main() {
	var input string = "madam"
	// fmt.Scanf("%s", &input)
	isPalindrome := checkPalindrome(input)
	if isPalindrome {
		println("String is palindrome!")
	} else {
		println("String is NOT palindrome!")
	}
	anagram := createAnagram(input)
	println("Anagram:", anagram)
	encodedInput := encode(input)
	println("Encoded:", encodedInput)
}

func checkPalindrome(input string) bool {
	start, end := 0, len(input) - 1
	for start < end {
		if input[start] != input[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func encode(input string) string {
	for i := 0; i < len(input); i++ {
		input = input[:i] + string(input[i] + 1) + input[i+1:]
	}
	return input
}

func createAnagram(input string) string {
	rand2.Seed(time.Now().UTC().UnixNano())
	for count := 1; count < len(input); count++ {
		r := rand2.Intn(len(input))
		rChar, countChar := string(input[r]), string(input[count])
		input = input[:count] + rChar + input[count+1:]
		input = input[:r] + countChar + input[r+1:]
	}
	return input
}
