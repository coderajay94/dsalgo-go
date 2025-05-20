package main

import (
	"fmt"
)

func main() {
	fmt.Println("running search in a 2d sorted array")

	// matrix := [][]int{
	// 	{1, 2, 3, 4},
	// 	{5, 6, 7, 8},
	// 	{9, 10, 11, 12},
	// 	{13, 14, 15, 16},
	// }

	// matrix := [][]int{
	// 	{1, 2, 3, 4},
	// }

	matrix := [][]int{
		{1}, {3},
	}

	// matrix := [][]int{
	// 	{1, 2, 3, 4},
	// 	{5, 6, 7, 8},
	// 	{9, 10, 11, 12},
	// 	{13, 14, 15, 16},
	// }

	row, col := FindTargetInSortedMatric(matrix, 1)
	fmt.Println(row, col)

}

func FindTargetInSortedMatric(matrix [][]int, target int) (int, int) {

	rows := len(matrix) - 1
	cols := len(matrix[0]) - 1

	if rows == 0 {
		fmt.Println("going to call 0 binary search")
		return BinarySearch(matrix, 0, target, 0, cols)
	}

	rowsStart := 0
	rowsEnd := rows

	colsMid := (0 + cols) / 2
	for rowsStart < rowsEnd-1 {
		fmt.Println("checking", rowsStart, rowsEnd)

		rowsMid := rowsStart + (rowsEnd-rowsStart)/2

		fmt.Println("checking", rowsMid)
		if matrix[rowsMid][colsMid] == target {
			return rowsMid, colsMid
		}

		if matrix[rowsMid][colsMid] > target {
			rowsEnd = rowsMid
		} else {
			rowsStart = rowsMid
		}
	}

	fmt.Println("rows and rows end", rowsStart, rowsEnd)

	if target >= matrix[rowsStart][0] && target <= matrix[rowsStart][cols] {
		return BinarySearch(matrix, rowsStart, target, 0, cols)
	} else {
		fmt.Println("checking here")
		return BinarySearch(matrix, rowsEnd, target, 0, cols)
	}

}

func BinarySearch(matrix [][]int, row, target, start, end int) (int, int) {
	fmt.Println("startting to check:", row, start, end)
	for start <= end {
		mid := start + (end-start)/2

		if matrix[row][mid] == target {
			fmt.Println("returned ", row, mid)
			return row, mid
		}

		if matrix[row][mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}

	}

	return -1, -1
}
