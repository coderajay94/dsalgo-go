package main

import (
	"fmt"
)

func main() {

	//nums := []int{12, 1, 34, 9, 33, 6}
	nums := []int{1}

	fmt.Println("before sorting", nums)
	fmt.Println("after sorting ", selectionSort(nums))
}

func selectionSort(nums []int) []int {
	//iterate over the array and find the minimum element and then
	//put it on the first place and let the first for loop move to next and next
	swaps := 0
	for i := 0; i < len(nums)-1; i++ {
		lowest := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[lowest] {
				lowest = j
				swaps++
			}
		}

		nums[i], nums[lowest] = nums[lowest], nums[i]
	}
	fmt.Println("total swaps done:", swaps)
	return nums
}
