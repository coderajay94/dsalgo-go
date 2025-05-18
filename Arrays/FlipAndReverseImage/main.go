package main

import "fmt"

func main() {
	nums := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	//Output: [[1,0,0],[0,1,0],[1,1,1]]

	fmt.Println("reverse and invert is:", flipAndInvertImage(nums))
}

func flipAndInvertImage(image [][]int) [][]int {
	result := make([][]int, len(image))
	for index, val := range image {
		arr := reverseAndZeroArray(val)
		fmt.Println("array reversed and inverted is", arr)
		result[index] = arr
	}
	return result
}

func reverseAndZeroArray(nums []int) []int {
	start := 0
	end := len(nums) - 1
	length := len(nums)

	if length == 1 {
		if nums[0] == 0 {
			nums[0] = 1
		} else if nums[0] == 1 {
			nums[0] = 0
		}
	}
	if length == 1 || length == 0 {
		return nums
	}

	for end > start {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
	for index, _ := range nums {
		if nums[index] == 0 {
			nums[index] = 1
		} else if nums[index] == 1 {
			nums[index] = 0
		}
	}
	return nums
}
