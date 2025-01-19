package CircularLinkedList

import "fmt"

type Node[T comparable] struct {
	data T
	next *Node[T]
}

type CircularLinkedList[T comparable] struct {
	head *Node[T]
}

func (l *CircularLinkedList[T]) Add(data T) {
	newNode := &Node[T]{data: data, next: nil}
	if l.head == nil {
		l.head = newNode
		newNode.next = l.head // Point to itself to make it circular
		return
	}

	curNode := l.head
	for curNode.next != l.head {
		curNode = curNode.next
	}
	curNode.next = newNode
	newNode.next = l.head // Make it circular by pointing to head
}

func (l *CircularLinkedList[T]) Remove() {
	if l.head == nil {
		return
	}
	if l.head.next == l.head {
		l.head = nil // Only one node in the list
		return
	}

	curNode := l.head
	for curNode.next.next != l.head {
		curNode = curNode.next
	}
	curNode.next = l.head // Point to head to maintain circular nature
}

func (l *CircularLinkedList[T]) AddFront(data T) {
	newNode := &Node[T]{data: data}
	if l.head == nil {
		l.head = newNode
		newNode.next = l.head // Point to itself to make it circular
		return
	}

	curNode := l.head
	for curNode.next != l.head {
		curNode = curNode.next
	}
	newNode.next = l.head
	l.head = newNode
	curNode.next = l.head // Update the last node to point to new head
}

func (l *CircularLinkedList[T]) Print() {
	if l.head == nil {
		fmt.Println()
		return
	}

	current := l.head
	for {
		fmt.Printf("%v ", current.data)
		current = current.next
		if current == l.head {
			break
		}
	}
	fmt.Println()
}
