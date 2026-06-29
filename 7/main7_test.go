package main

import (
	"sort"
	"testing"
)

// ───────────────────────────────────────────
// generate
// ───────────────────────────────────────────

func Test_generate_ReturnsAllValues(t *testing.T) {
	ch := generate(1, 2, 3)
	var result []int
	for v := range ch {
		result = append(result, v)
	}

	expected := []int{1, 2, 3}
	if !equalSlices(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func Test_generate_ClosesChannel(t *testing.T) {
	ch := generate(1, 2, 3)
	for range ch {
	}

	// если канал не закрыт — for range завис бы выше
	// доходим сюда — канал закрыт корректно
}

func Test_generate_Empty(t *testing.T) {
	ch := generate()
	var result []int
	for v := range ch {
		result = append(result, v)
	}

	if len(result) != 0 {
		t.Errorf("expected empty, got %v", result)
	}
}

func Test_generate_SingleValue(t *testing.T) {
	ch := generate(42)
	v := <-ch
	if v != 42 {
		t.Errorf("got %d, want 42", v)
	}
}

// ───────────────────────────────────────────
// merge
// ───────────────────────────────────────────

func Test_merge_AllValuesPresent(t *testing.T) {
	ch1 := generate(1, 2, 3)
	ch2 := generate(4, 5, 6)
	ch3 := generate(7, 8, 9)

	var result []int
	for v := range merge(ch1, ch2, ch3) {
		result = append(result, v)
	}

	// порядок не гарантирован — сортируем перед сравнением
	sort.Ints(result)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if !equalSlices(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func Test_merge_CorrectCount(t *testing.T) {
	ch1 := generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	ch2 := generate(11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	ch3 := generate(21, 22, 23, 24, 25, 26, 27, 28, 29, 30)

	count := 0
	for range merge(ch1, ch2, ch3) {
		count++
	}

	if count != 30 {
		t.Errorf("got %d values, want 30", count)
	}
}

func Test_merge_ClosesChannel(t *testing.T) {
	ch1 := generate(1, 2)
	ch2 := generate(3, 4)

	out := merge(ch1, ch2)
	for range out {
	}

	// если канал не закрыт — for range завис бы выше
}

func Test_merge_SingleChannel(t *testing.T) {
	ch := generate(1, 2, 3)

	var result []int
	for v := range merge(ch) {
		result = append(result, v)
	}

	sort.Ints(result)
	expected := []int{1, 2, 3}

	if !equalSlices(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func Test_merge_EmptyChannels(t *testing.T) {
	ch1 := generate()
	ch2 := generate()

	var result []int
	for v := range merge(ch1, ch2) {
		result = append(result, v)
	}

	if len(result) != 0 {
		t.Errorf("expected empty result, got %v", result)
	}
}

func Test_merge_NoChannels(t *testing.T) {
	var result []int
	for v := range merge() {
		result = append(result, v)
	}

	if len(result) != 0 {
		t.Errorf("expected empty result, got %v", result)
	}
}

func Test_merge_MixedSizeChannels(t *testing.T) {
	ch1 := generate(1)
	ch2 := generate(2, 3, 4)
	ch3 := generate(5, 6)

	count := 0
	for range merge(ch1, ch2, ch3) {
		count++
	}

	if count != 6 {
		t.Errorf("got %d values, want 6", count)
	}
}

// ───────────────────────────────────────────
// helpers
// ───────────────────────────────────────────

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
