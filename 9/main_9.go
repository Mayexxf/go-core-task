// pipeline.go
package main

import (
	"fmt"
	"math"
	"sync"
)

func ConveyerBelt(chUint <-chan uint8, chFloat chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range chUint {
		chFloat <- math.Pow(float64(v), 3)
	}
}

// RunPipeline прогоняет values через конвейер и возвращает кубы значений
// в том же порядке, в каком они были отправлены.
func RunPipeline(values []uint8) []float64 {
	wg := &sync.WaitGroup{}
	ch1 := make(chan uint8)
	ch2 := make(chan float64)

	wg.Add(1)
	go ConveyerBelt(ch1, ch2, wg)

	go func() {
		wg.Wait()
		close(ch2)
	}()

	go func() {
		for _, v := range values {
			ch1 <- v
		}
		close(ch1)
	}()

	var result []float64
	for v := range ch2 {
		result = append(result, v)
	}
	return result
}

func main() {
	result := RunPipeline(makeRange(5, 15))
	for _, v := range result {
		fmt.Println(v)
	}
}

func makeRange(start, end uint8) []uint8 {
	r := make([]uint8, 0, end-start)
	for i := start; i < end; i++ {
		r = append(r, i)
	}
	return r
}
