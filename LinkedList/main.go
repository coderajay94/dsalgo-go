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

//using actual reference as going to delete

func (list *LinkedList) DeleteDataNode(data interface{}) {

	//list can also be empty - of called like below
	//var list *LinkedList = nil
	//	list.DeleteDataNode(10) // ✅ This will NOT panic — but `list == nil`

	//case 1 has 2 parts : when trying to delete from empty list
	if list == nil || list.Head == nil {
		fmt.Println("Trying to delete, ", data, " from an empty LinkedList")
		return
	}
	previousNode := list.Head

	//case 2 : when head node data is deleted
	if data == previousNode.Data {
		list.Head = previousNode.Next
		list.Length--
		return
	}

	for previousNode.Next.Data != data {
		if previousNode.Next.Next == nil {
			fmt.Println(data, " not found")
			return
		}
		// to move pointed to correct position
		previousNode = previousNode.Next
	}

	previousNode.Next = previousNode.Next.Next
	list.Length--
	return
}

func (list LinkedList) CheckDataNode(data interface{}) {
	node := list.Head
	counter := 0
	for node.Next != nil {
		counter++
		if node.Data == data {
			fmt.Println(data, ": data found at position:", counter)
			return
		}
		node = node.Next
	}

}

func main() {
	list := LinkedList{}

	//as we're passing the actual reference, so add & to send the actual address of that Node
	node1 := &Node{Data: 8}
	node2 := &Node{Data: 18}
	node3 := &Node{Data: 81}
	node4 := &Node{Data: 108}
	node5 := &Node{Data: 28}

	list.prepend(node1)
	list.prepend(node2)
	list.prepend(node3)
	list.prepend(node4)
	list.prepend(node5)

	//just prints the address and length of that list
	//fmt.Println(list)

	list.PrintLinkedList()

	//list.CheckDataNode(81)
	// list.DeleteDataNode(81)
	// list.PrintLinkedList()

	//case 2 : when head node data is deleted
	list.DeleteDataNode(28)
	list.PrintLinkedList()

	//case 1 : when empty list, deleteing data
	// list2 := LinkedList{}
	// list2.DeleteDataNode(100)

}
