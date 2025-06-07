package main

import (
	"fmt"
)

func main() {
	fmt.Println("search in rotated search array")
	nums := []int{5, 6, 7, 9, 11, 1, 2, 4}
	target := 2

	fmt.Println(RotatedBinarySearch(nums, target, 0, len(nums)-1))

}

func RotatedBinarySearch(nums []int, target int, start, end int) int {

	// 5, 6, 7, 8, 0, 1, 2, 4
	//case 1 - nums[start]<nums[mid] && target also less than nums[mid] and greater than nums[start]

	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		}

		if nums[start] < nums[mid] {
			// already sorted sorted array - left side
			if target > nums[start] && target < nums[mid] {
				return RotatedBinarySearch(nums, target, start, mid-1)
			} else {
				return RotatedBinarySearch(nums, target, mid+1, end)
			}
		}

		//rotated sorted part - right side
		if target > nums[mid] && target < nums[end] {
			return RotatedBinarySearch(nums, target, mid+1, end)
		}

		return RotatedBinarySearch(nums, target, start, mid-1)

	}

	return -1
}
