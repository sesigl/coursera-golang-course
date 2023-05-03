package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	inputData := make(map[string]string)

	fmt.Println("Enter your name:")

	var name string
	mustScan(&name)
	inputData["name"] = name

	fmt.Println("Enter your Address:")

	var address string
	mustScan(&address)
	inputData["address"] = address

	inputDataJson, err := json.Marshal(&inputData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("JSON envoded input data: %s\n", inputDataJson)
}

func mustScan(a ...any) (n int) {
	n, err := fmt.Scan(a...)
	if err != nil {
		panic(err)
	}
	return n
}
