package main

import "fmt"

func main() {
	fmt.Println("check palindrome")
	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {

	if x < 0 {
		//added
		return false
	}
	temp := 0
	for x > 0 {
		last := x % 10
		temp = temp*10 + last
		x = x / 10
	}
	fmt.Println(x)
	if x == temp {
		return true
	}
	return false
}
