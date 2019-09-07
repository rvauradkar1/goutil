package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	responses := make([]string, 2)

	go service1(&wg, responses[0:1])

	go service2(&wg, responses[1:])

	wg.Wait()
	println(responses[0], " ", responses[1])

}

func service1(wg *sync.WaitGroup, resp []string) {
	defer wg.Done()
	resp[0] = "Performing Service 1"
}

func service2(wg *sync.WaitGroup, resp []string) {
	defer wg.Done()
	resp[0] = "Performing Service 2"
}
