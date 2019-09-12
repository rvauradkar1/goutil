package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_scanner_shutdown(t *testing.T) {
	b := &breaker{}
	b.init("name", time.Second, 1)
	b.shutdown()
	scanner(b)
	r := <-b.shutdownch
	if r != true {
		t.Errorf("Circuit should be closed")
	}
	if b.isShutdown != true {
		t.Errorf("Circuit should be closed")
	}
}

func Test_scanner(t *testing.T) {
	b := &breaker{}
	b.init("name", time.Second, 1)
	b.shutdown()
	scanner(b)
	r := <-b.shutdownch
	if r != true {
		t.Errorf("Circuit should be closed")
	}
	if b.isShutdown != true {
		t.Errorf("Circuit should be closed")
	}
}

func scanner1(b *breaker) {
	for {
		fmt.Println("Shiw = ", b.isShutdown)
		if b.isShutdown {
			return
		}
		time.Sleep(1000 * time.Millisecond)

		if !b.isOk {
			select {
			case <-b.shutdownch:
				fmt.Println("Shuttind down")
				return
			case b.semaphore <- true:
				<-b.semaphore
				b.closeCircuit()
				fmt.Println("Resetting circuit")
			default:
				fmt.Println("Circuit still bad!!!")
			}
		}
	}
}
