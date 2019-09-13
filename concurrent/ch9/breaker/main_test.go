package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/pkg/errors"
)

func Test_is_ok(t *testing.T) {
	fmt.Println("Testing Test_is_ok")
	b := &breaker{}
	b.init("name", time.Second, 0)
	b.shutdown()
	if b.isOk != true {
		t.Errorf("Circuite should have been ok")
	}
}
func Test_shutdown(t *testing.T) {
	fmt.Println("Testing Test_shutdown")
	b := &breaker{}
	b.isOk = false
	b.init("name", time.Second, 0)
	b.healthCheckInterval = 5
	b.shutdown()
	if b.status != iShutdown || !b.isShutdown {
		t.Errorf("Shutdown should have initiated")
	}
}

func Test_scanner_circuit_repaired(t *testing.T) {
	b := &breaker{}
	b.init("name", time.Second, 0)
	b.isOk = false
	b.healthCheckInterval = 10
	fmt.Println("starting Test_scanner_circuit_repaired")
	time.Sleep(15 * time.Millisecond)
	fmt.Println("Return = ", b.status)
	if b.status != iCircuitStillBad {
		t.Errorf("Circuit should have been repaired")
	}
	b.shutdown()
	fmt.Println("Return = ", b.status)
}

func Test_execute_t(t *testing.T) {
	errors.New("")
	fmt.Println("Throttle demo....")
	commands := &wrapper{}
	b := &breaker{}
	b.init("name", 10*time.Millisecond, 3)
	b.healthCheckInterval = 1000
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			err := b.execute(commands)
			if c := <-err; c != nil {
				fmt.Println("Err = ", i, "___", c, "___")
			}
		}(i)
	}
	fmt.Println("Number of go routines = ", runtime.NumGoroutine())
	wg.Wait()
	fmt.Printf("isOk 1 = %v\n", b.isOk)
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("isOk 2 = %v\n", b.isOk)
	b.shutdown()
	fmt.Println("Done!!")
}

func Test_execute_1(t *testing.T) {
	errors.New("")
	fmt.Println("Throttle demo....")
	b := &breaker{}
	b.init("name", 10*time.Millisecond, 3)
	if !b.isOk {
		t.Errorf("Circuit should be ok")
	}
	if b.isShutdown {
		t.Errorf("Circuit should NOT be shutdown")
	}
	b.shutdown()
}

func Test_execute_2(t *testing.T) {
	errors.New("")
	fmt.Println("Running Test_execute_2 demo....")
	b := &breaker{}
	b.init("name", 10*time.Millisecond, 3)
	w := &wrapper{exec: false}
	b.execute(w)
	time.Sleep(5 * time.Millisecond)
	if !w.exec {
		t.Errorf("Service should have executed")
	}
	b.shutdown()
}

func Test_execute_exceed_limit_wait_till_circuit_ok(t *testing.T) {
	errors.New("")
	fmt.Println("Running Test_execute_exceed_limit_wait_till_circuit_ok demo....")
	b := &breaker{}
	b.init("name", 10*time.Millisecond, 3)
	b.healthCheckInterval = 1000
	w := &wrapper2{exec: false}
	b.execute(w)
	w = &wrapper2{exec: false}
	b.execute(w)
	w = &wrapper2{exec: false}
	b.execute(w)
	time.Sleep(999 * time.Millisecond)
	if !w.exec {
		t.Errorf("Service should have executed")
	}
	b.shutdown()
}
func Test_scanner_circuit_reset(t *testing.T) {
	b := &breaker{}
	b.init("name", time.Second, 1)
	b.isOk = false
	b.healthCheckInterval = 10
	fmt.Println("starting Test_scanner_circuit_still_bad")
	time.Sleep(15 * time.Millisecond)
	fmt.Println("Return = ", b.status)
	if b.status != iCircuitGood {
		t.Errorf("Circuit should have been repaired")
	}
	b.shutdown()
	fmt.Println("Return = ", b.status)
}
