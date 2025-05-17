package main

import "fmt"

func main() {
	//nums2 := []int{2, 4, 5, 8, 3, 2, 1}
	nums := []int{1, 5, 2}
	target := 2
	fmt.Println("target index is ", findInMountainArray(target, nums))
}

func findInMountainArray(target int, nums []int) int {

	peakIndex := FindPeakElementindex(nums)
	fmt.Println("peak index is", peakIndex)

	if target == nums[peakIndex] {
		return peakIndex
	}
	left := BinarySearch(nums, target, true, 0, peakIndex-1)
	if left == -1 {
		return BinarySearch(nums, target, false, peakIndex+1, len(nums)-1)
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
	fmt.Println("start and end is", start, end)
	for start <= end {

		mid := (start + end) / 2
		midVal := nums[mid]
		if nums[mid] == target {
			return mid
		}
		/// 6, 4, 3, 2, 1
		if isAsc {
			if midVal > target {
				end = mid - 1
			} else if midVal < target {
				start = mid + 1
			}
		} else {
			if midVal > target {
				start = mid + 1
			} else if midVal < target {
				end = mid - 1
			}
		}
	}
	return -1
}

//
/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

// func findInMountainArray(target int, mountainArr *MountainArray) int {

// 	peakIndex := FindPeakElementindex(mountainArr)
// 	//fmt.Println("peak index is", peakIndex)

// 	if target == mountainArr.get(peakIndex) {
// 		return peakIndex
// 	}
// 	left := BinarySearch(mountainArr, target, true, 0, target-1)
// 	if left == -1 {
// 		return BinarySearch(mountainArr, target, false, target+1, mountainArr.length()-1)
// 	} else {
// 		return left
// 	}

// 	return -1
// }

// func FindPeakElementindex(mountainArr *MountainArray) int {

// 	start, end := 0, mountainArr.length()-1

// 	for start < end {
// 		mid := (start + end) / 2
// 		if mountainArr.get(mid) < mountainArr.get(mid+1) {
// 			start = mid + 1
// 		} else {
// 			//when mid is larger than mid+1, that would be the potential answer, but move to the left
// 			end = mid
// 		}

// 	}
// 	return start
// }

// //returns the index
// func BinarySearch(mountainArr *MountainArray, target int, isAsc bool, start int, end int) int {

// 	//start, end := 0, len(nums)-1

// 	for start <= end {

// 		mid := (start + end) / 2

// 		if mountainArr.get(mid) == target {
// 			return mid
// 		}
// 		/// 6, 4, 3, 2, 1
// 		if isAsc {
// 			if mid > target {
// 				end = mid - 1
// 			} else if mid < target {
// 				start = mid + 1
// 			}
// 		} else {
// 			if mid > target {
// 				start = mid + 1
// 			} else if mid < target {
// 				end = mid - 1
// 			}
// 		}
// 	}
// 	return -1
// }
