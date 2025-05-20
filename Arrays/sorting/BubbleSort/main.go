package main

import "fmt"

//bubble sort is also known as sinking sort, exchange sort
func main() {
	nums := []int{11, 23, 44, 33, 12, 89}

	fmt.Println("before sorting array :", nums)

	sortByReference(&nums)

	fmt.Println("sorted array with reference:", nums)
	fmt.Println("sorted array with value:", sortByValue(nums))
}

//bubble sort with pass by reference
func sortByValue(nums []int) []int {

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums)-i; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums

}

//bubble sort with pass by reference
func sortByReference(nums *[]int) {

	for i := 0; i < len(*nums)-1; i++ {
		for j := i + 1; j < len(*nums)-i; j++ {
			if (*nums)[i] > (*nums)[j] {
				(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
			}
		}
	}

}
