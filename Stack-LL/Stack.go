package stackll

import "fmt"

type Node[T comparable] struct {
	data T
	next *Node[T]
}

type Stack[T comparable] struct {
	head *Node[T]
}

func (s *Stack[T]) IsEmpty() bool {
	return s.head == nil
}

func (s *Stack[T]) Push(value T) {
	newNode := &Node[T]{
		data: value,
		next: s.head,
	}
	s.head = newNode
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	value := s.head.data
	s.head = s.head.next
	return value, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	return s.head.data, true
}

func (s *Stack[T]) Print() {
	current := s.head
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
}
