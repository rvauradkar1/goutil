package main

import (
	"fmt"
)

func main() {
	fmt.Println("Generator advanced....")
	done := make(chan bool)
	ch := generator(done)
	for s := range ch {
		fmt.Println(s)
		// Consumer wants to stop streaming at the fifth item
		if s == 5 {
			fmt.Println("Sending done")
			done <- true
		}
	}
}

func generator(done chan bool) chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; ; i++ {
			select {
			// If consumer is done, then stop streaming
			case <-done:
				fmt.Println("I am wrapping up.......")
				return
			case ch <- i:
			}
		}
	}()
	return ch
}
