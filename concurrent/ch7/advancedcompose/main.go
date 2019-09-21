package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	ch, _ := task1()
	println("called task1 ", ch)
	for s := range ch {
		fmt.Println(s.s, "---", s.err)
	}
}

type resp struct {
	s   string
	err error
}

func task1() (<-chan resp, error) {
	out := make(chan resp)
	go func() {
		defer close(out)
		for i := 0; i < 5; i++ {
			if i == 3 {
				s1 := resp{err: errors.Errorf("Service failure on %d", i)}
				out <- s1
			} else {
				out <- resp{s: fmt.Sprintf("Performing Task %d", i)}
			}
		}
	}()
	return out, nil
}

func task2(in <-chan resp) (<-chan string, chan<- error, error) {
	out := make(chan string)
	go func() {
		defer close(out)
		for s := range in {
			if s.err != nil {
				fmt.Println(s.err.Error())
			} else {
				out <- fmt.Sprintf(s.s + "\n" + "Performing Service \n")
			}
		}
	}()
	return out, nil, nil
}
