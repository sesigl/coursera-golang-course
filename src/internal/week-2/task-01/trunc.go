package main

import (
	"fmt"
)

func main() {
	fmt.Println("Please enter a floating point number (e.g. 4.937):")

	var floatFromUser float64
	_, err := fmt.Scanln(&floatFromUser)
	if err != nil {
		panic(err)
	}
	//math.Trunc
	numberCastedToInt := int64(floatFromUser)

	fmt.Printf("After receiving %v, the integer version of it is: %v", floatFromUser, numberCastedToInt)
}
