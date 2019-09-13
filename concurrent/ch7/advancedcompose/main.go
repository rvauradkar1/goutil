package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	ch, _ := service1("Start...")
	println("called service1 ", ch)
	ch1, _, _ := service2(ch)
	println("called service2 ", ch1)
	for s := range ch1 {
		println(s)
	}
}

type s1Resp struct {
	s   string
	err error
}

func service1(s string) (<-chan s1Resp, error) {
	out := make(chan s1Resp)
	go func() {
		defer close(out)
		for i := 0; i < 5; i++ {
			if i == 3 {
				s1 := s1Resp{err: errors.Errorf("Service failure on %d", i)}
				out <- s1
				return
			}
			out <- s1Resp{s: fmt.Sprintf(s+"\n"+"Performing Service %d\n", i)}
		}
	}()
	return out, nil
}

func service2(in <-chan s1Resp) (<-chan string, chan<- error, error) {
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
