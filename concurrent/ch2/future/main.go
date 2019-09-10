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
		ch <- "Result..."
	}()
	return ch
}
