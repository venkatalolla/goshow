// A guess challenge game to guess a random number
package main

import (
	"github.com/venkatalolla/goshow/guessgame"
)

func main() {
	// Random number generator
	randomNumber := guessgame.RandomNumberGenerator()

	// Print welcome messages
	guessgame.WelcomeStatements()

	// Call user input
	input := guessgame.UserInputFromPrompt()

	// Send the user input to game
	guessgame.Game(input, randomNumber)
}
