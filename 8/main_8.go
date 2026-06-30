package main

import (
	"fmt"
	"sync"
	"time"
)

type CustomerWaitGroup struct {
	sem chan struct{}
	idx int
	mu  sync.Mutex
}

func NewCustomerWaitGroup() *CustomerWaitGroup {
	return &CustomerWaitGroup{
		sem: make(chan struct{}),
	}
}

func main() {
	wg := NewCustomerWaitGroup()

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Second * 2)
			fmt.Printf("worker %d done\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("done")
}

func (cwg *CustomerWaitGroup) Add(delta int) {
	cwg.mu.Lock()
	defer cwg.mu.Unlock()
	cwg.idx += delta
}

func (cwg *CustomerWaitGroup) Done() {
	cwg.mu.Lock()
	defer cwg.mu.Unlock()

	cwg.idx--

	if cwg.idx < 0 {
		panic("negative index")
	}

	if cwg.idx == 0 {
		close(cwg.sem)
	}
}

func (cwg *CustomerWaitGroup) Wait() {
	<-cwg.sem
}
