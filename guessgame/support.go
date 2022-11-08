package guessgame

import (
	"fmt"
	"math/rand"
	"time"
)

func retriesLeft(minTries int) {
	fmt.Println("You have", 10-minTries, "left")
}

func oopsMessage(valueType string) {
	fmt.Println("Oops! Your guess was", valueType, "give it another try")
}

func ranOutOfTries(success bool, randomNumber int) {
	// Ran out of tries message
	if !success {
		fmt.Println("Oops! You ran out of tries, better luck next time!!! But here is my target", randomNumber)
	}
}

func WelcomeStatements() {
	// Welcome Statements
	fmt.Println("Welcome to the Guess game! \nI choose a number between 1 and 100 \nCan you guess it?")
	fmt.Println("Enter your number:")
}

func RandomNumberGenerator() int {
	// Set the seed and random number
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100) + 1
}

func customErrMsg(errorMsg string, errValue int) error {
	err := fmt.Errorf(errorMsg, errValue)
	return err
}
