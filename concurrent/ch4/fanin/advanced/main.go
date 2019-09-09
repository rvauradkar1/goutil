package main

import (
	"strconv"
	"sync"
)

func main() {
	println("Fan in........")

	tasks1 := taskGenerator("First ")
	tasks2 := taskGenerator("Second ")
	tasks3 := taskGenerator("Third ")

	out := merge(tasks1, tasks2, tasks3)

	for s := range out {
		println(s)
	}

}

func taskGenerator(s string) chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 0; i < 3; i++ {
			ch <- s + strconv.Itoa(i)
		}
	}()
	return ch
}

// From Samir Ajamni's excellent post on blog.golang
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
