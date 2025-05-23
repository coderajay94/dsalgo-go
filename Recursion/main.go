package main

import (
	"fmt"
)

func main() {

	PrintMe(1)
}

func PrintMe(n int) {
	fmt.Println("printing ", n)
	PrintMe(n + 1)
}
