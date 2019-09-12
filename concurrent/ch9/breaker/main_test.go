package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_scanner_shutdown(t *testing.T) {
	fmt.Println("Testing Test_scanner_shutdown")
	b := &breaker{}
	b.init("name", time.Second, 0)
	b.shutdown()
	scanner(b)
	if b.status != iShutdown || !b.isShutdown {
		t.Errorf("Shutdown should have initiated")
	}
	c := <-b.shutdownch
	if c != true {
		t.Errorf("Shutdown channel not populated")
	}
}

func scanner2(b *breaker) {
	for {
		fmt.Println("Scanner isShutdown = ", b.isShutdown)
		if b.isShutdown {
			b.status = iShutdown
			return
		}
		time.Sleep(1000 * time.Millisecond)
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
		b.status = iCircuitGood
		fmt.Println("Scanner status = ", b.status)
	}
}

func Test_scanner_close_ciucuit(t *testing.T) {
	b := &breaker{}
	b.init("name", time.Second, 1)
	b.isOk = false
	fmt.Println("starting Test_scanner_close_ciucuit")
	go func() { scanner(b) }()
	time.Sleep(110 * time.Millisecond)
	fmt.Println("Return = ", b.status)
	if b.status != iCircuitGood {
		t.Errorf("Circuit should have been repaired")
	}
	b.shutdown()
	fmt.Println("Return = ", b.status)

}
