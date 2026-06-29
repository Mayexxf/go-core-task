package main

import (
	"testing"
)

func Test_equalRows_BasicDifference(t *testing.T) {
	a := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	b := []string{"banana", "date", "fig"}
	expected := []string{"apple", "cherry", "43", "lead", "gno1"}

	result := equalRows(a, b)

	if !equalSlices(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func Test_equalRows_NoIntersection(t *testing.T) {
	a := []string{"apple", "cherry"}
	b := []string{"banana", "date"}

	result := equalRows(a, b)

	if !equalSlices(result, a) {
		t.Errorf("got %v, want %v", result, a)
	}
}

func Test_equalRows_AllIntersect(t *testing.T) {
	a := []string{"apple", "banana"}
	b := []string{"apple", "banana"}

	result := equalRows(a, b)

	if len(result) != 0 {
		t.Errorf("expected empty slice, got %v", result)
	}
}

func Test_equalRows_EmptyA(t *testing.T) {
	a := []string{}
	b := []string{"banana", "date"}

	result := equalRows(a, b)

	if len(result) != 0 {
		t.Errorf("expected empty slice, got %v", result)
	}
}

func Test_equalRows_EmptyB(t *testing.T) {
	a := []string{"apple", "banana"}
	b := []string{}

	result := equalRows(a, b)

	if !equalSlices(result, a) {
		t.Errorf("got %v, want %v", result, a)
	}
}

func Test_equalRows_BothEmpty(t *testing.T) {
	result := equalRows([]string{}, []string{})

	if len(result) != 0 {
		t.Errorf("expected empty slice, got %v", result)
	}
}

func Test_equalRows_Duplicates(t *testing.T) {
	a := []string{"apple", "apple", "banana"}
	b := []string{"banana"}
	expected := []string{"apple", "apple"}

	result := equalRows(a, b)

	if !equalSlices(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func Test_equalRows_ExtraElementsInB(t *testing.T) {
	a := []string{"apple"}
	b := []string{"apple", "banana", "cherry", "date"}

	result := equalRows(a, b)

	if len(result) != 0 {
		t.Errorf("expected empty slice, got %v", result)
	}
}

func Test_equalRows_DoesNotModifyInputs(t *testing.T) {
	a := []string{"apple", "banana"}
	b := []string{"banana"}

	originalA := make([]string, len(a))
	originalB := make([]string, len(b))
	copy(originalA, a)
	copy(originalB, b)

	equalRows(a, b)

	if !equalSlices(a, originalA) {
		t.Errorf("slice a was modified: got %v, want %v", a, originalA)
	}
	if !equalSlices(b, originalB) {
		t.Errorf("slice b was modified: got %v, want %v", b, originalB)
	}
}

func equalSlices(a, b []string) bool {
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
