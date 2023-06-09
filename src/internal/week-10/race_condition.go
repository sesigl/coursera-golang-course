package main

import "time"

func addition(i *int, j int) {

	// part of the race condition, can but not must be executed before the *2 operation
	*i = *i + j
}

func main() {

	cash := 1
	go addition(&cash, 10)

	time.Sleep(1)

	// part of the race condition, can but not must be executed before the +10 operation
	cash = cash * 2

	// hence, the value printed is not deterministic, usually it's 2 or 22 , depended on if +10 is already executed
	// I even believe there are other possibilities, e.g. that the goroutine of the addition func overrides
	// the value of the cash*2 because addition and multiplication is not atomic. though probability that this happens is a lot lower.
	print(cash)
}
