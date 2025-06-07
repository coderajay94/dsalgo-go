package main

import (
	"fmt"
)

func main() {

	fmt.Println("print pattern")
	//pattern1(4, 0)
	pattern2(4, 0)
}

// triange
func pattern2(r, c int) {
	if r <= 0 {
		return
	}

	if r > c {
		pattern2(r, c+1)
		fmt.Print("* ")
	} else {
		pattern2(r-1, 0)
		fmt.Println()
	}
}

// reverse triangle
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
