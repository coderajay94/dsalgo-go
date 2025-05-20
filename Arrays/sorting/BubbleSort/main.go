package main

import "fmt"

func main() {
	nums := []int{11, 23, 44, 33, 12, 89}

	fmt.Println("before sorting array :", nums)

	sort(&nums)
	fmt.Println("sorted array :", nums)
}

//bubble so
func sort(nums *[]int) {

	for i := 0; i < len(*nums)-1; i++ {
		for j := i + 1; j < len(*nums)-1; j++ {
			if (*nums)[i] > (*nums)[j] {
				(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
			}
		}
	}

}
