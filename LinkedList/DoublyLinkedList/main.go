package main

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
	Prev *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (list *DoublyLinkedList) InsertAtFirst(data interface{}) {

	// 1 <-> 2 <-> 3
	node := &Node{Data: data}
	if list.Size == 0 {
		list.Head = node
		list.Tail = node
		list.Size++
	} else {
		list.Head.Prev = node
		node.Next = list.Head
		list.Head = node
		list.Size++
	}
}

func (list DoublyLinkedList) PrintDoublyLinkedList() {

	if list.Head == nil {
		fmt.Println("Empty List")
		return
	}

	node := list.Head
	for node != nil {
		fmt.Print(node.Data, " -> ")
		node = node.Next
	}
	fmt.Println("End")
}

func (list *DoublyLinkedList) insertAtLast(data interface{}) {
	node := &Node{Data: data}
	if list.Size == 0 {
		list.Head = node
		list.Tail = node
		list.Size++
	} else {
		node.Prev = list.Tail
		list.Tail.Next = node
		list.Tail = node
		list.Size++
	}
}

func main() {

	list := NewDoublyLinkedList()

	list.InsertAtFirst(12)
	list.InsertAtFirst(14)
	list.InsertAtFirst(16)
	list.InsertAtFirst(19)
	list.InsertAtFirst(1)
	list.insertAtLast(100)
	list.insertAtLast(200)

	list.PrintDoublyLinkedList()
}
