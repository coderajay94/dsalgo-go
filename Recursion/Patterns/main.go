package main

import (
	"fmt"
)

func main() {

	fmt.Println("print pattern")
	pattern1(4, 0)
}

func pattern1(r, c int) {
	if r <= 0 {
		return
	}

	if r > c {
		fmt.Print("* ")
		pattern1(r, c+1)
	} else {
		fmt.Println()
		pattern1(r-1, 0)
	}
}
