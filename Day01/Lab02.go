package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Enter a number between 1 and 100")
	rand.Seed(time.Now().UTC().UnixNano())
	target, attempts, guess := rand.Intn(100), 0, -1
	var message string
	over := false
	fmt.Println(target)
	for !over {
		fmt.Scanf("%d", &guess)
		attempts++
		if guess > target {
			message = "Aim Lower"
		} else if guess < target {
			message = "Aim Higher"
		} else {
			message = "You've got it in " + strconv.Itoa(attempts) + " attempts"
			over = true
		}
		fmt.Println(message)
	}
}
