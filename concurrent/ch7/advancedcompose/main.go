package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println(cancel)
	ch, _ := task1(ctx)
	fmt.Println("called task1 ", ch)
	for s := range ch {
		select {
		case <-ctx.Done():
			fmt.Println("done!!!!")
		default:
			fmt.Println(s.s, "---", s.err)
			if e := s.err; e != nil {
				strings.Contains(s.err.Error(), "failure")
				fmt.Println("calling cancel")
				cancel()
			}
		}
	}
	time.Sleep(time.Second)
}

func task1(ctx context.Context) (<-chan resp1, error) {
	out := make(chan resp1)
	go func() {
		defer close(out)
		for i := 0; i < 5; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("Done, returning")
				return
			default:
				if i == 2 {
					s1 := resp1{err: errors.Errorf("Service failure on %d", i)}
					fmt.Println("not ok ", i)
					out <- s1
				} else {
					fmt.Println("ok ", i)
					out <- resp1{s: fmt.Sprintf("Performing Task %d", i)}
				}
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
				out <- resp2{s: fmt.Sprintf("Performing Task ")}
			}
		}
	}()
	return out, nil
}

type resp1 struct {
	s   string
	err error
}

type resp2 struct {
	s   string
	err error
}
