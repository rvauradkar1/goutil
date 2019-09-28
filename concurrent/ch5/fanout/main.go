package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	println("Fan out........")
	ch := tasks()

	for i := 0; i < 5; i++ {
		go func(ch chan string, i int) {
			for t := range ch {
				fmt.Println("Routine ", strconv.Itoa(i), "Processing ", t)
			}
		}(ch, i)
	}
	time.Sleep(1 * time.Second)
}

func tasks() chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- "Task " + strconv.Itoa(i)
		}
		close(ch)
	}()
	return ch
}
