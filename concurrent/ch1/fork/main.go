package main

import (
	"sync"
)

// Communicating by sharing
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	responses := make([]string, 2)

	go service1(responses[0:1])

	go service2(responses[1:])

	wg.Wait()
	println(responses[0], "\n", responses[1])

}

func service1(resp []string) {
	defer wg.Done()
	resp[0] = "Performing Service 1"
}

func service2(resp []string) {
	defer wg.Done()
	resp[0] = "Performing Service 2"
}
