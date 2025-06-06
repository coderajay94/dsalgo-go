package main

import "fmt"

func main() {
	target := 10201000
	fmt.Println("counting zeroes for ", target)
	fmt.Println("counting zeroes for ", Count(target))
}

func Count(n int) int {

	return countZeros(n)

}

func countZeros(n int) int {
	fmt.Println("calling with ", n)
	sum := 0
	if n == 0 {
		return 0
	}
	if n%10 == 0 {
		sum = 1
	} else {
		sum = 0
	}

	return sum + countZeros(n/10)
}
