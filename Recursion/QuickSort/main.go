package main

import (
	"fmt"
)

func main() {
	//find a pivot and move that on the right place
	// before that move all smaller before
	//and all greater after that

	nums := []int{6, 4, 3, 2, 1, 0}
	fmt.Println("calling quick sort", nums)
	quickSort(nums, 0, len(nums)-1)
	fmt.Println("calling quick sort", nums)
}

func quickSort(nums []int, start, end int) {

	//find a pivot element
	//also find 2 pointers to move the elments around
	if start > end {
		return
	}

	low := start
	high := end
	mid := start + (end-start)/2

	pivot := nums[mid]

	for high >= low {

		//6, 4, 3, 2, 1, 0
		//0, 4, 3, 2, 1, 6
		//0, 1, 3, 2, 4, 6
		//0, 1, 2, 3, 4, 6
		//check if array is sorted by
		//checking if start and end elments r are correct position

		for nums[low] < pivot {
			low++
		}
		for nums[high] > pivot {
			high--
		}

		for low <= high {
			//swap because they are already ahead
			//both the conditions violated
			nums[low], nums[high] = nums[high], nums[low]
			low++
			high--
		}

	}

	quickSort(nums, start, high)
	quickSort(nums, low, end)

}
