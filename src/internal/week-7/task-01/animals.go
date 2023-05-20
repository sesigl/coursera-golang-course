package main

import (
	"fmt"
)

type animal struct {
	food       string
	locomotion string
	noise      string
}

var cow = animal{
	food:       "grass",
	locomotion: "walk",
	noise:      "moo",
}
var bird = animal{
	food:       "worms",
	locomotion: "fly",
	noise:      "peep",
}
var snake = animal{
	food:       "mice",
	locomotion: "slither",
	noise:      "hsss",
}

func (a animal) Eat() {
	fmt.Println(a.food)
}

func (a animal) Move() {
	fmt.Println(a.locomotion)
}

func (a animal) Speak() {
	fmt.Println(a.noise)
}

func (a animal) Print(requestedInformation string) bool {
	switch requestedInformation {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	default:
		return false
	}

	return true
}

var animals = map[string]animal{
	"cow":   cow,
	"bird":  bird,
	"snake": snake,
}

func main() {

	for true {
		fmt.Print("How may I help?\n> ")
		requestedAnimalName, requestedInformation, scanCount := requestUserInput()
		if scanCount != 2 {
			fmt.Println("invalid input, syntax: <animal-name> <animal-information-request>")
			continue
		}

		foundAnimal, ok := animals[requestedAnimalName]
		if !ok {
			fmt.Printf("invalid input\nno animal with name '%s' found, pick one of cow|bird|snake", requestedAnimalName)
			continue
		}

		ok = foundAnimal.Print(requestedInformation)
		if !ok {
			fmt.Printf("invalid information request for '%s', pick one of eat|move|speak.", requestedInformation)
			continue
		}
	}

}

func requestUserInput() (string, string, int) {
	var requestedAnimalName, requestedInformation string
	scanCount, err := fmt.Scan(&requestedAnimalName, &requestedInformation)
	if err != nil {
		panic(err)
	}
	return requestedAnimalName, requestedInformation, scanCount
}
