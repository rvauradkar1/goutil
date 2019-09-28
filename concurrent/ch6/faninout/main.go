package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	println("Fan in........")

	tasksch := faninTasks()

	fanoutTasks(tasksch)

}

func faninTasks() chan string {
	tasks1 := taskGenerator("First ")
	tasks2 := taskGenerator("Second ")
	tasks3 := taskGenerator("Third ")

	out := merge(tasks1, tasks2, tasks3)

	return out
}

func taskGenerator(s string) chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 0; i < 30; i++ {
			ch <- s + strconv.Itoa(i)
		}
	}()
	return ch
}

func fanoutTasks(tasksch chan string) {
	num := 5
	m := sync.Mutex{}
	var wg sync.WaitGroup
	wg.Add(num)
	runs := make(map[int]int, num)
	for i := 0; i < num; i++ {
		go func(ch chan string, i int) {
			defer wg.Done()
			for t := range ch {
				fmt.Println("Routine ", strconv.Itoa(i), "Processing ", t)
				m.Lock()
				runs[i]++
				m.Unlock()
			}
		}(tasksch, i)
	}
	wg.Wait()
	fmt.Println()
	for key, value := range runs {
		fmt.Println("Goroutine ", key, " processed ", value)
	}
	fmt.Println()
}

func tasks() chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- "Task " + strconv.Itoa(i)
		}
		close(ch)
	}()
	return ch
}

// From Samir Ajamni's excellent post on blog.golang
func merge(channels ...chan string) chan string {
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
