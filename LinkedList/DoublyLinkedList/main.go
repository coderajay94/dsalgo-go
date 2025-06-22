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

func (list DoublyLinkedList) PrintReverse() {
	if list.Tail == nil {
		fmt.Println("Empty List")
		return
	}

	node := list.Tail

	for node != nil {
		fmt.Print(node.Data, " -> ")
		node = node.Prev
	}
	fmt.Println("Start")

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

func (list *DoublyLinkedList) DeleteAtFirst() {
	if list.Size == 0 {
		return
	} else if list.Size == 1 {
		list.Head = nil
		list.Tail = nil
	} else {
		list.Head = list.Head.Next
		list.Head.Prev = nil
		list.Size--
	}
}

func (list *DoublyLinkedList) DeleteAtLast() {
	if list.Size == 0 {
		return
	} else if list.Size == 1 {
		list.Head = nil
		list.Tail = nil
	} else {
		list.Tail = list.Tail.Prev
		list.Tail.Next = nil
		list.Size--
	}

}

func (list *DoublyLinkedList) DeleteValue(value interface{}) {

}

func main() {

	//empty := NewDoublyLinkedList()
	onlyOne := NewDoublyLinkedList()
	onlyOne.InsertAtFirst(11)

	list := NewDoublyLinkedList()

	list.InsertAtFirst(12)
	list.InsertAtFirst(14)
	list.InsertAtFirst(16)
	list.InsertAtFirst(19)
	list.InsertAtFirst(1)
	list.insertAtLast(100)
	list.insertAtLast(200)

	// 	list.PrintDoublyLinkedList()
	// 	list.PrintReverse()

	// 	empty.PrintReverse()
	// 	onlyOne.PrintReverse()

	// 	fmt.Println("------------------")
	// 	list.DeleteAtLast()
	// 	list.PrintDoublyLinkedList()
	// 	list.DeleteAtLast()
	// 	list.PrintDoublyLinkedList()
	// 	fmt.Println("------------------")
	// 	onlyOne.PrintDoublyLinkedList()
	// 	onlyOne.DeleteAtFirst()
	// 	onlyOne.PrintDoublyLinkedList()
	// 	onlyOne.DeleteAtFirst()
	// 	onlyOne.PrintDoublyLinkedList()
	fmt.Println("------------------")
	list.PrintDoublyLinkedList()
	fmt.Println("------------------")
}
