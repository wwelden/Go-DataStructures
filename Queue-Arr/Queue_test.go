package queuearr

import (
	"testing"
)

func TestQueueCreation(t *testing.T) {
	q := Queue[int]{}
	if !q.IsEmpty() {
		t.Error("New Queue should be empty")
	}
	if q.Length() != 0 {
		t.Error("New Queue should have length 0")
	}
}

func TestEnqueueDequeue(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		expected []int
	}{
		{
			name:     "Single value",
			values:   []int{1},
			expected: []int{1},
		},
		{
			name:     "Multiple values",
			values:   []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Empty queue",
			values:   []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := Queue[int]{}

			// Test Enqueue
			for _, v := range tt.values {
				q.Enqueue(v)
			}

			if q.Length() != len(tt.values) {
				t.Errorf("Expected length %d, got %d", len(tt.values), q.Length())
			}

			// Test Dequeue
			for i, expected := range tt.expected {
				if q.IsEmpty() {
					t.Errorf("Queue is empty before dequeuing all values at index %d", i)
					return
				}
				got := q.Dequeue()
				if got != expected {
					t.Errorf("Expected %v, got %v", expected, got)
				}
			}
		})
	}
}

func TestFrontBack(t *testing.T) {
	q := Queue[int]{}

	// Test with single element
	q.Enqueue(1)
	if q.Front() != 1 || q.Back() != 1 {
		t.Error("Front and Back should both be 1 for single element queue")
	}

	// Test with multiple elements
	q.Enqueue(2)
	q.Enqueue(3)
	if q.Front() != 1 {
		t.Errorf("Expected front to be 1, got %v", q.Front())
	}
	if q.Back() != 3 {
		t.Errorf("Expected back to be 3, got %v", q.Back())
	}
}

func TestWithStrings(t *testing.T) {
	q := Queue[string]{}

	strings := []string{"hello", "world", "!"}
	for _, s := range strings {
		q.Enqueue(s)
	}

	for i, expected := range strings {
		if q.IsEmpty() {
			t.Errorf("Queue is empty before dequeuing all values at index %d", i)
			return
		}
		got := q.Dequeue()
		if got != expected {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	}
}

func TestPrint(t *testing.T) {
	// Visual test to ensure it doesn't panic
	q := Queue[int]{}
	q.Print() // Should handle empty queue

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Print() // Should handle populated queue
}
