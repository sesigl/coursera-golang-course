package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var fileName string
	//var err error

	//fileName = "/home/juaneco/jcfr/desarrollo/go/input.txt"

	fmt.Printf("\nEnter a file name: ")
	_, _ = fmt.Scan(&fileName)

	file, _ := os.Open(fileName)

	scanner := bufio.NewScanner(file)

	var persons = make([]PersonStruct, 0)

	for scanner.Scan() {

		line := scanner.Text()
		lineSplit := strings.Split(line, " ")

		// creating a new person
		person := PersonStruct{
			fname: lineSplit[0],
			lname: lineSplit[1],
		}

		// append person to the slice persons
		persons = append(persons, person)
	}

	_ = file.Close()

	for _, person := range persons {
		fmt.Printf("%s %s\n", person.fname, person.lname)
	}
}

type PersonStruct struct {
	fname string
	lname string
}
