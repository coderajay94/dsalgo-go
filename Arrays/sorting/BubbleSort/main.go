package main

import "fmt"

//bubble sort is also known as sinking sort, exchange sort
func main() {
	//nums := []int{11, 23, 44, 33, 12, 89, 56}
	nums := []int{1, 2, 12, 19, 23, 45}

	fmt.Println("before sorting array :", nums)

	sortByReference(&nums)

	fmt.Println("sorted array with reference:", nums)
	fmt.Println("sorted array with value:", sortByValue(nums))
}

//bubble sort with pass by reference
func sortByValue(nums []int) []int {
	swapCounter := 0
	for i := 0; i < len(nums)-1; i++ {
		swapCounter = 0
		for j := i + 1; j < len(nums)-i; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
				swapCounter++
			}
		}
		if swapCounter == 0 {
			//array is already sorted
			return nums
		}
	}
	return nums

}

//bubble sort with pass by reference
func sortByReference(nums *[]int) {
	swapCounter := 0
	for i := 0; i < len(*nums)-i; i++ {
		swapCounter = 0
		for j := i + 1; j < len(*nums)-i; j++ {
			if (*nums)[i] > (*nums)[j] {
				(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
				swapCounter++
			}
		}
		if swapCounter == 0 {
			//array is already sorted
			return
		}

	}
	fmt.Println("executing this outside the loop")
}
