package LinkedList

import (
	"testing"
)

func TestLinkedListCreation(t *testing.T) {
	ll := LinkedList[int]{}
	if ll.head != nil {
		t.Error("New LinkedList should have nil head")
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		expected []int
	}{
		{
			name:     "Add single value",
			values:   []int{1},
			expected: []int{1},
		},
		{
			name:     "Add multiple values",
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
			ll := LinkedList[int]{}
			for _, v := range tt.values {
				ll.Add(v)
			}

			// Verify the list contents
			current := ll.head
			for i, expected := range tt.expected {
				if current == nil {
					t.Errorf("List is shorter than expected at index %d", i)
					return
				}
				if current.data != expected {
					t.Errorf("Expected %v at position %d, got %v", expected, i, current.data)
				}
				current = current.next
			}
			if current != nil {
				t.Error("List is longer than expected")
			}
		})
	}
}

func TestAddFront(t *testing.T) {
	ll := LinkedList[int]{}

	// Test adding to empty list
	ll.AddFront(1)
	if ll.head == nil || ll.head.data != 1 {
		t.Error("AddFront failed on empty list")
	}

	// Test adding to non-empty list
	ll.AddFront(2)
	if ll.head == nil || ll.head.data != 2 {
		t.Error("AddFront failed to update head")
	}
	if ll.head.next == nil || ll.head.next.data != 1 {
		t.Error("AddFront failed to maintain existing elements")
	}
}

func TestRemove(t *testing.T) {
	ll := LinkedList[int]{}

	// Test remove on empty list
	ll.Remove()
	if ll.head != nil {
		t.Error("Remove on empty list should keep head nil")
	}

	// Test remove with one element
	ll.Add(1)
	ll.Remove()
	if ll.head != nil {
		t.Error("Remove failed to remove single element")
	}

	// Test remove with multiple elements
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ll.Remove()

	current := ll.head
	expected := []int{1, 2}
	for i, exp := range expected {
		if current == nil {
			t.Errorf("List is shorter than expected at index %d", i)
			return
		}
		if current.data != exp {
			t.Errorf("Expected %v at position %d, got %v", exp, i, current.data)
		}
		current = current.next
	}
}

func TestWithStrings(t *testing.T) {
	ll := LinkedList[string]{}

	// Test that the generic implementation works with strings
	strings := []string{"hello", "world", "!"}
	for _, s := range strings {
		ll.Add(s)
	}

	current := ll.head
	for i, expected := range strings {
		if current == nil {
			t.Errorf("List is shorter than expected at index %d", i)
			return
		}
		if current.data != expected {
			t.Errorf("Expected %v at position %d, got %v", expected, i, current.data)
		}
		current = current.next
	}
}

func TestPrint(t *testing.T) {
	// This is more of a visual test, but we can at least ensure it doesn't panic
	ll := LinkedList[int]{}
	ll.Print() // Should handle empty list

	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ll.Print() // Should handle populated list
}
