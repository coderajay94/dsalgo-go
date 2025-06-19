package main

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (list *LinkedList) AddLast(data interface{}) {
	node := Node{Data: data, Next: nil}
	if list == nil || list.Head == nil {
		list.Head = &node
		list.Tail = &node
		list.Size++
	} else {
		list.Tail.Next = &node
		list.Tail = &node
		list.Size++
	}
}

func (list LinkedList) PrintLinkedList() {
	node := list.Head
	for list.Size > 0 {
		fmt.Print(node.Data, " -> ")
		node = node.Next
		list.Size--
	}
	fmt.Println(" end")
}

func (list *LinkedList) InsertAtPosition(position int, data interface{}) {
	node := Node{Data: data}

	if position == 1 {
		node.Next = list.Head
		list.Head = &node
		list.Size++
		return
	}

	previous := list.Head
	counter := 2
	for previous.Next != nil {
		if counter == position {
			node.Next = previous.Next
			previous.Next = &node
			list.Size++
			return
		} else {
			counter++
			previous = previous.Next
		}
	}
	previous.Next = &node
	node.Next = previous
	list.Size++
	return

}

func (list *LinkedList) InsertAtIndex(index int, data interface{}) {

	if index > list.Size {
		fmt.Println("Index not found")
		return
	}
	//case 1 := starting element
	if index == 0 {
		node := Node{Data: data}
		node.Next = list.Head
		list.Head = &node
		return
	}
	//case 2 := last element
	if index == list.Size {
		list.AddLast(data)
		return
	}
	pointer := list.Head
	for i := 1; i < index; i++ {
		pointer = pointer.Next
	}

	node := Node{Data: data, Next: pointer.Next}
	pointer.Next = &node
	list.Size++
}

func (list *LinkedList) DeleteFirst() {
	if list.Size == 0 {
		fmt.Println("delete can't be performed")
	} else if list.Size == 1 {
		list.Head = nil
		list.Tail = nil
		list.Size--
	} else {
		list.Head = list.Head.Next
		list.Size--
	}
}

func (list *LinkedList) DeleteLast() {
	if list.Size == 0 {
		fmt.Println("Can't delete from an empty list")
	} else if list.Size == 1 {
		list.Head = nil
		list.Tail = nil
		list.Size--
	} else {
		pointer := list.Head
		for pointer.Next != nil {
			pointer = pointer.Next
		}
		pointer.Next = nil
		list.Size--
	}
}

func main() {
	fmt.Println("linkedlist with head and tail")

	list := LinkedList{}
	list.AddLast(12)
	list.PrintLinkedList()

	list.AddLast(121)
	list.PrintLinkedList()

	list.AddLast(112)
	list.PrintLinkedList()

	list.AddLast(102)

	list.PrintLinkedList()
	fmt.Println("------------------------")
	// list.InsertAtPosition(6, 100)
	// list.PrintLinkedList()
	fmt.Println("------------------------")

	// list.InsertAtIndex(4, 100)
	// list.PrintLinkedList()

	fmt.Println("------------------------")

	list.DeleteFirst()
	list.PrintLinkedList()
	fmt.Println("------------------------")

	list.DeleteFirst()
	list.PrintLinkedList()
	fmt.Println("------------------------")

	list.DeleteLast()
	list.PrintLinkedList()
	fmt.Println("------------------------")

	list.DeleteLast()
	list.PrintLinkedList()
	fmt.Println("******")

	list.DeleteLast()
	list.PrintLinkedList()
}
