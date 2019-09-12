package main

import (
	"fmt"
	"testing"
	"time"
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

func Test_scanner_circuit_still_bad(t *testing.T) {
	b := &breaker{}
	b.init("name", time.Second, 0)
	b.isOk = false
	b.healthCheckInterval = 10
	fmt.Println("starting Test_scanner_circuit_still_bad")
	time.Sleep(15 * time.Millisecond)
	fmt.Println("Return = ", b.status)
	if b.status != iCircuitStillBad {
		t.Errorf("Circuit should have been repaired")
	}
	b.shutdown()
	fmt.Println("Return = ", b.status)
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
