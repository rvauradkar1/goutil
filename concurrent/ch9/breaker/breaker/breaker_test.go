package breaker

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
	b := &Breaker{}
	b.Init("name", time.Second, 0)
	b.Shutdown()
	if b.isOk != true {
		t.Errorf("Circuite should have been ok")
	}
}
func Test_Shutdown(t *testing.T) {
	fmt.Println("Testing Test_Shutdown")
	b := &Breaker{}
	b.isOk = false
	b.Init("name", time.Second, 0)
	b.HealthCheckInterval = 5
	b.Shutdown()
	if b.status != iShutdown || !b.isShutdown {
		t.Errorf("Shutdown should have initiated")
	}
}

func Test_scanner_circuit_repaired(t *testing.T) {
	b := &Breaker{}
	b.Init("name", time.Second, 0)
	b.isOk = false
	b.HealthCheckInterval = 10
	fmt.Println("starting Test_scanner_circuit_repaired")
	time.Sleep(15 * time.Millisecond)
	fmt.Println("Return = ", b.status)
	if b.status != iCircuitStillBad {
		t.Errorf("Circuit should have been repaired")
	}
	b.Shutdown()
	fmt.Println("Return = ", b.status)
}

func Test_Execute_t(t *testing.T) {
	errors.New("")
	fmt.Println("Throttle demo....")
	commands := &wrapper{}
	b := &Breaker{}
	b.Init("name", 10*time.Millisecond, 3)
	b.HealthCheckInterval = 1000
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			err := b.Execute(commands)
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
	b.Shutdown()
	fmt.Println("Done!!")
}

func Test_Execute_1(t *testing.T) {
	errors.New("")
	fmt.Println("Throttle demo....")
	b := &Breaker{}
	b.Init("name", 10*time.Millisecond, 3)
	if !b.isOk {
		t.Errorf("Circuit should be ok")
	}
	if b.isShutdown {
		t.Errorf("Circuit should NOT be Shutdown")
	}
	b.Shutdown()
}

func Test_Execute_2(t *testing.T) {
	errors.New("")
	fmt.Println("Running Test_Execute_2 demo....")
	b := &Breaker{}
	b.Init("name", 10*time.Millisecond, 3)
	w := &wrapper{exec: false}
	b.Execute(w)
	time.Sleep(5 * time.Millisecond)
	if !w.exec {
		t.Errorf("Service should have executed")
	}
	b.Shutdown()
}

func Test_Execute_exceed_limit_wait_till_circuit_ok(t *testing.T) {
	errors.New("")
	fmt.Println("Running Test_execute_exceed_limit_wait_till_circuit_ok demo....")
	b := &Breaker{}
	b.Init("name", 10*time.Millisecond, 3)
	b.HealthCheckInterval = 1000
	w1 := &wrapper2{"Service 1", false}
	b.Execute(w1)
	w2 := &wrapper2{"Service 2", false}
	b.Execute(w2)
	w3 := &wrapper2{"Service 3", false}
	b.Execute(w3)
	w4 := &wrapper2{"Service 4", false}
	b.Execute(w4)
	w5 := &wrapper2{"Service 5", false}
	b.Execute(w5)
	time.Sleep(2020 * time.Millisecond)
	if !b.isOk {
		t.Errorf("Circuit should have been repaired")
	}

	b.Shutdown()
}

func Test_execute_exceed_limit_wait_tillok_submit_more(t *testing.T) {
	errors.New("")
	fmt.Println("Running Test_execute_exceed_limit_wait_till_circuit_ok demo....")
	b := &Breaker{}
	b.Init("name", 10*time.Millisecond, 3)
	b.HealthCheckInterval = 1000
	w1 := &wrapper2{"Service 1", false}
	b.Execute(w1)
	w2 := &wrapper2{"Service 2", false}
	b.Execute(w2)
	w3 := &wrapper2{"Service 3", false}
	b.Execute(w3)
	w4 := &wrapper2{"Service 4", false}
	b.Execute(w4)
	w5 := &wrapper2{"Service 5", false}
	b.Execute(w5)
	time.Sleep(2020 * time.Millisecond)
	if !b.isOk {
		t.Errorf("Circuit should have been repaired")
	}

	b.Shutdown()
}
func Test_scanner_circuit_multipl_Shutdown(t *testing.T) {
	b := &Breaker{}
	b.Init("name", time.Second, 1)
	b.isOk = false
	b.HealthCheckInterval = 10
	fmt.Println("starting Test_scanner_circuit_multipl_Shutdown")
	time.Sleep(15 * time.Millisecond)
	fmt.Println("Return = ", b.status)
	if b.status != iCircuitGood {
		t.Errorf("Circuit should have been repaired")
	}
	b.Shutdown()
	fmt.Println("Return = ", b.status)
	b.Shutdown()
}