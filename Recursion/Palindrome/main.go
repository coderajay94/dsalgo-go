package main

import (
	"fmt"
	"math"
)

func main() {
	target := 12421
	fmt.Println("lets check palindome number", target)
	fmt.Println("lets check palindome number", palindome(target))

}

func palindome(n int) bool {
	countdigits := int(math.Log10(math.Abs(float64(n))) + 1)
	return n == helper(n, countdigits)
}

func helper(n int, digits int) int {

	if n == 0 {
		return 0
	}

	last := n % 10
	digits = digits - 1
	return last*int(math.Pow10(digits)) + helper(n/10, digits)
}
