package main

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	Head *Node
	Tail *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 {
		return -1
	}

	node := this.Head
	for i := 0; i < index && node != nil; i++ {
		node = node.Next
	}

	if node == nil {
		return -1
	}

	return node.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &Node{Val: val, Next: this.Head}

	if this.Tail == nil {
		this.Tail = newNode
	}

	// newNode := &Node{Val:val, Next:this.Head}
	this.Head = newNode
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &Node{Val: val}

	if this.Head == nil {
		this.Head = newNode
		this.Tail = newNode
		return
	}

	this.Tail.Next = newNode
	this.Tail = newNode

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {

	if index < 0 {
		return
	}

	newNode := &Node{Val: val}
	if index == 0 {
		newNode.Next = this.Head
		this.Head = newNode
		if this.Tail == nil {
			this.Tail = newNode
		}
		return
	} else {
		node := this.Head
		for i := 0; i < index-1 && node != nil; i++ {
			node = node.Next
		}

		if node == nil {
			return
		}
		newNode.Next = node.Next
		node.Next = newNode

		if node == this.Tail {
			this.Tail = newNode
		}
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || this.Head == nil {
		return
	}

	if index == 0 {
		this.Head = this.Head.Next
		if this.Head == nil {
			this.Tail = nil
		}
		return
	} else {
		node := this.Head
		for i := 0; i < index-1; i++ {
			if node == nil || node.Next == nil {
				return
			}
			node = node.Next
		}

		//nil check
		if node.Next == nil {
			return
		}

		if node.Next == this.Tail {
			this.Tail = node
		}
		node.Next = node.Next.Next
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func main() {

}
