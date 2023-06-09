package main

import (
	"fmt"
	"reflect"
)

type Animal interface {
	Eat()
	Move()
	Speak()

	Print(requestedInformation string) bool
}

type animalData struct {
	food       string
	locomotion string
	noise      string
}

var cow = animalData{
	food:       "grass",
	locomotion: "walk",
	noise:      "moo",
}
var bird = animalData{
	food:       "worms",
	locomotion: "fly",
	noise:      "peep",
}
var snake = animalData{
	food:       "mice",
	locomotion: "slither",
	noise:      "hsss",
}

type Cow struct {
	animalData
}

type Bird struct {
	animalData
}

type Snake struct {
	animalData
}

func (a animalData) Eat() {
	fmt.Println(a.food)
}

func (a animalData) Move() {
	fmt.Println(a.locomotion)
}

func (a animalData) Speak() {
	fmt.Println(a.noise)
}

func (a animalData) Print(requestedInformation string) bool {
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

var animals = map[string]Animal{}

func main() {

	for true {
		fmt.Print("How may I help?\n> ")
		requestedCommand, requestedAnimalName, requestedInformationOrTypeOfAnimal, scanCount := requestUserInput()
		if scanCount != 3 {
			fmt.Println("invalid input, \n" +
				"syntax: newanimal <animal-name> <type-of-animal>\n" +
				"syntax: query <animalData-name> <information-request>")
			continue
		}

		if requestedCommand == "query" {
			foundAnimal, ok := animals[requestedAnimalName]
			if !ok {
				fmt.Printf("invalid input\nno animal with name '%s' found, create one before querying\n", requestedAnimalName)
				continue
			}

			ok = foundAnimal.Print(requestedInformationOrTypeOfAnimal)
			if !ok {
				fmt.Printf("invalid information request for '%s', pick one of eat|move|speak\n", requestedInformationOrTypeOfAnimal)
				continue
			}
		} else if requestedCommand == "newanimal" {

			var createdAnimal Animal
			switch requestedInformationOrTypeOfAnimal {
			case "cow":
				createdAnimal = Cow{cow}
			case "bird":
				createdAnimal = Bird{bird}
			case "snake":
				createdAnimal = Snake{snake}
			default:
				{
					fmt.Printf("invalid animal '%s', pick one of cow|bird|snake\n", requestedInformationOrTypeOfAnimal)
					continue
				}
			}

			animals[requestedAnimalName] = createdAnimal

			fmt.Printf("Created animal '%s': %v\n", requestedAnimalName, reflect.TypeOf(createdAnimal).Name())

		} else {
			fmt.Printf("invalid command '%s', pick one of query|newanimal\n", requestedCommand)
			continue
		}

	}

}

func requestUserInput() (string, string, string, int) {
	var command, requestedAnimalName, requestedInformation string
	scanCount, err := fmt.Scan(&command, &requestedAnimalName, &requestedInformation)
	if err != nil {
		panic(err)
	}
	return command, requestedAnimalName, requestedInformation, scanCount
}
