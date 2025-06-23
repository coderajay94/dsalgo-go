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

func main() {

	ll := LinkedList{}
	ll.AddLast(1)
	ll.AddLast(3)
	ll.AddLast(4)
	ll.AddLast(5)
	ll.AddLast(6)
	ll.PrintLinkedList()
}

func hasCycle(head *Node) bool {

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}

//count cycle length
func hasCycleCount(head *Node) int {

	slow := head
	fast := head
	counter := 0

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			counter = 1
			slow = slow.Next
			for slow != fast {
				counter++
				slow = slow.Next
			}
			return counter
		}
	}
	return counter
}
