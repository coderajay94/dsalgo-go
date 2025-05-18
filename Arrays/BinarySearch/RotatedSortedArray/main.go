package main

import (
	"fmt"
)

func main() {

	//nums := []int{4, 5, 6, 7, 0, 1, 2}
	nums := []int{1, 3, 5}
	target := 3
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {

	pivot := findPivot(nums)
	fmt.Println("pivot", pivot)
	if pivot == -1 {
		//array is not rotated
		return binarySearch(nums, target, 0, len(nums)-1)
	}
	if nums[pivot] == target {
		return pivot
	}

	if target < nums[0] {
		//look on the right side
		return binarySearch(nums, target, pivot+1, len(nums)-1)
	} else {
		//look on the left side
		return binarySearch(nums, target, 0, pivot-1)
	}
	return -1
}

func findPivot(nums []int) int {

	start, end := 0, len(nums)-1

	for start <= end {
		//instead of checking this way
		//mid := (start + end) / 2

		mid := start + (end-start)/2

		if mid < end && nums[mid] > nums[mid+1] {
			//condition 1 -  when found the pivot like mid is 7 and mid+1 = 0
			return mid
		} else if mid > start && nums[mid-1] > nums[mid] {
			//condition 2- when mid is 0 and mid-1 is 7
			return mid - 1
		} else if nums[mid] <= nums[start] {
			//condition 3 - on the right side as lesser than start of nums, so leave the right side part
			end = mid - 1
		} else {
			//condition 4 - when on the left side of the array
			start = mid + 1
		}

	}

	return -1
}

// returns the index of the target
func binarySearch(nums []int, target, start, end int) int {

	for start <= end {
		mid := (start + end) / 2

		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			return mid
		}

	}
	return -1
}
