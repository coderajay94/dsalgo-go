package main

import "fmt"

func main() {

	fmt.Println("print numbers between 1 and target ")
	target := 40

	seive(target)
}

func seive(target int) {
	primes := make([]bool, target+1)

	for i := 2; i*i <= target; i++ {
		if primes[i] == false {
			for j := i * 2; j <= target; j = j + i {
				primes[j] = true
			}
		}

	}

}
