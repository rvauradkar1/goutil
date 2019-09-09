package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Fan in....")
	tasks := taskGenerator()
	for s := range tasks {
		fmt.Println("val = " + s)
	}
}

func taskGenerator() chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 0; i < 3; i++ {
			ch <- strconv.Itoa(i)
		}
	}()
	return ch
}
