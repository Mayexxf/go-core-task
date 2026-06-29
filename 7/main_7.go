package main

import (
	"fmt"
	"sync"
)

func main() {

	ch1 := generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	ch2 := generate(11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	ch3 := generate(21, 22, 23, 24, 25, 26, 27, 28, 29, 30)

	for ch := range merge(ch1, ch2, ch3) {
		fmt.Println(ch)
	}

}

func generate(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func merge(chs ...<-chan int) <-chan int {
	resultChan := make(chan int)
	wg := &sync.WaitGroup{}

	for _, ch := range chs {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for v := range ch {
				resultChan <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	return resultChan
}
