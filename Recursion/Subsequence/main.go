package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}

func subsets(nums []int) [][]int {

	res := [][]int{}
	output := []int{}
	index := 0

	solve(nums, index, output, &res)
	return res
}

func solve(nums []int, index int, output []int, res *[][]int) {

	if index >= len(nums) {
		// copy output to avoid mutation in future calls
		temp := make([]int, len(output))
		copy(temp, output)
		*res = append(*res, temp)
		return
	}

	//exclude
	solve(nums, index+1, output, res)
	//include
	num := nums[index]
	output = append(output, num)
	solve(nums, index+1, output, res)
}
