package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Demo of OneOf pattern - using first returned value")
	providers := [5]provider{}
	result := returnFirstResult(providers[0:])

	fmt.Println(result)
}

// Code modified from Samir Ajamni's blog at blog.golang.org

func returnFirstResult(providers []provider) string {
	ch := make(chan string)
	for _, p := range providers {
		go func(p provider) {
			select {
			case ch <- p.query():
			default:
				fmt.Println("Defaulting - doing nothing")
			}
		}(p)
	}
	return <-ch
}

type provider struct {
}

func (p provider) query() string {
	r := rand.Int63n(1000)
	time.Sleep(time.Duration(r))
	return "Finished in " + strconv.FormatInt(r, 10)
}
