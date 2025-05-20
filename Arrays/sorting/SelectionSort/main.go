package main

import (
	"fmt"
)

func main() {

	nums := []int{12, 1, 34, 9, 33, 6}
	//nums := []int{1}

	fmt.Println("before sorting", nums)
	fmt.Println("after sorting ", selectionSort(nums))
	fmt.Println("after sorting ", selectionSortWithPositives(nums))
}

// note: we have found lowest no and then its position that is starting of the array
// you can also find the greatest no and its position that is last
// nums := []int{12, 1, 34, 9, 33, 6}
func selectionSortWithPositives(nums []int) []int {
	length := len(nums)

	for i := 0; i < length-1; i++ {
		max := findMaximum(nums, length-i)
		nums[max], nums[length-i-1] = nums[length-i-1], nums[max]
		fmt.Println("array after swap:", nums)
	}
	return nums
}

func findMaximum(nums []int, end int) int {
	max := 0
	for i := 1; i < end; i++ {
		if nums[i] > nums[max] {
			max = i
		}
	}
	return max
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
