package queuell

import (
	"testing"
)

func TestQueueCreation(t *testing.T) {
	q := Queue[int]{}
	if !q.IsEmpty() {
		t.Error("New Queue should be empty")
	}
}

func TestEnqueueAndPeek(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		expected int
	}{
		{
			name:     "Enqueue single value",
			values:   []int{1},
			expected: 1,
		},
		{
			name:     "Enqueue multiple values",
			values:   []int{1, 2, 3},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := Queue[int]{}
			for _, v := range tt.values {
				q.Enqueue(v)
			}

			value, exists := q.Peek()
			if !exists {
				t.Error("Peek should return true for non-empty queue")
			}
			if value != tt.expected {
				t.Errorf("Expected front value %d, got %d", tt.expected, value)
			}
		})
	}
}

func TestDequeue(t *testing.T) {
	q := Queue[int]{}

	// Test dequeue on empty queue
	_, exists := q.Dequeue()
	if exists {
		t.Error("Dequeue should return false on empty queue")
	}

	// Test dequeue with values
	values := []int{1, 2, 3}
	for _, v := range values {
		q.Enqueue(v)
	}

	// Test dequeuing all values
	for i, expected := range values {
		value, exists := q.Dequeue()
		if !exists {
			t.Errorf("Dequeue should return true for value at index %d", i)
		}
		if value != expected {
			t.Errorf("Expected dequeued value %d, got %d", expected, value)
		}
	}

	// Verify queue is empty after dequeuing all values
	if !q.IsEmpty() {
		t.Error("Queue should be empty after dequeuing all values")
	}
}

func TestWithStrings(t *testing.T) {
	q := Queue[string]{}

	strings := []string{"hello", "world", "!"}
	for _, s := range strings {
		q.Enqueue(s)
	}

	// Test FIFO order with strings
	for _, expected := range strings {
		value, exists := q.Dequeue()
		if !exists {
			t.Error("Dequeue should return true for non-empty queue")
		}
		if value != expected {
			t.Errorf("Expected dequeued value '%s', got '%s'", expected, value)
		}
	}
}

func TestPrint(t *testing.T) {
	// Visual test to ensure Print() doesn't panic
	q := Queue[int]{}
	q.Print() // Should handle empty queue

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Print() // Should handle populated queue
}
