package main

import (
	"fmt"
)

func main() {

	//nums := []int{11, 23, 13, 45, 56, 14, 6}
	nums := []int{-23, 0, -1}

	// 11 23
	// 11 13 23 45 56 14
	// 11 13 23 14 45 56
	fmt.Println("sort using insertion sort", insertionSort(nums))
}

func insertionSort(nums []int) []int {
	length := len(nums)

	for i := 0; i < length-1; i++ {

		for j := i + 1; j > 0; j-- {

			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			} else {
				break
			}

		}

	}

	return nums
}
