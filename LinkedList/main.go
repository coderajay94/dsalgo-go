package main

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

// mostly pass by reference, not by value
func (l *LinkedList) prepend(newNode *Node) {
	//step 1
	oldNode := l.Head
	//step2
	l.Head = newNode
	//step 3
	l.Head.Next = oldNode

	l.Length++
}

//this time we only need value reciever, as we're not going to modify the list
//just printing the node
func (list LinkedList) PrintLinkedList() {
	pointer := list.Head

	for list.Length != 0 {
		fmt.Print(pointer.Data, " -> ")
		pointer = pointer.Next
		//not changing actual length, as its passed as a value, so only changed in that copy
		//not in the actual list
		list.Length--
	}
	fmt.Println("End")
}

func main() {
	list := LinkedList{}

	//as we're passing the actual reference, so add & to send the actual address of that Node
	node1 := &Node{Data: 8}
	node2 := &Node{Data: 18}
	node3 := &Node{Data: 81}
	node4 := &Node{Data: 108}
	list.prepend(node1)
	list.prepend(node2)
	list.prepend(node3)
	list.prepend(node4)

	//just prints the address and length of that list
	fmt.Println(list)

	list.PrintLinkedList()
}
