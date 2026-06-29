package main

import "testing"

// ───────────────────────────────────────────
// Add
// ───────────────────────────────────────────

func Test_Add_NewKey(t *testing.T) {
	s := make(StringIntMap)
	s.Add("key1", 1)

	val, ok := s["key1"]
	if !ok {
		t.Fatal("key1 not found after Add")
	}
	if val != 1 {
		t.Errorf("got %d, want 1", val)
	}
}

func Test_Add_OverwriteExistingKey(t *testing.T) {
	s := make(StringIntMap)
	s.Add("key1", 1)
	s.Add("key1", 99)

	if s["key1"] != 99 {
		t.Errorf("got %d, want 99", s["key1"])
	}
}

func Test_Add_MultipleKeys(t *testing.T) {
	s := make(StringIntMap)
	s.Add("a", 1)
	s.Add("b", 2)
	s.Add("c", 3)

	if len(s) != 3 {
		t.Errorf("got len %d, want 3", len(s))
	}
}

// ───────────────────────────────────────────
// Remove
// ───────────────────────────────────────────

func Test_Remove_ExistingKey(t *testing.T) {
	s := make(StringIntMap)
	s.Add("key1", 1)
	s.Remove("key1")

	if _, ok := s["key1"]; ok {
		t.Error("key1 still exists after Remove")
	}
}

func Test_Remove_NonExistingKey(t *testing.T) {
	s := make(StringIntMap)

	// не должно паниковать
	s.Remove("nonexistent")
}

func Test_Remove_DoesNotAffectOtherKeys(t *testing.T) {
	s := make(StringIntMap)
	s.Add("key1", 1)
	s.Add("key2", 2)
	s.Remove("key1")

	if _, ok := s["key2"]; !ok {
		t.Error("key2 was removed along with key1")
	}
}

// ───────────────────────────────────────────
// Copy
// ───────────────────────────────────────────

func Test_Copy_SameElements(t *testing.T) {
	s := make(StringIntMap)
	s.Add("key1", 1)
	s.Add("key2", 2)

	cp := s.Copy()

	for k, v := range s {
		if cp[k] != v {
			t.Errorf("key %q: got %d, want %d", k, cp[k], v)
		}
	}
}

func Test_Copy_IndependentFromOriginal(t *testing.T) {
	s := make(StringIntMap)
	s.Add("key1", 1)

	cp := s.Copy()
	s.Add("key2", 2)

	if _, ok := cp["key2"]; ok {
		t.Error("copy was affected by Add to original")
	}
}

func Test_Copy_MutationDoesNotAffectOriginal(t *testing.T) {
	s := make(StringIntMap)
	s.Add("key1", 1)

	cp := s.Copy()
	cp["key1"] = 999

	if s["key1"] == 999 {
		t.Error("mutating copy affected original")
	}
}

func Test_Copy_EmptyMap(t *testing.T) {
	s := make(StringIntMap)
	cp := s.Copy()

	if len(cp) != 0 {
		t.Errorf("expected empty copy, got len %d", len(cp))
	}
}

// ───────────────────────────────────────────
// Exists
// ───────────────────────────────────────────

func Test_Exists_KeyPresent(t *testing.T) {
	s := make(StringIntMap)
	s.Add("key1", 1)

	if !s.Exists("key1") {
		t.Error("expected key1 to exist")
	}
}

func Test_Exists_KeyAbsent(t *testing.T) {
	s := make(StringIntMap)

	if s.Exists("key1") {
		t.Error("expected key1 to not exist")
	}
}

func Test_Exists_AfterRemove(t *testing.T) {
	s := make(StringIntMap)
	s.Add("key1", 1)
	s.Remove("key1")

	if s.Exists("key1") {
		t.Error("expected key1 to not exist after Remove")
	}
}

func Test_Exists_ZeroValueKey(t *testing.T) {
	s := make(StringIntMap)
	s.Add("key1", 0) // zero value int

	if !s.Exists("key1") {
		t.Error("key with zero value should still exist")
	}
}

// ───────────────────────────────────────────
// Get
// ───────────────────────────────────────────

func Test_Get_ExistingKey(t *testing.T) {
	s := make(StringIntMap)
	s.Add("key1", 42)

	val, ok := s.Get("key1")
	if !ok {
		t.Fatal("expected ok=true for existing key")
	}
	if val != 42 {
		t.Errorf("got %d, want 42", val)
	}
}

func Test_Get_NonExistingKey(t *testing.T) {
	s := make(StringIntMap)

	val, ok := s.Get("missing")
	if ok {
		t.Error("expected ok=false for missing key")
	}
	if val != 0 {
		t.Errorf("expected zero value, got %d", val)
	}
}

func Test_Get_AfterRemove(t *testing.T) {
	s := make(StringIntMap)
	s.Add("key1", 1)
	s.Remove("key1")

	_, ok := s.Get("key1")
	if ok {
		t.Error("expected ok=false after Remove")
	}
}

func Test_Get_ZeroValue(t *testing.T) {
	s := make(StringIntMap)
	s.Add("key1", 0)

	val, ok := s.Get("key1")
	if !ok {
		t.Fatal("expected ok=true for key with zero value")
	}
	if val != 0 {
		t.Errorf("got %d, want 0", val)
	}
}
