package queuearr

import "fmt"

type Queue[T comparable] struct {
	q []T
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.q) == 0
}

func (q *Queue[T]) Length() int {
	return len(q.q)
}

func (q *Queue[T]) Print() {
	length := len(q.q)
	for i := 0; i < length; i++ {
		fmt.Print(q.q[i], " ")
	}
	fmt.Println()
}

func (q *Queue[T]) Enqueue(value T) {
	q.q = append(q.q, value)
}
func (q *Queue[T]) Dequeue() T {
	res := q.q[0]
	q.q = q.q[1:]
	return res
}

func (q *Queue[T]) Front() T {
	return q.q[0]
}

func (q *Queue[T]) Back() T {
	return q.q[len(q.q)-1]
}
