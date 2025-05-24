package main

import "fmt"

func main() {

	//13 elements
	nums := []int{1, 3, 5, 7, 9, 11, 12, 13, 23, 34, 56, 67, 90}
	//nums := []int{2}

	fmt.Println(binarySearchUsingRecursion(nums, 0, len(nums)-1, 1))
}

//returns index of the target
func binarySearchUsingRecursion(nums []int, start, end, target int) int {
	if start > end {
		return -1
	}

	mid := start + (end-start)/2

	if nums[mid] == target {
		return mid
	}

	if nums[mid] > target {
		end = mid - 1
	} else {
		start = mid + 1
	}

	return binarySearchUsingRecursion(nums, start, end, target)
}
