// pipeline_test.go
package main

import (
	"reflect"
	"testing"
	"time"
)

// runWithTimeout запускает fn в отдельной горутине и роняет тест,
// если результат не пришёл за d — так ловим дедлоки, а не висим вечно.
func runWithTimeout(t *testing.T, fn func() []float64, d time.Duration) []float64 {
	t.Helper()
	resultCh := make(chan []float64, 1)

	go func() {
		resultCh <- fn()
	}()

	select {
	case res := <-resultCh:
		return res
	case <-time.After(d):
		t.Fatal("RunPipeline не завершился вовремя — похоже на дедлок")
		return nil
	}
}

func TestRunPipeline_Cubes(t *testing.T) {
	input := []uint8{5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	want := make([]float64, len(input))
	for i, v := range input {
		want[i] = float64(v) * float64(v) * float64(v)
	}

	got := runWithTimeout(t, func() []float64 {
		return RunPipeline(input)
	}, 2*time.Second)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("RunPipeline(%v) = %v, want %v", input, got, want)
	}
}

func TestRunPipeline_EmptyInput(t *testing.T) {
	got := runWithTimeout(t, func() []float64 {
		return RunPipeline(nil)
	}, 2*time.Second)

	if len(got) != 0 {
		t.Errorf("RunPipeline(nil) = %v, ожидался пустой слайс", got)
	}
}

func TestRunPipeline_SingleValue(t *testing.T) {
	got := runWithTimeout(t, func() []float64 {
		return RunPipeline([]uint8{3})
	}, 2*time.Second)

	want := []float64{27}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("RunPipeline([3]) = %v, want %v", got, want)
	}
}

func TestRunPipeline_BoundaryValues(t *testing.T) {
	// 0 и максимальное значение uint8 — проверяем, что float64(v)
	// до возведения в степень не даёт переполнения.
	got := runWithTimeout(t, func() []float64 {
		return RunPipeline([]uint8{0, 255})
	}, 2*time.Second)

	want := []float64{0, 255 * 255 * 255}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("RunPipeline([0,255]) = %v, want %v", got, want)
	}
}

func TestRunPipeline_PreservesOrder(t *testing.T) {
	// Один producer и один consumer гарантируют порядок FIFO,
	// явно проверяем, что он не нарушается.
	input := []uint8{9, 1, 5, 2, 8}
	got := runWithTimeout(t, func() []float64 {
		return RunPipeline(input)
	}, 2*time.Second)

	for i, v := range input {
		want := float64(v) * float64(v) * float64(v)
		if got[i] != want {
			t.Errorf("на позиции %d: got %v, want %v", i, got[i], want)
		}
	}
}

func TestRunPipeline_NoGoroutineLeak(t *testing.T) {
	// Запускаем несколько раз подряд — если где-то остаётся
	// висящая горутина или незакрытый канал, таймаут это поймает.
	for i := 0; i < 50; i++ {
		runWithTimeout(t, func() []float64 {
			return RunPipeline([]uint8{1, 2, 3})
		}, 1*time.Second)
	}
}
