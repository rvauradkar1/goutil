package breaker

import (
	"fmt"
	"time"
)

type wrapper struct {
	name string
	exec bool
}

func (w *wrapper) CommandFunc() {
	time.Sleep(1 * time.Millisecond)
	fmt.Println("Executing ", w.name)
	w.exec = true
}
func (w *wrapper) DefaultFunc() {
	fmt.Println("Defaulting command.....")
}
func (w *wrapper) CleanupFunc() {
	fmt.Println("Canceling command.....")
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
