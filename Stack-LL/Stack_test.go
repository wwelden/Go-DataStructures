package stackll

import (
	"testing"
)

func TestStackCreation(t *testing.T) {
	s := Stack[int]{}
	if !s.IsEmpty() {
		t.Error("New Stack should be empty")
	}
}

func TestPushAndPeek(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		expected int
	}{
		{
			name:     "Push single value",
			values:   []int{1},
			expected: 1,
		},
		{
			name:     "Push multiple values",
			values:   []int{1, 2, 3},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Stack[int]{}
			for _, v := range tt.values {
				s.Push(v)
			}

			value, exists := s.Peek()
			if !exists {
				t.Error("Peek should return true for non-empty stack")
			}
			if value != tt.expected {
				t.Errorf("Expected top value %d, got %d", tt.expected, value)
			}
		})
	}
}

func TestPop(t *testing.T) {
	s := Stack[int]{}

	// Test pop on empty stack
	_, exists := s.Pop()
	if exists {
		t.Error("Pop on empty stack should return false")
	}

	// Push some values
	values := []int{1, 2, 3}
	for _, v := range values {
		s.Push(v)
	}

	// Test popping all values
	for i := len(values) - 1; i >= 0; i-- {
		value, exists := s.Pop()
		if !exists {
			t.Error("Pop should return true for non-empty stack")
		}
		if value != values[i] {
			t.Errorf("Expected popped value %d, got %d", values[i], value)
		}
	}

	// Verify stack is empty after popping all values
	if !s.IsEmpty() {
		t.Error("Stack should be empty after popping all values")
	}
}

func TestWithStrings(t *testing.T) {
	s := Stack[string]{}

	strings := []string{"hello", "world", "!"}
	for _, str := range strings {
		s.Push(str)
	}

	// Test peek
	value, exists := s.Peek()
	if !exists {
		t.Error("Peek should return true for non-empty stack")
	}
	if value != "!" {
		t.Errorf("Expected top value '!', got '%s'", value)
	}

	// Test pop
	for i := len(strings) - 1; i >= 0; i-- {
		value, exists := s.Pop()
		if !exists {
			t.Error("Pop should return true for non-empty stack")
		}
		if value != strings[i] {
			t.Errorf("Expected popped value '%s', got '%s'", strings[i], value)
		}
	}
}

func TestPrint(t *testing.T) {
	// Visual test to ensure Print() doesn't panic
	s := Stack[int]{}
	s.Print() // Should handle empty stack

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Print() // Should handle populated stack
}
