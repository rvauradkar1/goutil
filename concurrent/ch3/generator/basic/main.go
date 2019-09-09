package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Generator basic....")
	ch := generator()
	for s := range ch {
		fmt.Println(s)
	}
}

func generator() chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- strconv.Itoa(i)
		}
	}()
	return ch
}
