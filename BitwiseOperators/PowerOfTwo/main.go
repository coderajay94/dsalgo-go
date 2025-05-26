package main

import "fmt"

func main() {
	fmt.Println("how to check if no is power of two")

	// lets see 1000
	// if any no is power of 2, only right most would be the one
	// so and it with n-1
	// 1000 &
	// 0111
	// 0000 if it becomes zero that means its power of 2

	// 1000000
	// 0000001
	// 1111110
	// 1111111

	fmt.Println(powerOfTwo(12))
	fmt.Println(powerOfTwo(4))
	fmt.Println(powerOfTwo(16))
	fmt.Println(powerOfTwo(-1))
	fmt.Println(powerOfTwo(0))
}

func powerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	return (n & (n - 1)) == 0
}
