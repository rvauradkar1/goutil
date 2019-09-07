package main

import "fmt"

// Sharing by communicating
func main() {
	fmt.Println("Channels....")

	ch := service()

	fmt.Println(<-ch)
}

func service() chan string {
	ch := make(chan string)
	go func() {
		ch <- "Result..."
	}()
	return ch
}
