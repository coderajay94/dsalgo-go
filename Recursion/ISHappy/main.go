package main

import "fmt"

//logic is when you want to break the loop for infinite
//lopok for no when they start repeating in a map

func main() {
	//this no will keep on looping forever
	// to avoid that keep the processed records in map
	// if no repeats that means its in forever loop or cycle then break it
	fmt.Println(isHappy(78))
	fmt.Println(isHappy(19))
}

func isHappy(n int) bool {
	m := make(map[int]bool)
	for n != 1 && !m[n] {
		m[n] = true
		n = squareOfSum(n)
	}

	return n == 1
}

func squareOfSum(n int) int {
	if n == 0 {
		return 0
	}
	mod := n % 10
	return mod*mod + squareOfSum(n/10)
}
