package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("lets count the digits of binary number")

	b := 0b11111011

	noOfBits := int(math.Log2(float64(b))) + 1
	fmt.Println(noOfBits)
}
