package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	number := rand.Intn(100)
	playGame(number)
}

func playGame(number int) {
	// println(number)
	success := false
	attempts := 0
	for !success {
		fmt.Println("Enter a number between 1 and 100:")
		var guess int
		_, err := fmt.Scanf("%d", &guess)
		// _, err := fmt.Scanln(&guess)
		if err == nil {
			if guess > 100 {
				println("Only numbers below 100 allowed!")
				continue
			}
			attempts++
			success = isSuccess(number, guess, attempts)
		}
	}
}

func isSuccess(number int, guess int, attempts int) bool {
	success := false
	if guess > number {
		println("Aim lower")
	} else if guess < number {
		println("Aim higher")
	} else {
		success = true
		fmt.Printf("You've got it in %d attempts", attempts)
	}
	return success
}
