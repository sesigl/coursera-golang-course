package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type person struct {
	fname string
	lname string
}

func (p person) String() string {
	return fmt.Sprintf(`
		first name:			%s
		last name:			%s`, p.fname, p.lname)
}

func main() {
	filePath := requestFilePath()

	file := mustOpen(filePath)
	defer file.Close()

	persons := scanPersonsFromFile(file)

	printSummary(persons)
}

func requestFilePath() string {
	fmt.Print("enter a path to a file to parse:")

	var filePath string
	_, err := fmt.Scan(&filePath)
	if err != nil {
		panic(err)
	}

	return filePath
}

func mustOpen(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	return file
}

func scanPersonsFromFile(file *os.File) []person {
	scanner := bufio.NewScanner(file)

	persons := []person{}
	for scanner.Scan() {
		lineText := scanner.Text()
		lineSplit := strings.Split(lineText, " ")
		if len(lineSplit) != 2 {
			panic(fmt.Sprintf("Can not parse line '%v'", lineText))
		}

		persons = append(persons, person{
			fname: lineSplit[0],
			lname: lineSplit[1],
		})
	}

	return persons
}

func printSummary(persons []person) {
	fmt.Printf("Found %v persons:", len(persons))
	for _, person := range persons {
		fmt.Println(person)
	}
}
