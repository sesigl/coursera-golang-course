package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sortedNumbers := make([]int, 3)

	var userInput string
	for userInput != "X" {
		fmt.Print("Enter \"X\" to exit or enter a number:")
		userInput = requestUserInput()

		parseIntAndExecuteIfSuccessful(userInput, func(parsedInt int) {
			sortedNumbers = sortNumbers(sortedNumbers, parsedInt)
			fmt.Printf("Current sorted array: %v\n", sortedNumbers)
		})
	}

}

func requestUserInput() string {
	return mustScanLn()
}

func mustScanLn() string {
	var userInput string
	_, err := fmt.Scanln(&userInput)
	if err != nil {
		panic(err)
	}

	return userInput
}

func parseIntAndExecuteIfSuccessful(userInput string, fn func(parsedInt int)) {
	parsedUserInput, err := parseInt(userInput)
	if err != nil {
		fmt.Println("Invalid number, ... try again")
	} else {
		fn(parsedUserInput)
	}
}

func sortNumbers(sortedNumbers []int, parsedUserInput int) []int {
	sortedNumbers = appendAndSort(sortedNumbers, parsedUserInput)
	return sortedNumbers
}

func parseInt(userInput string) (int, error) {
	parsedUserInput, err := strconv.Atoi(userInput)
	return parsedUserInput, err
}

func appendAndSort(sortedNumbers []int, parsedUserInput int) []int {
	sortedNumbers = append(sortedNumbers, parsedUserInput)
	sort.Ints(sortedNumbers)
	return sortedNumbers
}
