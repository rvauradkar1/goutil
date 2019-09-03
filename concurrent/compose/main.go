package main

import "fmt"

func main() {
	ch := service1("Start...")
	println("called service1 ", ch)
	ch1 := service2(ch)
	println("called service2 ", ch1)
	for s := range ch1 {
		println(s)
	}
}

func service1(s string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 0; i < 5; i++ {
			out <- fmt.Sprintf(s+"\n"+"Performing Service %d\n", i)
		}
	}()
	return out
}

func service2(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for s := range in {
			out <- fmt.Sprintf(s + "\n" + "Performing Service \n")
		}
	}()
	return out
}
