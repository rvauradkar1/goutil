package main

import "fmt"

// Sharing by communicating
func main() {
	fmt.Println("Futures....")

	ch := futureTask()

	fmt.Println(<-ch)
}

func futureTask() chan string {
	ch := make(chan string)
	go func() {
		// Make sure to close channel after using it
		defer close(ch)
		ch <- "Result..."
	}()
	return ch
}
