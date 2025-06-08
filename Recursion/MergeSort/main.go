package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{5, 4, 8, 1, 9}
	// fmt.Println("before sorting array ", nums)
	// fmt.Println("after sorting array ", mergeSort(nums), 0, len(nums))
	mergeSortInPlace(nums, 0, len(nums))
	fmt.Println("after sorting array", nums)

	testSlice(nums)
	fmt.Println(nums)

	m := make(map[string]string)
	m["ajay"] = "kumar"
	m["raghu"] = "kumar"
	fmt.Println(m)
	changeMap(m)
	fmt.Println(m)
	fmt.Println(digitCount("1210"))
}

func digitCount(num string) bool {

	m := make(map[int]int)

	for _, value := range num {
		numb, _ := strconv.Atoi(string(value))
		fmt.Println("iterating num", numb)
		m[numb]++
	}

	for i := 0; i < len(num); i++ {

		count, _ := strconv.Atoi(string(num[i]))
		if m[i] != count {
			return false
		}
	}
	return true
}

func changeMap(m map[string]string) {
	m["riddhi"] = "kumaar"
}

func testSlice(nums []int) {
	nums[0] = 0
}

func mergeSortInPlace(nums []int, start, end int) {

	if end-start == 1 {
		return
	}

	mid := start + (end-start)/2

	mergeSortInPlace(nums, start, mid)
	mergeSortInPlace(nums, mid, end)

	mergeInPlace(nums, start, mid, end)
}

func mergeInPlace(nums []int, start, mid, end int) {

	results := make([]int, end-start)
	i, j, k := start, mid, 0

	//compare both the arrays and put the smallest element in the results array k
	for i < mid && j < end {

		if nums[i] < nums[j] {
			results[k] = nums[i]
			i++
		} else {
			results[k] = nums[j]
			j++
		}
		k++
	}

	//add the remaining elements from left array
	for i < mid {
		results[k] = nums[i]
		i++
		k++
	}

	//add the remaining elements from right array
	for j < end {
		results[k] = nums[j]
		j++
		k++
	}

	// you
	for i := 0; i < len(results); i++ {
		nums[start+i] = results[i]
	}

	//fmt.Println("printing results", results)
}

func mergeSort(nums []int) []int {
	//if found 1 elements that means its already sorted
	if len(nums) == 1 {
		return nums
	}

	mid := len(nums) / 2
	//here end is not exclusive while copying, thats why start to mid-1, mid to len(nums)-1
	// will be copied
	left := mergeSort(nums[0:mid])
	right := mergeSort(nums[mid:len(nums)])

	return merge(left, right)

}

func merge(left, right []int) []int {

	ll := len(left)
	lr := len(right)

	results := make([]int, ll+lr)
	i, j, k := 0, 0, 0

	//compare both the arrays and put the smallest element in the results array k
	for i < ll && j < lr {

		if left[i] < right[j] {
			results[k] = left[i]
			i++
			fmt.Println("1", left, right, k)
		} else {
			results[k] = right[j]
			j++

			fmt.Println("2", left, right, k)
		}
		k++
		fmt.Println("3", left, right, k)
	}

	//add the remaining elements from left array
	for i < ll {
		results[k] = left[i]
		i++
		k++
	}

	//add the remaining elements from right array
	for j < lr {
		results[k] = right[j]
		j++
		k++
	}

	return results

}
