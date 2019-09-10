package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Throttle demo....")
	command := func() {
		time.Sleep(10 * time.Millisecond)
		fmt.Println("Executing command.....")
	}
	t := &throttler{}
	t.init(command, 10*time.Millisecond, 3)
	for i := 0; i < 5; i++ {
		go func(i int) {
			err := t.execute()
			c := <-err
			fmt.Println("Err = ", i, "___", c, "___", c == nil)
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
	fn := func() {
		fmt.Println("Releaseing.....")
		<-t.semaphore
	}
	go func() {
		select {
		case t.semaphore <- true:
			defer fn()
			t.command()
			errorch <- nil

		default:
			fmt.Println("Reached threshold, cannot run your function")
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
