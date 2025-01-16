package main

import "fmt"

// Define the Node structure
type Node struct {
    data int
    next *Node
}

// Define the LinkedList structure
type LinkedList struct {
    head *Node
}

// Function to insert a new node at the beginning of the list
func (l *LinkedList) Insert(data int) {
    newNode := &Node{data: data}
    if l.head == nil {
        l.head = newNode
    } else {
        newNode.next = l.head
        l.head = newNode
    }
}

// Function to display the linked list
func (l *LinkedList) Display() {
    current := l.head
    for current != nil {
        fmt.Printf("%d ", current.data)
        current = current.next
    }
    fmt.Println()
}

func main() {
    list := LinkedList{}
    list.Insert(30)
    list.Insert(20)
    list.Insert(10)

    list.Display()
}