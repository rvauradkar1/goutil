package main

import (
	"strconv"
	"sync"
)

func main() {
	println("Fan in........")

	ch := makech("First ")
	ch1 := makech("Second ")
	ch2 := makech("Third ")

	out := merge(ch, ch1, ch2)

	for s := range out {
		println(s)
	}

}

func makech(s string) chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 0; i < 3; i++ {
			ch <- s + strconv.Itoa(i)
		}
	}()
	return ch
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
