package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var INR string



//const PI float64 = 3.14

func main() {
	word := "helloz"
	palindrome(word)
	anagram(word)
	encode(word)
}

func encode(word string) {
	encodedWord := ""
	for i := 0; i < len(word); i++ {
		letter := string(word[i] + 1)
		if string(word[i]) == "z" {
			letter = "a"
		} else if string(word[i]) == "Z" {
			letter = "A"
		}
		encodedWord += letter
	}
	fmt.Println(encodedWord)
}

func anagram(word string) {
	jumbledWord := ""

	for count := 0; count < len(word); {
		rand.Seed(time.Now().UTC().UnixNano())
		randomIndex := rand.Intn(len(word))
		letter := string(word[randomIndex])
		if letter != "_" {
			jumbledWord += letter
			word = strings.Replace(word, letter, "_", 1)
			count++
		}
	}
	fmt.Println(jumbledWord)
}

func palindrome(word string) {
	reversedWord := ""
	for i := len(word) - 1; i >= 0 ; i--  {
		reversedWord += string(word[i])
	}
	//fmt.Println(word, word == reversedWord)
	fmt.Println(word, strings.EqualFold(word, reversedWord)) //equalsIgnoreCase in Java
}