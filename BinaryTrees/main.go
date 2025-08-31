package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// takes only thing before space
	fmt.Println("welcome")
	// var name string
	// fmt.Println("Enter the root node:")
	// fmt.Scanln(&name)

	// fmt.Println("welcome ", name)

	reader := bufio.NewReader(os.Stdin)

	//fmt.Println("enter your name:")
	//name, _ := reader.ReadString('\n')

	//fmt.Println("welcome ", name)

	fmt.Println("Wanna proceeed next:")
	isTrue, _ := reader.ReadString('\n')
	isTrue = strings.TrimSpace(isTrue)

	val, err := strconv.ParseBool(isTrue)
	if err != nil {
		fmt.Println("Invalid Input, please enter true or false", val, err)
	}
	fmt.Println("entered bool is", val)
}
