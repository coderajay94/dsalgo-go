package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println(lowestSetBit(12))

	// calculate total set bits
	fmt.Println(countSetBits(12))
	fmt.Println(countSetBits(13))
	fmt.Println(countSetBits(0))
	// will only check for positives
	fmt.Println(countSetBits(-8))

	binary := strconv.FormatInt(int64(1244323), 2)
	//100101111110010100011
	fmt.Printf("Binary of %d is %s\n", 1244323, binary)
	fmt.Println(countSetBits(1244323))

}

func countSetBits(n int) int {
	count := 0
	for n > 0 {
		// if it returns 4, we need to subtract that from the number to cal remaining set bits
		n = n - (n & -n)
		count++
	}
	return count
}

func lowestSetBit(n int) int {

	//how to calculate the lowest set bit
	// and n and -n which will give you lowest set bit
	// 12 = 1100
	// -12 = 0011 + 1 =
	// 0100 = 4 - which is the lowest set bit

	return n & (-n)
}
