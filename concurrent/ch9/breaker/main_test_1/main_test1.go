package main

import (
	"fmt"
	"time"
)

type wrapper struct {
	exec bool
}

func (w *wrapper) commandFunc() {
	time.Sleep(1 * time.Millisecond)
	fmt.Println("Executing command.....")
	w.exec = true
}
func (w *wrapper) defaultFunc() {
	fmt.Println("Defaulting command.....")
}
func (w *wrapper) cleanupFunc() {
	fmt.Println("Canceling command.....")
}

type wrapper2 struct {
	exec bool
}

func (w *wrapper2) commandFunc() {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Executing command.....")
	w.exec = true
}
func (w *wrapper2) defaultFunc() {
	fmt.Println("Defaulting command.....")
}
func (w *wrapper2) cleanupFunc() {
	fmt.Println("Canceling command.....")
}
