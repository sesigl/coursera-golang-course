package main

import (
	"fmt"
	"math"
)

func main() {

	a := requestInput("Enter a value for acceleration: ")
	v0 := requestInput("Enter a value for initial velocity: ")
	s0 := requestInput("Enter a value for initial displacement: ")

	computeDisplacement := GenDisplaceFn(a, v0, s0)

	t := requestInput("Enter a value for time: ")

	s := computeDisplacement(t)
	fmt.Printf("Calculated displacement: %v", s)
}

func requestInput(printMessage string) float64 {
	fmt.Print(printMessage)

	var f float64
	_, err := fmt.Scan(&f)
	if err != nil {
		panic(err)
	}

	return f
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return (0.5 * a * math.Pow(t, 2)) + (v0 * t) + s0
	}
}
