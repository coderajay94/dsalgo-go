package main

import (
	"fmt"
)

func main() {
	nums := []int{5, 4, 8, 1, 9}
	fmt.Println("before sorting array ", nums)
	fmt.Println("after sorting array ", mergeSort(nums))
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
