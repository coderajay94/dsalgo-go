package main

import (
	"fmt"
	"math"
)

func main() {
	target := -1432
	fmt.Println("reverse no", target)
	fmt.Println("reverse no", reverse(target))
	fmt.Println("reverse2 no", reverse2(target))

}

//concept
// i++ and ++1 are not same
// in recursion i++ will always return the same value
// as it provides the value and then increment, and seems like function call never ended so value
// didn't get increment
// but ++1 will increase the count then uses it
// so its safe to use it

// will only work for positive numbers
func reverse2(n int) int {
	countdigits := int(math.Log10(float64(n)) + 1)
	fmt.Println("digits", countdigits)
	return helper(n, countdigits)
}

func helper(n int, digits int) int {

	if n == 0 {
		return 0
	}

	last := n % 10
	digits = digits - 1
	fmt.Println("printing ", last*int(math.Pow10(digits)))
	return last*int(math.Pow10(digits)) + helper(n/10, digits)
}

var sum = 0

// reverse using external variable
func reverse(n int) int {

	//x = 1432
	if n == 0 {
		return 0
	}
	sum = sum*10 + (n % 10)
	fmt.Println(sum)
	reverse(n / 10)
	return sum
}
