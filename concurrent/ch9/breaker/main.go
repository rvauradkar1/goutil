package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/pkg/errors"
)

func command() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Executing command.....")
}

func main() {
	errors.New("")
	fmt.Println("Throttle demo....")
	breaker := &breaker{}
	breaker.init("name", 10*time.Millisecond, 1)
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			err := breaker.execute(command)
			if c := <-err; c != nil {
				fmt.Println("Err = ", i, "___", c, "___")
			}
		}(i)
	}
	fmt.Println("Number of go routines = ", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Done!!")
}

type breakerFuncs interface {
	commandFunc()
	defaultFunc()
}

type commandFunc func()

// Handler is a means for client of circuit breaker to provide a method value bound to a struct that will make actual service call
type Handler func() error

// Breaker struct for circuit breaker control parameters
type breaker struct {
	name          string
	timeout       time.Duration
	numConcurrent int
	semaphore     chan bool
}

// New initializes the circuit breaker
func (b *breaker) init(name string, timeout time.Duration, numConcurrent int) {
	b.name = name
	b.timeout = timeout
	b.numConcurrent = numConcurrent
	b.semaphore = make(chan bool, b.numConcurrent)
}
func (b *breaker) execute(command commandFunc) chan error {
	errorch := make(chan error, 1)
	go func() {
		select {
		case b.semaphore <- true:
			go func() {
				defer func() { <-b.semaphore }()
				command()
				errorch <- nil
			}()
		default:
			errorch <- errors.New("reached threshold, cannot run your function")
		}
	}()
	return errorch
}

type BreakerError struct {
	Err error
}

func (e *BreakerError) Unwrap() error   { return e.Err }
func (e *BreakerError) Error() string   { return "error reading DNS config: " + e.Err.Error() }
func (e *BreakerError) Timeout() bool   { return false }
func (e *BreakerError) Temporary() bool { return false }
