package main

import "fmt"

func main() {
	// nums := []int{1, 2, 3}
	// fmt.Println(subsets(nums))

	str := "abc"
	fmt.Println(subseq(str))
}

func subseq(str string) []string {
	return sub("", str, []string{})
}

func sub(pr string, un string, resp []string) []string {

	if un == "" {
		res := []string{pr}
		return res
	}

	ch := string(un[0])
	//resp = append(resp, ch)

	left := sub(pr+ch, un[1:], resp)
	right := sub(pr, un[1:], resp)
	return append(left, right...)
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
