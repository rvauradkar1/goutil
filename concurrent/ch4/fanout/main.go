package main

import (
	"strconv"
	"sync"
)

func main() {
	println("Fan in........")

	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 0; i < 3; i++ {
			ch <- "First " + strconv.Itoa(i)
		}
	}()

	ch1 := make(chan string)
	go func() {
		defer close(ch1)
		for i := 0; i < 3; i++ {
			ch1 <- "Second " + strconv.Itoa(i)
		}
	}()

	ch2 := make(chan string)
	go func() {
		defer close(ch2)
		for i := 0; i < 3; i++ {
			ch2 <- "Third " + strconv.Itoa(i)
		}
	}()

	out := merge(ch, ch1, ch2)

	for s := range out {
		println(s)
	}

}

func merge(channels ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	wg.Add(len(channels))
	out := make(chan string)

	for _, ch := range channels {
		go func(c <-chan string) {
			for s := range c {
				out <- s
			}
			wg.Done()
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
