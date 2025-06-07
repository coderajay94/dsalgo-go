package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 2, 88, 1, 99}
	fmt.Println("bubble sort using recursion", nums)
	fmt.Println("bubble sort using recursion", bubbleSort(nums, 0, len(nums)-1))
	fmt.Println("selection sort using recursion", selectionSort(nums, 0, len(nums)-1, 0))
}

//selection sort
// find the max then swap with the last element

func selectionSort(nums []int, start, end int, max int) []int {
	if end == 0 {
		return nums
	}

	if start < end {
		if nums[start+1] > nums[max] {
			max = start + 1
		}
		return selectionSort(nums, start+1, end, max)
	}

	nums[max], nums[end] = nums[end], nums[max]
	return selectionSort(nums, 0, end-1, 0)
}

// bubble sort
// swap till the largest reaches end
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
