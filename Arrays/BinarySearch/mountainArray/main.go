package main

import "fmt"

func main() {
	nums := []int{2, 4, 5, 8, 3, 2, 1}
	target := 1
	fmt.Println("target index is ", findInMountainArray(target, nums))
}

func findInMountainArray(target int, nums []int) int {

	peakIndex := FindPeakElementindex(nums)
	fmt.Println("peak index is", peakIndex)

	if target == nums[peakIndex] {
		return peakIndex
	}
	left := BinarySearch(nums, target, true, 0, target-1)
	if left == -1 {
		return BinarySearch(nums, target, false, target+1, len(nums)-1)
	} else {
		return left
	}

	return -1
}

func FindPeakElementindex(nums []int) int {

	start, end := 0, len(nums)-1

	for start < end {
		mid := (start + end) / 2
		if nums[mid] < nums[mid+1] {
			start = mid + 1
		} else {
			//when mid is larger than mid+1, that would be the potential answer, but move to the left
			end = mid
		}

	}
	return start
}

//returns the index
func BinarySearch(nums []int, target int, isAsc bool, start int, end int) int {

	//start, end := 0, len(nums)-1

	for start <= end {

		mid := (start + end) / 2

		if nums[mid] == target {
			return mid
		}
		/// 6, 4, 3, 2, 1
		if isAsc {
			if mid > target {
				end = mid - 1
			} else if mid < target {
				start = mid + 1
			}
		} else {
			if mid > target {
				start = mid + 1
			} else if mid < target {
				end = mid - 1
			}
		}
	}
	return -1
}
