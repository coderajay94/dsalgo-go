package main

import "fmt"

func main() {

	// range xor from a to b; xor(b) ^ xor(a-1)
	a := 3
	b := 9

	ans := xor(b) ^ xor(a-1)
	fmt.Println(ans)
}

func xor(n int) int {

	//pattern observed when n xored
	if n%4 == 0 {
		return n
	} else if n%4 == 1 {
		return 1
	} else if n%4 == 2 {
		return n + 1
	} else {
		// n%4 == 3
		return 0
	}

}
