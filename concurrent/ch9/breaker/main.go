package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/pkg/errors"
)

func main() {
	errors.New("")
}

type commandFuncs interface {
	commandFunc()
	defaultFunc()
	cleanupFunc()
}

// Handler is a means for client of circuit breaker to provide a method value bound to a struct that will make actual service call
type Handler func() error

// Breaker struct for circuit breaker control parameters
type breaker struct {
	name                string        // For debudding purposes
	timeout             time.Duration // Timeout for each task
	numConcurrent       int           // Number of concurrent request taken
	semaphore           chan bool     // Controls access to execute tasks
	isOk                bool          // Can circuit take more load?
	isShutdown          bool          // Has circuit been shutdown completely?
	shutdownch          chan bool     // Used to kill healthcheck goroutine when shutdown
	status              int           // States that a ciucuit can go through
	healthCheckInterval time.Duration // Scanning interval to reset circuit if it has been tripped
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
	iCircuitStillBad = 2
	iCircuitGood     = 3
)

func scanner(b *breaker) {
	for {
		if b.isShutdown {
			b.status = iShutdown
			return
		}
		time.Sleep(b.healthCheckInterval * time.Millisecond)
		if !b.isOk {
			select {
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
	if b.isShutdown {
		return
	}
	b.shutdownch <- true
	fmt.Println("shut down....")
	mutex.Lock()
	b.isShutdown = true
	mutex.Unlock()
	b.status = iShutdown
}

func (b *breaker) execute(commands commandFuncs) chan error {
	errorch := make(chan error, 1)
	if b.isShutdown {
		be := &BreakerError{Err: errors.New("cicuit has been permanently shutdown. create a new one")}
		errorch <- be
		return errorch
	}
	go func() {
		select {
		case b.semaphore <- true:
			go func() {
				defer func() { <-b.semaphore }()
				commands.commandFunc()
				errorch <- nil
			}()
		default:
			commands.defaultFunc()
			commands.cleanupFunc()
			b.openCircuit()
			errorch <- errors.New("reached threshold, cannot run your function")
		}
	}()
	return errorch
}

type BreakerError struct {
	Err       error
	isTimeout bool
}

func (b *BreakerError) Unwrap() error  { return b.Err }
func (b *BreakerError) Error() string  { return b.Err.Error() }
func (b *BreakerError) Timeout() bool  { return b.isTimeout }
func (b *BreakerError) Shutdown() bool { return false }
