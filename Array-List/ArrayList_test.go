package arraylist

import (
	"testing"
)

func TestArrayListCreation(t *testing.T) {
	al := ArrayList[int]{}
	if !al.IsEmpty() {
		t.Error("New ArrayList should be empty")
	}
	if al.Length() != 0 {
		t.Error("New ArrayList should have length 0")
	}
}

func TestAppend(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		expected []int
	}{
		{
			name:     "Append single value",
			values:   []int{1},
			expected: []int{1},
		},
		{
			name:     "Append multiple values",
			values:   []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Empty list",
			values:   []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			al := ArrayList[int]{}
			for _, v := range tt.values {
				al.Append(v)
			}

			if al.Length() != len(tt.expected) {
				t.Errorf("Expected length %d, got %d", len(tt.expected), al.Length())
			}

			for i, expected := range tt.expected {
				if got := al.Get(i); got != expected {
					t.Errorf("Expected %v at position %d, got %v", expected, i, got)
				}
			}
		})
	}
}

func TestRemove(t *testing.T) {
	al := ArrayList[int]{}

	// Setup test data
	values := []int{1, 2, 3, 4}
	for _, v := range values {
		al.Append(v)
	}

	// Test removing from middle
	removed := al.Remove(1)
	if removed != 2 {
		t.Errorf("Expected removed value 2, got %v", removed)
	}
	if al.Length() != 3 {
		t.Errorf("Expected length 3, got %d", al.Length())
	}
	if al.Get(1) != 3 {
		t.Errorf("Expected value 3 at index 1, got %v", al.Get(1))
	}

	// Test removing from beginning
	removed = al.Remove(0)
	if removed != 1 {
		t.Errorf("Expected removed value 1, got %v", removed)
	}

	// Test removing from end
	removed = al.Remove(al.Length() - 1)
	if removed != 4 {
		t.Errorf("Expected removed value 4, got %v", removed)
	}
}

func TestInsert(t *testing.T) {
	al := ArrayList[int]{}

	// Test insert into empty list
	al.Insert(0, 1)
	if al.Get(0) != 1 {
		t.Errorf("Expected value 1 at index 0, got %v", al.Get(0))
	}

	// Test insert at beginning
	al.Insert(0, 2)
	if al.Get(0) != 2 || al.Get(1) != 1 {
		t.Error("Insert at beginning failed")
	}

	// Test insert in middle
	al.Insert(1, 3)
	if al.Get(0) != 2 || al.Get(1) != 3 || al.Get(2) != 1 {
		t.Error("Insert in middle failed")
	}

	// Test insert at end
	al.Insert(al.Length(), 4)
	if al.Get(al.Length()-1) != 4 {
		t.Error("Insert at end failed")
	}
}

func TestWithStrings(t *testing.T) {
	al := ArrayList[string]{}

	strings := []string{"hello", "world", "!"}
	for _, s := range strings {
		al.Append(s)
	}

	for i, expected := range strings {
		if got := al.Get(i); got != expected {
			t.Errorf("Expected %v at position %d, got %v", expected, i, got)
		}
	}
}

func TestPrint(t *testing.T) {
	// This is more of a visual test, but we can at least ensure it doesn't panic
	al := ArrayList[int]{}
	al.Print() // Should handle empty list

	al.Append(1)
	al.Append(2)
	al.Append(3)
	al.Print() // Should handle populated list
}
