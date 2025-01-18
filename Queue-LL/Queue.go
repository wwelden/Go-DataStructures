package queuell

import "fmt"

type Node[T comparable] struct {
	data T
	next *Node[T]
}

type Queue[T comparable] struct {
	head *Node[T]
	tail *Node[T]
}

func (q *Queue[T]) IsEmpty() bool {
	return q.head == nil
}

func (q *Queue[T]) Enqueue(data T) {
	newNode := &Node[T]{data: data, next: nil}

	if q.IsEmpty() {
		q.head = newNode
		q.tail = newNode
		return
	}

	q.tail.next = newNode
	q.tail = newNode
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}

	data := q.head.data
	q.head = q.head.next

	if q.head == nil {
		q.tail = nil
	}

	return data, true
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	return q.head.data, true
}

func (q *Queue[T]) Print() {
	current := q.head
	for current != nil {
		fmt.Printf("%v ", current.data)
		current = current.next
	}
	fmt.Println()
}
