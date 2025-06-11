package main

import (
	"fmt"
	"strconv"
)

func main() {
	// nums := []int{1, 2, 3}
	// fmt.Println(subsets(nums))

	// str := "abc"
	// fmt.Println(subseq(str))

	//subAscii("", "abc")
	subSequenceIteratively("abc")
	subsetsOfNumbersIteratively([]int{1, 2, 3, 4})

}

func subsetsOfNumbersIteratively(n []int) {

	//logic to remember
	// 	Example Flow with [1, 2]:
	// Start: all = [[]]
	// num = 1:

	// Add 1 to each existing subset:
	// [] + 1 → [1]
	// Now all = [[], [1]]

	// num = 2:
	// Add 2 to each existing subset:
	// [] + 2 → [2]
	// [1] + 2 → [1, 2]

	// Now all = [[], [1], [2], [1, 2]]
	//starts with empty
	all := [][]int{{}}

	//update each existing array with new element
	for _, value := range n {
		currentSize := len(all)

		for i := 0; i < currentSize; i++ {
			new := append(all[i], value)
			all = append(all, new)
		}
	}

	fmt.Println(len(all), all)
}

// iterative solutions not efficient
func subSequenceIteratively(str string) {

	list := []string{""}

	for _, val := range str {

		currentSize := len(list)
		newlist := []string{}
		for j := 0; j < currentSize; j++ {
			newlist = append(newlist, list[j]+string(val))
		}
		list = append(list, newlist...)
	}
	fmt.Println(len(list), list)
}

func subseq(str string) []string {
	return sub("", str)
}

func sub(pr string, un string) []string {

	if un == "" {
		res := []string{pr}
		return res
	}

	ch := string(un[0])
	//resp = append(resp, ch)

	left := sub(pr+ch, un[1:])
	right := sub(pr, un[1:])
	return append(left, right...)
}

func subAscii(pr string, un string) {
	//fmt.Println("inside this", pr)
	if un == "" {
		fmt.Println(pr)
		return
	}

	ch := un[0]
	//resp = append(resp, ch)

	subAscii(pr+string(ch), un[1:])
	subAscii(pr, un[1:])
	subAscii(pr+strconv.Itoa(int(ch)), un[1:])
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
