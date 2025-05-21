package main

import (
	"fmt"
)

func main() {

	nums := []int{3, 2, 1, 4, 5, 7, 6}
	fmt.Println("before sorting", nums)
	fmt.Println("before sorting", cyclicSort(nums))

	nums2 := []int{1, 0, 3, 4}

	fmt.Println(missingNumber(nums2))
}

// https://leetcode.com/problems/missing-number/
func missingNumber(nums []int) int {

	start := 0

	for start < len(nums) {

		actual := nums[start]

		if nums[start] != len(nums) && nums[start] != nums[actual] {
			nums[start], nums[actual] = nums[actual], nums[start]
		} else {
			start++
		}
	}

	//iterate over all the elements and check if index is not same as value
	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	//if array is already sorted, then return the last index+1,
	return len(nums)
}

func cyclicSort(nums []int) []int {

	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}

	start := 0
	for start < len(nums) {
		actual := nums[start] - 1

		if nums[start] != nums[actual] {
			nums[start], nums[actual] = nums[actual], nums[start]
		} else {
			start++
		}

	}
	return nums
}
