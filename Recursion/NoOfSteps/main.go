package main

import (
	"fmt"
)

func main() {
	fmt.Println("count no of steps", callHelper(14, 0))
}

func callHelper(num int, count int) int {
	fmt.Println(num, count)
	if num == 0 {
		return count
	}

	if num&1 == 0 {
		return callHelper(num/2, count+1)
	} else {
		return callHelper(num-1, count+1)
	}
}
