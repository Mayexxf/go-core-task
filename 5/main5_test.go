package main

import "testing"

// ───────────────────────────────────────────
// int
// ───────────────────────────────────────────

func Test_equalSlices_Int_HasIntersection(t *testing.T) {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	expected := []int{3, 64}

	ok, result := equalSlices(a, b)

	if !ok {
		t.Error("expected ok=true, got false")
	}
	if !equalSlicesHelper(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func Test_equalSlices_Int_NoIntersection(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	ok, result := equalSlices(a, b)

	if ok {
		t.Error("expected ok=false, got true")
	}
	if len(result) != 0 {
		t.Errorf("expected empty result, got %v", result)
	}
}

func Test_equalSlices_Int_OneMatch(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{3, 4, 5}
	expected := []int{3}

	ok, result := equalSlices(a, b)

	if !ok {
		t.Error("expected ok=true, got false")
	}
	if !equalSlicesHelper(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func Test_equalSlices_Int_AllMatch(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}

	ok, result := equalSlices(a, b)

	if !ok {
		t.Error("expected ok=true, got false")
	}
	if !equalSlicesHelper(result, a) {
		t.Errorf("got %v, want %v", result, a)
	}
}

func Test_equalSlices_Int_DuplicatesInA(t *testing.T) {
	a := []int{1, 1, 2}
	b := []int{1}
	expected := []int{1, 1}

	ok, result := equalSlices(a, b)

	if !ok {
		t.Error("expected ok=true, got false")
	}
	if !equalSlicesHelper(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func Test_equalSlices_Int_EmptyA(t *testing.T) {
	a := []int{}
	b := []int{1, 2, 3}

	ok, result := equalSlices(a, b)

	if ok {
		t.Error("expected ok=false, got true")
	}
	if len(result) != 0 {
		t.Errorf("expected empty result, got %v", result)
	}
}

func Test_equalSlices_Int_EmptyB(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{}

	ok, result := equalSlices(a, b)

	if ok {
		t.Error("expected ok=false, got true")
	}
	if len(result) != 0 {
		t.Errorf("expected empty result, got %v", result)
	}
}

func Test_equalSlices_Int_BothEmpty(t *testing.T) {
	ok, result := equalSlices([]int{}, []int{})

	if ok {
		t.Error("expected ok=false, got true")
	}
	if len(result) != 0 {
		t.Errorf("expected empty result, got %v", result)
	}
}

// ───────────────────────────────────────────
// string — проверяем что дженерик работает
// ───────────────────────────────────────────

func Test_equalSlices_String_HasIntersection(t *testing.T) {
	a := []string{"apple", "banana", "cherry"}
	b := []string{"banana", "date"}
	expected := []string{"banana"}

	ok, result := equalSlices(a, b)

	if !ok {
		t.Error("expected ok=true, got false")
	}
	if !equalSlicesHelper(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func Test_equalSlices_String_NoIntersection(t *testing.T) {
	a := []string{"apple", "cherry"}
	b := []string{"banana", "date"}

	ok, result := equalSlices(a, b)

	if ok {
		t.Error("expected ok=false, got true")
	}
	if len(result) != 0 {
		t.Errorf("expected empty result, got %v", result)
	}
}

// ───────────────────────────────────────────
// helpers
// ───────────────────────────────────────────

func equalSlicesHelper[T comparable](a, b []T) bool {
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
