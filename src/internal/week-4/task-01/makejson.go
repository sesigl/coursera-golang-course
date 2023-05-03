package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	inputData := make(map[string]string)

	fmt.Print("Enter your name: ")
	var name string
	mustScan(&name)
	inputData["name"] = name

	fmt.Print("Enter your Address: ")
	var address string
	mustScan(&address)
	inputData["address"] = address

	inputDataJson := mustMarshal(inputData)

	fmt.Printf("JSON encoded input data: %s\n", inputDataJson)
}

func mustScan(a ...any) (n int) {
	n, err := fmt.Scan(a...)
	if err != nil {
		panic(err)
	}
	return n
}

func mustMarshal(v any) []byte {
	inputDataJson, err := json.Marshal(&v)
	if err != nil {
		panic(err)
	}
	return inputDataJson
}
