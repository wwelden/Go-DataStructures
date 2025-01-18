package bst

import (
	"testing"
)

func TestBSTCreation(t *testing.T) {
	b := BST{}
	if b.root != nil {
		t.Error("New BST should have nil root")
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		expected []int // values in-order
	}{
		{
			name:     "Insert single value",
			values:   []int{5},
			expected: []int{5},
		},
		{
			name:     "Insert multiple values",
			values:   []int{5, 3, 7, 1, 4, 6, 8},
			expected: []int{1, 3, 4, 5, 6, 7, 8},
		},
		{
			name:     "Insert duplicate values",
			values:   []int{5, 5, 5},
			expected: []int{5, 5, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BST{}
			for _, v := range tt.values {
				b.Insert(v)
			}

			// Collect values in-order
			var result []int
			var inOrder func(*Node)
			inOrder = func(n *Node) {
				if n != nil {
					inOrder(n.Left)
					result = append(result, n.Value)
					inOrder(n.Right)
				}
			}
			inOrder(b.root)

			if len(result) != len(tt.expected) {
				t.Errorf("Expected %d nodes, got %d", len(tt.expected), len(result))
			}

			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Expected %v at position %d, got %v", tt.expected[i], i, result[i])
				}
			}
		})
	}
}

func TestSearch(t *testing.T) {
	b := BST{}
	values := []int{5, 3, 7, 1, 4, 6, 8}
	for _, v := range values {
		b.Insert(v)
	}

	tests := []struct {
		name     string
		search   int
		expected bool
	}{
		{"Search existing value", 4, true},
		{"Search root value", 5, true},
		{"Search leaf value", 8, true},
		{"Search non-existing value", 10, false},
		{"Search negative value", -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.Search(tt.search); got != tt.expected {
				t.Errorf("Search(%d) = %v, want %v", tt.search, got, tt.expected)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name           string
		initialValues  []int
		deleteValue    int
		expectedValues []int
		expectedExists bool
	}{
		{
			name:           "Delete leaf node",
			initialValues:  []int{5, 3, 7, 1, 4, 6, 8},
			deleteValue:    8,
			expectedValues: []int{1, 3, 4, 5, 6, 7},
			expectedExists: false,
		},
		{
			name:           "Delete node with one child",
			initialValues:  []int{5, 3, 7, 1},
			deleteValue:    3,
			expectedValues: []int{1, 5, 7},
			expectedExists: false,
		},
		{
			name:           "Delete node with two children",
			initialValues:  []int{5, 3, 7, 1, 4, 6, 8},
			deleteValue:    7,
			expectedValues: []int{1, 3, 4, 5, 6, 8},
			expectedExists: false,
		},
		{
			name:           "Delete root",
			initialValues:  []int{5, 3, 7},
			deleteValue:    5,
			expectedValues: []int{3, 7},
			expectedExists: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BST{}
			for _, v := range tt.initialValues {
				b.Insert(v)
			}

			b.Delete(tt.deleteValue)

			// Verify the value was deleted
			if exists := b.Search(tt.deleteValue); exists != tt.expectedExists {
				t.Errorf("After deletion, Search(%d) = %v, want %v",
					tt.deleteValue, exists, tt.expectedExists)
			}

			// Collect remaining values
			var result []int
			var inOrder func(*Node)
			inOrder = func(n *Node) {
				if n != nil {
					inOrder(n.Left)
					result = append(result, n.Value)
					inOrder(n.Right)
				}
			}
			inOrder(b.root)

			if len(result) != len(tt.expectedValues) {
				t.Errorf("Expected %d nodes after deletion, got %d",
					len(tt.expectedValues), len(result))
			}

			for i := range result {
				if result[i] != tt.expectedValues[i] {
					t.Errorf("Expected %v at position %d, got %v",
						tt.expectedValues[i], i, result[i])
				}
			}
		})
	}
}

func TestTraversals(t *testing.T) {
	// This is more of a visual test to ensure traversal methods don't panic
	b := BST{}

	// Test empty tree
	b.InOrder()
	b.PreOrder()
	b.PostOrder()
	b.LevelOrder()

	// Test with values
	values := []int{5, 3, 7, 1, 4, 6, 8}
	for _, v := range values {
		b.Insert(v)
	}

	b.InOrder()    // Should print: 1 3 4 5 6 7 8
	b.PreOrder()   // Should print: 5 3 1 4 7 6 8
	b.PostOrder()  // Should print: 1 4 3 6 8 7 5
	b.LevelOrder() // Should print: 5 3 7 1 4 6 8
}
