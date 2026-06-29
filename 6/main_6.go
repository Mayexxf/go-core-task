package main

import (
	"fmt"
	"math/rand"
)

func main() {
	done := make(chan struct{})
	gen := generator(done)
	fmt.Println(<-gen)
	close(done)
}

func generator(done <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case out <- rand.Intn(100):
			}
		}
	}()
	return out
}
