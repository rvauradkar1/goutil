package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	println("Fan out........")
	ch := tasks()
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
		}(ch, i)
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
