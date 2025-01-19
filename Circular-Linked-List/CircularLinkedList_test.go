package CircularLinkedList

import (
	"testing"
)

func TestLinkedListCreation(t *testing.T) {
	ll := CircularLinkedList[int]{}
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
			ll := CircularLinkedList[int]{}
			for _, v := range tt.values {
				ll.Add(v)
			}

			// Verify the list contents and circular nature
			if len(tt.expected) == 0 {
				if ll.head != nil {
					t.Error("Empty list should have nil head")
				}
				return
			}

			// Start at head and count elements until we get back to head
			current := ll.head
			count := 0
			values := make([]int, 0)

			for current != nil && (count == 0 || current != ll.head) {
				values = append(values, current.data)
				current = current.next
				count++
			}

			// Verify we got back to head (circular property)
			if current != ll.head {
				t.Error("List is not circular - did not get back to head")
			}

			// Verify the values
			if len(values) != len(tt.expected) {
				t.Errorf("Expected %d elements, got %d", len(tt.expected), len(values))
			}

			for i, exp := range tt.expected {
				if values[i] != exp {
					t.Errorf("Expected %v at position %d, got %v", exp, i, values[i])
				}
			}
		})
	}
}

func TestAddFront(t *testing.T) {
	ll := CircularLinkedList[int]{}

	// Test adding to empty list
	ll.AddFront(1)
	if ll.head == nil || ll.head.data != 1 {
		t.Error("AddFront failed on empty list")
	}
	if ll.head.next != ll.head {
		t.Error("Single element should point to itself")
	}

	// Test adding to non-empty list
	ll.AddFront(2)
	if ll.head == nil || ll.head.data != 2 {
		t.Error("AddFront failed to update head")
	}
	if ll.head.next == nil || ll.head.next.data != 1 {
		t.Error("AddFront failed to maintain existing elements")
	}
	if ll.head.next.next != ll.head {
		t.Error("List is not circular after AddFront")
	}
}

func TestRemove(t *testing.T) {
	ll := CircularLinkedList[int]{}

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

	// Verify remaining elements and circular nature
	if ll.head == nil {
		t.Error("Head should not be nil after removing from multiple elements")
	}

	values := make([]int, 0)
	current := ll.head
	for current != nil && (len(values) == 0 || current != ll.head) {
		values = append(values, current.data)
		current = current.next
	}

	expected := []int{1, 2}
	if len(values) != len(expected) {
		t.Errorf("Expected %d elements after remove, got %d", len(expected), len(values))
	}

	for i, exp := range expected {
		if values[i] != exp {
			t.Errorf("Expected %v at position %d, got %v", exp, i, values[i])
		}
	}

	if current != ll.head {
		t.Error("List is not circular after remove")
	}
}

func TestWithStrings(t *testing.T) {
	ll := CircularLinkedList[string]{}

	strings := []string{"hello", "world", "!"}
	for _, s := range strings {
		ll.Add(s)
	}

	// Verify values and circular nature
	values := make([]string, 0)
	current := ll.head
	for current != nil && (len(values) == 0 || current != ll.head) {
		values = append(values, current.data)
		current = current.next
	}

	if len(values) != len(strings) {
		t.Errorf("Expected %d elements, got %d", len(strings), len(values))
	}

	for i, expected := range strings {
		if values[i] != expected {
			t.Errorf("Expected %v at position %d, got %v", expected, i, values[i])
		}
	}

	if current != ll.head {
		t.Error("String list is not circular")
	}
}

func TestPrint(t *testing.T) {
	ll := CircularLinkedList[int]{}
	ll.Print() // Should handle empty list

	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ll.Print() // Should handle populated list
}
