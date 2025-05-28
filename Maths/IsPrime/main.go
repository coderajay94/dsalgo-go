package main

import "fmt"

func main() {
	fmt.Println("math tricks")
	// for i := 2; i <= 13; i++ {
	// 	fmt.Println(isPrime(i))
	// }
	//fmt.Println(isPrime(691731))
	optimisedForPrime(36)
	fmt.Println(differenceOfSums(10, 3))

}

func differenceOfSums(n int, m int) int {
	num1 := 0
	num2 := 0

	for i := 1; i <= n; i++ {
		if i%m != 0 {
			num1 = num1 + i
			fmt.Println(i)
		}
	}

	for i := 1; i <= m; i++ {
		if i%m == 0 {
			num2 = num2 + i
			fmt.Println(i)
		}
	}
	fmt.Println(num2)

	return num1 - num2
}

func optimisedForPrime(n int) {
	arr := make([]bool, n)

	for i := 2; i < n; i++ {
		arr[i] = true
		fmt.Println(i, arr[i])
	}

	for i := 3; i <= n; i++ {
		if n%i == 0 {
			start := i
			for start < n {
				if n%i == 0 {
					arr[start-1] = false
					fmt.Println(i, arr[start-1])
				}
				start = start * i
			}
		}
	}
	fmt.Println(arr)
}

func isPrime(n int) bool {
	//prime no starts from 2, any no divisible by 1 or itself called primes
	if n <= 1 {
		return false
	}
	//lets check if n is 36, it doesn't make sense to check from 2 to n
	// square root of n

	//trick to use start*start < n
	// will basically run for square root of n times
	//  start <= squareroot(n)
	//take square on both sides
	// start*start = n
	start := 2
	for start*start <= n {
		fmt.Println("start", start)
		if n%start == 0 {
			return false
		}
		start++
	}
	return true
}
