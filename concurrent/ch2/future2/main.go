package main

import "fmt"

// Sharing by communicating
func main() {
	fmt.Println("Future chaining....")

	//ch1 := futureTask1()
	//ch2 := futureTask2(ch1)

	ch1 := futureTask2(futureTask1())

	fmt.Println(<-ch1)
}

func futureTask1() chan string {
	out := make(chan string)
	go func() {
		out <- "Result 1 "
	}()
	return out
}

func futureTask2(in chan string) chan string {
	out := make(chan string)
	go func() {
		out <- <-in + "Result 2 "
	}()
	return out
}
