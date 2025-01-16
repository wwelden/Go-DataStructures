package LinkedList

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Add(data int) {
	newNode := &Node{data: data, next: nil}
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

func (l *LinkedList) Remove() {
	curNode := l.head

	for curNode.next.next != nil {
		curNode.next = curNode
	}
	curNode.next = nil
}

func (l *LinkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}
