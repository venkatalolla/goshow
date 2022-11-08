package guessgame

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func UserInputFromPrompt() int {
	// Read user input
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	// Error handling the user input
	if err != nil {
		log.Fatal(err)
	}

	// Convert the string to int
	input = strings.TrimSpace(input)
	userInput, err := strconv.Atoi(input)

	// Error handling if the number is a string
	if err != nil {
		log.Fatal("No words or characters please! Just pass numbers")
	}

	// Error handling if the number is negative or higher than 100
	if userInput < 0 {
		log.Fatal(customErrMsg("%d is a negative numbers, use a positive number", userInput))
	} else if userInput > 100 {
		log.Fatal(customErrMsg("%d is higher than 100, choose a number between 1 and 100", userInput))
	}

	return userInput
}

func Game(userInput int, randomNumber int) {
	//  Loop the user to try 10 times
	success := false

	for minTries := 1; minTries <= 10; minTries++ {
		// Check for the low/high numbers
		if userInput < randomNumber {
			oopsMessage("low")
			retriesLeft(minTries)
			userInput = UserInputFromPrompt()
		} else if userInput > randomNumber {
			oopsMessage("high")
			retriesLeft(minTries)
			userInput = UserInputFromPrompt()
		} else {
			success = true
			fmt.Println("Hurrah! you guessed the correct number in", minTries, "tries")
			break
		}
	}
	// If ran out of tries
	ranOutOfTries(success, randomNumber)
}
