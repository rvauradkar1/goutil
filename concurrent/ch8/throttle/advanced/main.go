package main

import (
	"errors"
	"fmt"
	"time"
)

func command() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Executing command.....")
}

func main() {
	fmt.Println("Throttle demo....")
	t := &throttler{}
	t.init(command, 10*time.Millisecond, 3)
	for i := 0; i < 5; i++ {
		go func(i int) {
			err := t.execute()
			if c := <-err; err != nil {
				fmt.Println("Err = ", i, "___", c, "___")
			}
		}(i)
	}
	time.Sleep(time.Second)
}

type throttler struct {
	semaphore chan bool
	limit     int
	timeout   time.Duration
	command   commandFunc
}

type commandFunc func()

func (t *throttler) execute() chan error {
	errorch := make(chan error, 1)
	go func() {
		select {
		case t.semaphore <- true:
			defer func() { <-t.semaphore }()
			t.command()
			errorch <- nil

		default:
			errorch <- errors.New("reached threshold, cannot run your function")
		}
	}()
	return errorch
}
func (t *throttler) init(command commandFunc, timeout time.Duration, limit int) {
	t.command = command
	t.timeout = timeout
	t.limit = limit
	t.semaphore = make(chan bool, limit)
}
