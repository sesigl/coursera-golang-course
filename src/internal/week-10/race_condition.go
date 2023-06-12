package main

import (
	"fmt"
	"sync"
)

//Taking two integers and modify them by nullifying, multiplying by 2 or incrementing by 2.
//Due to race condition the variables a and b are being modified by whichever thread is currently being scheduled to execute.
//Output result is unpredictable.

// 3 passes are given to introduce the race condition behavior
func main() {

	numberOfPasses := 3
	counter := 0

	// 3 trials to run
	for counter < numberOfPasses {

		//Variables being assigned
		a := 2
		b := 2

		fmt.Println("++++++++++++++++++")
		fmt.Println("Starting pass", counter)

		var wg sync.WaitGroup
		wg.Add(3)
		//Executing goroutines concurently

		go incrementBy2(&a, &b, &wg)
		go multiplyBy2(&a, &b, &wg)
		go nullifyIntegers(&a, &b, &wg)

		wg.Wait()

		fmt.Println("end of main number:", counter)
		fmt.Println("++++++++++++++++++")
		counter++

	}

}

func incrementBy2(a, b *int, wg *sync.WaitGroup) {
	*a = *a + 2
	*b = *b + 2
	fmt.Println("incremented", *a, *b)
	wg.Done()

}
func multiplyBy2(a, b *int, wg *sync.WaitGroup) {
	*a = *a * 2
	*b = *b * 2
	fmt.Println("multiplied", *a, *b)
	wg.Done()

}

func nullifyIntegers(a, b *int, wg *sync.WaitGroup) {
	*a = 0
	*b = 0
	fmt.Println("nullified", *a, *b)
	wg.Done()

}
