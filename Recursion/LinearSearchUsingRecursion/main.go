package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 5, 7, 3, 3, 11, 3}
	target := 14

	fmt.Println("before linear search using recursion", nums)
	fmt.Println("after linear search using recursion", linearSearch(nums, target, 0, len(nums)-1))
	fmt.Println("all matched elements as target", linearSearchAllOccurances(nums, 3, 0, []int{}))
}

// var resp = []int{}
// or
// directly pass an slice argument to return the response
// get all the occurances of the target
func linearSearchAllOccurances(nums []int, target, start int, resp []int) []int {
	if start == len(nums) {
		return resp
	}

	if nums[start] == target {
		resp = append(resp, start)
	}
	return linearSearchAllOccurances(nums, target, start+1, resp)
}

func linearSearch(nums []int, target int, start, end int) int {
	if start > end {
		return -1
	}

	if nums[start] == target {
		return start
	}
	return linearSearch(nums, target, start+1, end)
}
