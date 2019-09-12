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
	breaker.init("name", 10*time.Millisecond, 3)
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
	fmt.Printf("isOk = %v\n", breaker.isOk)
	time.Sleep(2 * time.Second)
	fmt.Printf("isOk = %v\n", breaker.isOk)
	breaker.shutdown()
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
	isOk          bool
	isShutdown    bool
	shutdownch    chan bool
	status        int
}

// New initializes the circuit breaker
func (b *breaker) init(name string, timeout time.Duration, numConcurrent int) {
	b.name = name
	b.timeout = timeout
	b.numConcurrent = numConcurrent
	b.semaphore = make(chan bool, b.numConcurrent)
	b.isOk = true
	b.shutdownch = make(chan bool, 1)
	go scanner(b)
}

const (
	iShutdown        = 0
	iShuttingDown    = 1
	iCircuitStillBad = 2
	iCircuitGood     = 3
)

func scanner(b *breaker) {
	for {
		fmt.Println("Scanner isShutdown = ", b.isShutdown)
		if b.isShutdown {
			b.status = iShutdown
			return
		}
		time.Sleep(100 * time.Millisecond)
		if !b.isOk {
			select {
			case <-b.shutdownch:
				fmt.Println("Shuttind down")
				b.status = iShuttingDown
			case b.semaphore <- true:
				<-b.semaphore
				b.closeCircuit()
				fmt.Println("Resetting circuit")
				b.status = iCircuitGood
			default:
				fmt.Println("Circuit still bad!!!")
				b.status = iCircuitStillBad

			}
		}
		fmt.Println("Scanner status = ", b.status)
	}
}

func (b *breaker) openCircuit() bool {
	b.isOk = false
	return b.isOk
}

func (b *breaker) closeCircuit() bool {
	b.isOk = true
	return b.isOk
}

func (b *breaker) shutdown() {
	fmt.Println("SHudtting down....")
	b.shutdownch <- true
	fmt.Println("shut down....")
	b.isShutdown = true
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
			b.openCircuit()
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
