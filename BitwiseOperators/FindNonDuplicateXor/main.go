package main

import "fmt"

func main() {

	//if you xor any number with same nmber gives 0
	// and zero exor number gives the number itself
	//1 * 0 = 1
	//0 * 1 = 1 // any number exor with 0 gives the same number
	//1 * 1 = 0 // any number exor with same number gives zero
	//0 * 0 = 0

	nums := []int{1, 2, 3, 3, 4, 2, 6, 4, 1}
	fmt.Println(ans(nums))

}

func ans(nums []int) int {
	result := 0
	for _, value := range nums {
		//if any number will get xor twice that will become zero
		//only single element would be left
		result = result ^ value
	}
	return result
}
