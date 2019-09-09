package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Generator basic....")
	done := make(chan bool)
	ch := generator(done)
	for s := range ch {
		fmt.Println(s)
		if s == "1000" {
			fmt.Println("Sending done")
			done <- true
		}
	}
}

func generator(done chan bool) chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 0; ; i++ {
			select {
			case <-done:
				fmt.Println("I am wrapping up.......")
				return
			case ch <- strconv.Itoa(i):
			}
		}
	}()
	return ch
}
