package main

import (
	"errors"
	"testing"
)

// ───────────────────────────────────────────
// sliceExample
// ───────────────────────────────────────────

func Test_sliceExample_FilterOdd(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	result := sliceExample(input, func(v int) bool { return v%2 != 0 })
	expected := []int{2, 4}

	if !equalSlices(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func Test_sliceExample_DoesNotModifyOriginal(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	original := make([]int, len(input))
	copy(original, input)

	sliceExample(input, func(v int) bool { return v%2 != 0 })

	if !equalSlices(input, original) {
		t.Errorf("original slice was modified: got %v, want %v", input, original)
	}
}

func Test_sliceExample_AllMatch(t *testing.T) {
	input := []int{1, 3, 5}
	result := sliceExample(input, func(v int) bool { return v%2 != 0 })

	if len(result) != 0 {
		t.Errorf("expected empty slice, got %v", result)
	}
}

func Test_sliceExample_NoneMatch(t *testing.T) {
	input := []int{2, 4, 6}
	result := sliceExample(input, func(v int) bool { return v%2 != 0 })

	if !equalSlices(result, input) {
		t.Errorf("got %v, want %v", result, input)
	}
}

func Test_sliceExample_EmptyInput(t *testing.T) {
	result := sliceExample([]int{}, func(v int) bool { return v%2 != 0 })

	if len(result) != 0 {
		t.Errorf("expected empty slice, got %v", result)
	}
}

// ───────────────────────────────────────────
// copySlice
// ───────────────────────────────────────────

func Test_copySlice_IndependentFromOriginal(t *testing.T) {
	original := []int{1, 2, 3}
	cp := copySlice(original)

	original = append(original, 100)

	if len(cp) == len(original) {
		t.Errorf("copy should not be affected by append to original")
	}
}

func Test_copySlice_SameElements(t *testing.T) {
	original := []int{10, 20, 30}
	cp := copySlice(original)

	if !equalSlices(original, cp) {
		t.Errorf("got %v, want %v", cp, original)
	}
}

func Test_copySlice_MutationIndependent(t *testing.T) {
	original := []int{1, 2, 3}
	cp := copySlice(original)

	cp[0] = 999

	if original[0] == 999 {
		t.Error("mutating copy affected original — shared backing array")
	}
}

func Test_copySlice_Empty(t *testing.T) {
	cp := copySlice([]int{})

	if len(cp) != 0 {
		t.Errorf("expected empty slice, got %v", cp)
	}
}

// ───────────────────────────────────────────
// addElements
// ───────────────────────────────────────────

func Test_addElements_AppendsToEnd(t *testing.T) {
	input := []int{1, 2, 3}
	result := addElements(input, 99)
	expected := []int{1, 2, 3, 99}

	if !equalSlices(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func Test_addElements_DoesNotModifyOriginal(t *testing.T) {
	input := []int{1, 2, 3}
	original := make([]int, len(input))
	copy(original, input)

	addElements(input, 99)

	if !equalSlices(input, original) {
		t.Errorf("original was modified: got %v, want %v", input, original)
	}
}

func Test_addElements_ToEmptySlice(t *testing.T) {
	result := addElements([]int{}, 42)
	expected := []int{42}

	if !equalSlices(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

// ───────────────────────────────────────────
// removeElement
// ───────────────────────────────────────────

func Test_removeElement_MiddleIndex(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	result, err := removeElement(input, 2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	expected := []int{1, 2, 4, 5}

	if !equalSlices(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func Test_removeElement_FirstIndex(t *testing.T) {
	input := []int{1, 2, 3}
	result, err := removeElement(input, 0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	expected := []int{2, 3}

	if !equalSlices(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func Test_removeElement_LastIndex(t *testing.T) {
	input := []int{1, 2, 3}
	result, err := removeElement(input, len(input))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	expected := []int{1, 2}

	if !equalSlices(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func Test_removeElement_OutOfRange(t *testing.T) {
	input := []int{1, 2, 3}
	result, err := removeElement(input, 999)
	if err != nil {
		if errors.Is(err, ErrIndexOutOfRange) {
			t.Logf("Log: %v", err)
		}
	}
	expected := []int{1, 2, 3}

	if !equalSlices(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func Test_removeElement_DoesNotModifyOriginal(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	original := make([]int, len(input))
	copy(original, input)

	removeElement(input, 2)

	if !equalSlices(input, original) {
		t.Errorf("original was modified: got %v, want %v", input, original)
	}
}

// ───────────────────────────────────────────
// helpers
// ───────────────────────────────────────────

func equalSlices[T comparable](a, b []T) bool {
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
