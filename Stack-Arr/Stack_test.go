package stackarr

import (
	"testing"
)

func TestStackCreation(t *testing.T) {
	s := Stack[int]{}
	if !s.IsEmpty() {
		t.Error("New Stack should be empty")
	}
	if s.Length() != 0 {
		t.Error("New Stack should have length 0")
	}
}

func TestPushAndTop(t *testing.T) {
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

			if s.Length() != len(tt.values) {
				t.Errorf("Expected length %d, got %d", len(tt.values), s.Length())
			}

			if s.Peek() != tt.expected {
				t.Errorf("Expected top value %d, got %d", tt.expected, s.Peek())
			}
		})
	}
}

func TestPop(t *testing.T) {
	s := Stack[int]{}

	// Push some values
	values := []int{1, 2, 3}
	for _, v := range values {
		s.Push(v)
	}

	// Test popping all values
	for i := len(values) - 1; i >= 0; i-- {
		popped := s.Pop()
		if popped != values[i] {
			t.Errorf("Expected popped value %d, got %d", values[i], popped)
		}
	}

	if !s.IsEmpty() {
		t.Error("Stack should be empty after popping all values")
	}
}

func TestBottomAndTop(t *testing.T) {
	s := Stack[int]{}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Bottom() != 1 {
		t.Errorf("Expected bottom value 1, got %d", s.Bottom())
	}

	if s.Peek() != 3 {
		t.Errorf("Expected top value 3, got %d", s.Peek())
	}
}

func TestWithStrings(t *testing.T) {
	s := Stack[string]{}

	strings := []string{"hello", "world", "!"}
	for _, str := range strings {
		s.Push(str)
	}

	if s.Length() != len(strings) {
		t.Errorf("Expected length %d, got %d", len(strings), s.Length())
	}

	if s.Peek() != "!" {
		t.Errorf("Expected top value '!', got '%s'", s.Peek())
	}

	if s.Bottom() != "hello" {
		t.Errorf("Expected bottom value 'hello', got '%s'", s.Bottom())
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
