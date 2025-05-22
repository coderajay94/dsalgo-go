package main

import (
	"fmt"
)

func main() {
	fmt.Println("Printing Patterns")
	// pattern1(5)
	// fmt.Println()
	// pattern2(5)
	// fmt.Println()
	// pattern3(5)
	// fmt.Println()
	// pattern4(5)

	// fmt.Println()
	// pattern5(5)

	fmt.Println()
	pattern6(5)
}

func pattern6(times int) {
	totalCols := 0
	for row := 1; row <= 2*times-1; row++ {
		//print spaces
		//1,2,3,4,5,
		if row > times {
			totalCols = 2*times - row
		} else {
			totalCols = row
		}

		totalSpaces := times - totalCols
		for s := 1; s <= totalSpaces; s++ {
			fmt.Print(" ")
		}
		//print *
		for j := 1; j <= totalCols; j++ {
			fmt.Print("* ")
		}
		//counter++
		fmt.Println()
	}

}

func pattern5(times int) {
	counter := 2
	for i := 1; i <= 2*times-1; i++ {
		end := i
		if i > times {
			end = end - counter
			counter = counter + 2
		}
		for j := 1; j <= end; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func pattern4(times int) {
	for i := 1; i <= times; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

func pattern3(times int) {
	for i := times; i > 0; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func pattern1(times int) {
	for i := 1; i <= times; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func pattern2(times int) {
	for i := 1; i <= times; i++ {
		for j := 1; j <= times; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
