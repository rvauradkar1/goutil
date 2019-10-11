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
	}
}

func generator(done chan bool) chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; ; i++ {
			if i == 3 {
				// Producer want to stop streaming after 3 items
				return
			}
			ch <- i
		}
	}()
	return ch
}
