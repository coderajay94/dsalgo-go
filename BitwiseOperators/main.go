package main

import "fmt"

func main() {

	//whenever you and anything with , it will give the same result
	// 111001 & 111111 = result in same number as 111001
	//so to check even odd we just need to and it with one and it will give us the last digit in binary
	//that is responsible for even or odd
	fmt.Println("testing bitwise operator")

	//this last bit is called lsb - least significant bit
	fmt.Println(isOdd(12))
	fmt.Println(isOdd(1))
	fmt.Println(isOdd(0))
}

func isOdd(n int) bool {
	return n&1 == 1
}
