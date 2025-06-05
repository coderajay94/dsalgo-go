package main

import "fmt"

func main() {
	PrintMe(6)
	fmt.Println()
	PrintMeReverse(6)
	fmt.Println(Factorial(6))
}

func Factorial(n int) int {
	if n == 1 {
		return 1
	}

	return n * Factorial(n-1)
}

func PrintMe(n int) {

	if n == 0 {
		return
	}
	fmt.Println(n)
	PrintMe(n - 1)
}

func PrintMeReverse(n int) {
	if n == 0 {
		return
	}
	PrintMeReverse(n - 1)
	fmt.Println(n)
}
