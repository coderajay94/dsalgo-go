package main

import "fmt"

func main() {
	PrintMe(6)
}

func PrintMe(n int) {
	fmt.Println(n)
	if n == 1 {
		return
	}
	PrintMe(n - 1)
}
