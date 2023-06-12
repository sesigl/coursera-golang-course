package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func sortInts(ints []int, wg *sync.WaitGroup) {
	defer wg.Done()
	sort.Ints(ints)
}

func main() {

	ints := getIntsFromUserInput()

	sortGroupCount := int(math.Ceil(float64(len(ints)) / 4))

	sortedGroups := concurrentSort(sortGroupCount, ints)
	sortedInts := appendIntGroups(ints, sortGroupCount, sortedGroups)
	sortFinalList(sortedInts)

}

func getIntsFromUserInput() []int {
	fmt.Println("Provide a list of integers to sort comma-separated")

	var commaSeparatedInts string
	fmt.Scanln(&commaSeparatedInts)

	intsAsStr := strings.Split(commaSeparatedInts, ",")

	ints := make([]int, 0, len(commaSeparatedInts))

	for _, val := range intsAsStr {
		parsedInt, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}

		ints = append(ints, parsedInt)
	}
	return ints
}

func concurrentSort(sortGroupCount int, ints []int) [][]int {
	sortGroups := make([][]int, 0, sortGroupCount)
	var wg sync.WaitGroup

	for i := 0; i < sortGroupCount; i++ {
		wg.Add(1)

		sortGroupSliceStart := i * 4
		sortGroupSliceEnd := i*4 + 4

		if sortGroupSliceEnd > len(ints) {
			sortGroupSliceEnd = len(ints)
		}
		sortGroups = append(sortGroups, ints[sortGroupSliceStart:sortGroupSliceEnd])
		fmt.Printf("Group %v before sorting: %v\n", i, sortGroups[i])
		go sortInts(sortGroups[i], &wg)
	}

	wg.Wait()
	return sortGroups
}

func appendIntGroups(ints []int, sortGroupCount int, sortedGroups [][]int) []int {
	sortedInts := make([]int, 0, len(ints))

	for i := 0; i < sortGroupCount; i++ {
		fmt.Printf("Group %v after sorting sorting: %v\n", i, sortedGroups[i])
		sortedInts = append(sortedInts, sortedGroups[i]...)
	}
	return sortedInts
}

func sortFinalList(sortedInts []int) {
	fmt.Printf("merged list before final sort %v: ", sortedInts)
	sort.Ints(sortedInts)
	fmt.Printf("final list sorted list %v: ", sortedInts)
}
