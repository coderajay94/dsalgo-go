package main

import (
	"fmt"
)

type Node struct {
	Val  interface{}
	Next *Node
}

type CircularLinkedList struct {
	Head *Node
	Tail *Node
}

func NewCircularLinkedList() *CircularLinkedList {
	return &CircularLinkedList{}
}

func (list *CircularLinkedList) Insert(data interface{}) {
	node := &Node{Val: data}
	if list.Head == nil {
		list.Head = node
		list.Tail = node
		node.Next = node
	} else {
		node.Next = list.Head
		list.Tail.Next = node
		list.Tail = node
	}
}

func (list CircularLinkedList) Print() {
	if list.Head == nil {
		fmt.Println("Empty Circular LinkedList")
		return
	}
	node := list.Head

	for {
		fmt.Print(node.Val, " -> ")
		node = node.Next
		if node == list.Head {
			break
		}
	}
	fmt.Println("End")
}

func main() {
	fmt.Println("circular linked list")

	list := NewCircularLinkedList()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Print()
	list.InsertAtIndex(4, 100)
}

func (list *CircularLinkedList) InsertAtIndex(index int, data interface{}) {
	insertAt(data, list, index, list.Head)
	list.Print()
}

func insertAt(data interface{}, list *CircularLinkedList, index int, node *Node) {
	fmt.Println(index, data, node.Val)
	if index == 0 {
		newNode := &Node{Val: data, Next: node.Next}
		node.Next = newNode
		return
	} else {
		index = index - 1
		node = node.Next
		insertAt(data, list, index, node.Next)
	}

}
