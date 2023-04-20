package main

import (
	"fmt"
	"strings"
)

const requiredFirstLetter = 'i'
const requiredMiddleLetter = 'a'
const requiredEndLetter = 'n'

func main() {
	userInput := readUserInput()

	fmt.Println("Enter a string")

	if len(userInput) < 3 {
		panic("to fulfil the task input needs to have at least 3 characters")
	}

	userInputLower := strings.ToLower(userInput)

	if containsRequiredLetters(userInputLower) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}

func containsRequiredLetters(userInputLower string) bool {
	return userInputLower[0] == requiredFirstLetter && userInputLower[len(userInputLower)-1] == requiredEndLetter && strings.ContainsRune(userInputLower, requiredMiddleLetter)
}

func readUserInput() string {
	var userInput string

	_, err := fmt.Scanln(&userInput)
	if err != nil {
		panic(err)
	}

	return userInput
}
