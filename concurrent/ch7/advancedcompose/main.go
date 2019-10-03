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

type resp1 struct {
	s   string
	err error
}

type resp2 struct {
	s   string
	err error
}

func task1() (<-chan resp1, error) {
	out := make(chan resp1)
	go func() {
		defer close(out)
		for i := 0; i < 5; i++ {
			if i == 3 {
				s1 := resp1{err: errors.Errorf("Service failure on %d", i)}
				out <- s1
			} else {
				out <- resp1{s: fmt.Sprintf("Performing Task %d", i)}
			}
		}
	}()
	return out, nil
}

func task2(in <-chan resp1) (<-chan resp2, error) {
	out := make(chan resp2)
	go func() {
		defer close(out)
		for s := range in {
			if s.err != nil {
				fmt.Println(s.err.Error())
			} else {
				out <- resp2{s: fmt.Sprintf("Performing Task %d", i)}
			}
		}
	}()
	return out, nil
}
