package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16},
	}

	target := 15

	i, j := BinarySearch2d(matrix, target)
	fmt.Println("target found at", i, j)

}

func BinarySearch2d(matrix [][]int, target int) (int, int) {

	rowLength := len(matrix)
	colLength := len(matrix[0])

	if rowLength == 0 || colLength == 0 {
		return -1, -1
	}

	start := 0
	end := (rowLength * colLength) - 1

	for start <= end {
		mid := start + (end-start)/2
		i, j := mid/colLength, mid%colLength
		if matrix[i][j] == target {
			return i, j
		}

		if matrix[i][j] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1, -1
}
