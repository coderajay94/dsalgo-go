package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 2, 88, 1, 99}
	fmt.Println("bubble sort using recursion", nums)
	fmt.Println("bubble sort using recursion", bubbleSort(nums, 0, len(nums)-1))
}

func bubbleSort(nums []int, start int, end int) []int {

	if end == 0 {
		return nums
	}

	if start < end {
		if nums[start] > nums[start+1] {
			nums[start], nums[start+1] = nums[start+1], nums[start]
		}
		return bubbleSort(nums, start+1, end)
	} else {
		return bubbleSort(nums, 0, end-1)
	}
}
