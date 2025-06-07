package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 5, 7, 3, 0, 11, 4}
	nums2 := []int{1, 2, 5, 7, 11}
	fmt.Println("check if array is sorted", nums, isSorted(nums, 0))
	fmt.Println("check if array is sorted", nums2, isSorted(nums2, 0))
}

func isSorted(nums []int, start int) bool {
	if start == len(nums)-1 {
		return true
	}

	return nums[start] < nums[start+1] && isSorted(nums, start+1)
}
