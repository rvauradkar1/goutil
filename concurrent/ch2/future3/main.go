package main

import "fmt"

// Sharing by communicating
func main() {
	fmt.Println("Future chaining....")

	ch := futureTask2(futureTask1())
	fmt.Println(<-ch)
	// Alternate way of chaining them
	//ch1 := futureTask1()
	//ch2 := futureTask2(ch1)
	//fmt.Println(<-ch2)

}

func futureTask1() chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		out <- "Result 1 "
	}()
	return out
}

func futureTask2(in chan string) chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		out <- <-in + "Result 2 "
	}()
	return out
}
