package main

import (
	"fmt"
)

func main() {

	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60}}

	r, c := searchMatrix(matrix, 34)
	fmt.Println("target at index:", r, c)

}

func searchMatrix(matrix [][]int, target int) (int, int) {
	rows := 0
	cols := len(matrix[0]) - 1

	for cols >= 0 && rows < len(matrix) {

		//condition match
		if matrix[rows][cols] == target {
			return rows, cols
		}

		//greater element, remove colum
		if matrix[rows][cols] > target {
			cols--
		} else {
			// remove row
			rows++
		}
	}

	return -1, -1
}
