package main

import (
	//"errors"
	"time"

	"github.com/pkg/errors"
)

func main() {
	errors.New("")
}

// Breaker struct for circuit breaker control parameters
type Breaker struct {
	name          string
	timeout       time.Duration
	numConcurrent int
	semaphore     chan bool
}

// New initializes the circuit breaker
func (b *Breaker) New(name string, timeout time.Duration, numConcurrent int) {
	b.name = name
	b.timeout = timeout
	b.numConcurrent = numConcurrent
	b.semaphore = make(chan bool, b.numConcurrent)
}

// Handler is a means for client of circuit breaker to provide a method value bound to a struct that will make actual service call
type Handler func() error

type DNSConfigError struct {
	Err error
}

func (e *DNSConfigError) Unwrap() error   { return e.Err }
func (e *DNSConfigError) Error() string   { return "error reading DNS config: " + e.Err.Error() }
func (e *DNSConfigError) Timeout() bool   { return false }
func (e *DNSConfigError) Temporary() bool { return false }

//func (b *Breaker) Execute()
