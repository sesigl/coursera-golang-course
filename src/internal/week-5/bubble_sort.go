package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	numbers := RequestNumbersFromUser()
	BubbleSort(numbers)

	fmt.Println(numbers)
}

func RequestNumbersFromUser() []int {
	var userInput string
	fmt.Print("Type up to 10 integers separated by comma without whitespaces:")

	_, err := fmt.Scan(&userInput)
	if err != nil {
		panic(err)
	}

	numbers := make([]int, 0, 10)

	splits := strings.Split(userInput, ",")

	if len(splits) > 10 {
		panic("Don't provide more than 10 integers")
	}

	for _, split := range splits {
		parsedInt, err := strconv.Atoi(split)
		if err != nil {
			panic(fmt.Sprintf("Only provide numbers, you povided: %s", split))
		}
		numbers = append(numbers, parsedInt)
	}
	return numbers
}

func BubbleSort(numbers []int) {
	firstRun := true
	sorted := true

	for sorted == false || firstRun == true {

		sorted = true
		firstRun = false

		for i := 0; i < len(numbers)-1; i++ {

			if numbers[i+1] < numbers[i] {
				Swap(numbers, i)
				sorted = false
			}

		}
	}
}

func Swap(numbers []int, i int) {
	tmp := numbers[i]
	numbers[i] = numbers[i+1]
	numbers[i+1] = tmp
}
