// a simple guessing program

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Guess a number between 1 t0 15, you have five (5) trials to guess correctly")

	// generate a random number
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(5) // generate random number between 0 and 10


	var guess int
	var MaxAttempt = 5
	var count int


	for {
	if count < MaxAttempt {
		fmt.Println("Please input your guess")
		fmt.Scan(&guess)
		if guess > secretNumber {
			fmt.Println("Too big, please try again")
		} else if guess < secretNumber {
			fmt.Println("Too Small, please try again")
		} else {
			fmt.Println("Congratulations! You won.")
			fmt.Printf("you guessed %d times", count)
			break
		}
	} else {
		fmt.Println("Sorry you have exceeded your maximum amount of attempts")
		break
	}
		count++
	}

}