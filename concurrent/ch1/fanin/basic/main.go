package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Fan in....")
	ch := s1()
	for s := range ch {
		fmt.Println("val = " + s)
	}
}

func s1() chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 0; i < 3; i++ {
			ch <- strconv.Itoa(i)
		}
	}()
	return ch
}
