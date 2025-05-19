package main

import (
	"fmt"
)

func main() {
	grid := [][]int{
		{4, 3, 2, -1},
		{3, 2, 1, -1},
		{1, 1, -1, -2},
		{-1, -1, -2, -3}}
	fmt.Println(countNegatives(grid))
}

func countNegatives(grid [][]int) int {

	rowsLength := len(grid) - 1
	rows := 0
	cols := len(grid[0]) - 1

	sum := 0

	for rows < len(grid) && cols >= 0 {
		//fmt.Println("row and columns is:", rows, cols, grid[rows][cols])

		if grid[rows][cols] < 0 {
			//fmt.Println("adding sum now is:", rows, cols, sum)
			sum = sum + rowsLength - rows + 1
			cols--
		} else {
			rows++
		}
	}
	return sum
}
