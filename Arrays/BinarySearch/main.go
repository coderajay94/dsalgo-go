// package main

// import "fmt"

// func main() {

// 	fmt.Println("Running Binary Search...")

// 	elements := []int{12, 40, 56, 122, 134, 245, 4567, 5678}
// 	target := 245

// 	fmt.Println("Finding ", target, " using binary search on index:", BinarySearch(elements, target))
// }

// func BinarySearch(elements []int, target int) int {

// 	start := 0
// 	end := len(elements) - 1
// 	for start <= end {
// 		fmt.Println("checking start and end is", start, end)
// 		mid := (start + end) / 2
// 		fmt.Println("mid is", mid)
// 		if elements[mid] == target {
// 			return mid
// 		} else if elements[mid] > target {
// 			end = mid - 1
// 		} else if elements[mid] < target {
// 			start = mid + 1
// 		}
// 	}

// 	return -1
// }

//infinite array problem
// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	nums := []int{
		1, 3, 5, 7, 8, 10, 11, 12, 15, 17, 18, 19, 20, 21, 23, 25, 26, 29,
		31, 32, 34, 35, 38, 40, 42, 43, 45, 47, 48, 50, 52, 54, 55, 56, 58,
		60, 62, 64, 66, 68, 70, 71, 73, 75, 77, 79, 81, 83, 86, 89, 90, 93,
		95, 97, 99,
	}
	fmt.Println("target", findTarget(nums, 21))
}

func findTarget(nums []int, target int) int {
	start := 0
	end := 1
	for start < end {
		fmt.Println("start", start, "end", end)
		if nums[start] <= target && target <= nums[end] {
			return search(nums, target, start, end)
		} else {
			start = end + 1
			end = end + (end+1)*2
		}
	}
	return -1
}

func search(nums []int, target, start, end int) int {

	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		}
		if target > nums[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

