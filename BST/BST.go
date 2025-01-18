package bst

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BST struct {
	root *Node
}

func (b *BST) Insert(value int) {
	newNode := &Node{Value: value}

	if b.root == nil {
		b.root = newNode
		return
	}

	current := b.root
	for {
		if value < current.Value {
			if current.Left == nil {
				current.Left = newNode
				return
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = newNode
				return
			}
			current = current.Right
		}
	}
}

func (b *BST) Search(value int) bool {
	current := b.root
	for current != nil {
		if current.Value == value {
			return true
		}
		if value < current.Value {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	return false
}

func (b *BST) Delete(value int) {
	b.root = deleteNode(b.root, value)
}

func deleteNode(root *Node, value int) *Node {
	if root == nil {
		return nil
	}

	if value < root.Value {
		root.Left = deleteNode(root.Left, value)
	} else if value > root.Value {
		root.Right = deleteNode(root.Right, value)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		minValue := findMin(root.Right)
		root.Value = minValue
		root.Right = deleteNode(root.Right, minValue)
	}
	return root
}

func findMin(node *Node) int {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current.Value
}

func (b *BST) InOrder() {
	inOrderTraversal(b.root)
	fmt.Println()
}

func inOrderTraversal(node *Node) {
	if node != nil {
		inOrderTraversal(node.Left)
		fmt.Print(node.Value, " ")
		inOrderTraversal(node.Right)
	}
}

func (b *BST) PreOrder() {
	preOrderTraversal(b.root)
	fmt.Println()
}

func preOrderTraversal(node *Node) {
	if node != nil {
		fmt.Print(node.Value, " ")
		preOrderTraversal(node.Left)
		preOrderTraversal(node.Right)
	}
}

func (b *BST) PostOrder() {
	postOrderTraversal(b.root)
	fmt.Println()
}

func postOrderTraversal(node *Node) {
	if node != nil {
		postOrderTraversal(node.Left)
		postOrderTraversal(node.Right)
		fmt.Print(node.Value, " ")
	}
}

func (b *BST) LevelOrder() {
	if b.root == nil {
		return
	}

	queue := []*Node{b.root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Print(node.Value, " ")

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	fmt.Println()
}

func (b *BST) Print() {
	fmt.Println("In-Order:")
	b.InOrder()
	fmt.Println("Pre-Order:")
	b.PreOrder()
	fmt.Println("Post-Order:")
	b.PostOrder()
	fmt.Println("Level-Order:")
	b.LevelOrder()
}
