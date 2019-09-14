package main

import (
	"errors"
	"fmt"
	"testing"
	"time"

	//"github.com/rvauradkar1/goutil/concurrent/ch9/breaker/breaker"
	"github.com/rvauradkar1/goutil/concurrent/ch9/breaker/breaker"
)

func main() {

}

func Test_Execute_exceed_limit_wait_till_circuit_ok(t *testing.T) {
	errors.New("")
	fmt.Println("Running Test_execute_exceed_limit_wait_till_circuit_ok demo....")
	b := &breaker.Breaker{}
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

	b.Shutdown()
}

type wrapper2 struct {
	name string
	exec bool
}

func (w *wrapper2) CommandFunc() {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Executing ", w.name)
	w.exec = true
}
func (w *wrapper2) DefaultFunc() {
	fmt.Println("Defaulting ", w.name)
}
func (w *wrapper2) CleanupFunc() {
	fmt.Println("Cleaning ", w.name)
}
