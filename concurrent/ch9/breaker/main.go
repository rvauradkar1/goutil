package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/pkg/errors"
)

type wrapper struct {
}

func (w *wrapper) commandFunc() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Executing command.....")
}
func (w *wrapper) defaultFunc() {
	fmt.Println("Defaulting command.....")
}
func (w *wrapper) cleanupFunc() {
	fmt.Println("Canceling command.....")
}

func main() {
	errors.New("")
	fmt.Println("Throttle demo....")
	commands := &wrapper{}
	breaker := &breaker{}
	breaker.init("name", 10*time.Millisecond, 3)
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			err := breaker.execute(commands)
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

type commandFunc func()
type commandFuncs interface {
	commandFunc()
	defaultFunc()
	cleanupFunc()
}

// Handler is a means for client of circuit breaker to provide a method value bound to a struct that will make actual service call
type Handler func() error

// Breaker struct for circuit breaker control parameters
type breaker struct {
	name                string
	timeout             time.Duration
	numConcurrent       int
	semaphore           chan bool
	isOk                bool
	isShutdown          bool
	shutdownch          chan bool
	status              int
	healthCheckInterval time.Duration
}

// New initializes the circuit breaker
func (b *breaker) init(name string, timeout time.Duration, numConcurrent int) {
	b.name = name
	b.timeout = timeout
	b.numConcurrent = numConcurrent
	b.semaphore = make(chan bool, b.numConcurrent)
	b.isOk = true
	b.shutdownch = make(chan bool, 1)
	b.healthCheckInterval = 100
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
		time.Sleep(b.healthCheckInterval * time.Millisecond)
		if !b.isOk {
			select {
			case <-b.shutdownch:
				fmt.Println("aa Shuttind down")
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
		//b.status = iCircuitGood
	}
}

func (b *breaker) openCircuit() bool {
	b.isOk = false
	b.status = iCircuitStillBad
	return b.isOk
}

func (b *breaker) closeCircuit() bool {
	b.isOk = true
	b.status = iCircuitGood
	return b.isOk
}

var mutex = &sync.Mutex{}

func (b *breaker) shutdown() {
	fmt.Println("SHudtting down....")
	b.shutdownch <- true
	fmt.Println("shut down....")
	mutex.Lock()
	b.isShutdown = true
	mutex.Unlock()
	b.status = iShutdown
}

func (b *breaker) execute(commands commandFuncs) chan error {
	errorch := make(chan error, 1)
	go func() {
		select {
		case b.semaphore <- true:
			go func() {
				defer func() { <-b.semaphore }()
				commands.commandFunc()
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
