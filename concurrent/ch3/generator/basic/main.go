package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Generator basic....")
	ch := generate()
	for s := range ch {
		fmt.Println(s)
	}
}

func generate() chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- strconv.Itoa(i)
		}
	}()
	return ch
}
