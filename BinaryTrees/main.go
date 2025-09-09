package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome")
	root := Populate()
	Display(root)
}

type Node struct {
	Value int64
	Left  *Node
	Right *Node
}

func Display(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(" ", root.Value)

	if root.Left != nil {
		Display(root.Left)
	}

	if root.Right != nil {
		Display(root.Right)
	}
}

func Populate() *Node {

	fmt.Println("Enter the Root Node")
	reader := bufio.NewReader(os.Stdin)

	rootvalue, _ := reader.ReadString('\n')
	rootvalue = strings.TrimSpace(rootvalue)

	val, err := strconv.ParseInt(rootvalue, 10, 64)

	if err != nil {
		fmt.Println("err : Please enter a valid root value in integer format")
	}

	node := Node{Value: val}
	populate(&node, reader)
	return &node
}

func populate(node *Node, reader *bufio.Reader) {

	fmt.Println("Do you want to enter left of ", node.Value)
	inputValueLeft, _ := reader.ReadString('\n')
	inputValueLeft = strings.TrimSpace(inputValueLeft)
	inputLeft, err := strconv.ParseBool(inputValueLeft)
	if err != nil {
		fmt.Println("err : enter a valid boolean value for left")
	}

	if inputLeft {
		fmt.Println("Enter the value of left of ", node.Value)
		stringvalue, _ := reader.ReadString('\n')
		stringvalue = strings.TrimSpace(stringvalue)
		intvalue, err := strconv.ParseInt(stringvalue, 10, 64)
		if err != nil {
			fmt.Println("err : enter a valid int value for left")
		}
		node.Left = &Node{Value: intvalue}
		populate(node.Left, reader)
	}

	fmt.Println("Do you want to enter Right of ", node.Value)
	inputValueRight, _ := reader.ReadString('\n')
	inputValueRight = strings.TrimSpace(inputValueRight)
	inputRight, err := strconv.ParseBool(inputValueRight)
	if err != nil {
		fmt.Println("err : enter a valid boolean value for left")
	}

	if inputRight {
		fmt.Println("Enter the value of Right of ", node.Value)
		stringvalue, _ := reader.ReadString('\n')
		stringvalue = strings.TrimSpace(stringvalue)
		intvalue, err := strconv.ParseInt(stringvalue, 10, 64)
		if err != nil {
			fmt.Println("err : enter a valid int value for Right")
		}
		node.Right = &Node{Value: intvalue}
		populate(node.Right, reader)
	}
	fmt.Println("coming out of stack for node ", node.Value)
}
