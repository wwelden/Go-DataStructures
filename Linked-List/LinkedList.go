package LinkedList

import "fmt"

type Node[T comparable] struct {
	data T
	next *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
}

func (l *LinkedList[T]) Add(data T) {
	newNode := &Node[T]{data: data, next: nil}
	if l.head.next == nil {
		l.head = newNode
		return
	}

	curNode := l.head
	for curNode != nil {
		curNode = curNode.next
	}
	curNode.next = newNode
}

func (l *LinkedList[T]) Remove() {
	curNode := l.head

	for curNode.next.next != nil {
		curNode.next = curNode
	}
	curNode.next = nil
}

func (l *LinkedList[T]) AddFront(data T) {

}

func (l *LinkedList[T]) Print() {
	current := l.head
	for current != nil {
		fmt.Printf("%v ", current.data)
		current = current.next
	}
	fmt.Println()
}
