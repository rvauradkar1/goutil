package main

import "fmt"

func main() {
	fmt.Println("Streaming futures....")

	ch := streamingFuture()
	for i := range ch {
		fmt.Println(i)
	}
}

func streamingFuture() chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()
	return ch
}
