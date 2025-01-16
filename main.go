package main

import (
	"fmt"

	LinkedList "github.com/wwelden/Go-DataStructures/Linked-List"
)

func main() {

	fmt.Println("Linked List")
	ll := LinkedList.LinkedList{}
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ll.Print()

}
