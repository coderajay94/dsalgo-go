package main

import "fmt"

func main() {

	//nums := make([]int, 2)

	// for i := 1; i < 1026; i++ {
	// 	//if i/10 == 0 {
	// 	fmt.Println(cap(nums), len(nums))
	// 	//}
	// 	nums = append(nums, i)
	// }

	nums := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(len(nums), cap(nums))
	fmt.Println(nums)
	nums2 := append(nums[:2])
	fmt.Println(nums2)

	nums3 := make([]int, 3)
	fmt.Println(len(nums3), cap(nums3))
	nums3 = append(nums[:])
	fmt.Println(len(nums3), cap(nums3))
	nums3[0] = 10
	nums4 := make([]int, 5)
	fmt.Println(nums4)
	n := copy(nums4, nums)
	nums4[0] = 100
	fmt.Println(nums)
	fmt.Println(nums2)
	fmt.Println(nums3)
	fmt.Println(n)
	fmt.Println(nums4)

	// arr := []int{}
	// arr[0] = 12
	// if arr != nil {
	// 	fmt.Println("arr is not nil", len(arr), cap(arr), arr)
	// } else {
	// 	fmt.Println("arr is nil", len(arr), cap(arr), arr)
	// }

	s := make([]int, 3)

	callme(s)
	fmt.Println("main s:", s)
}

func callme(nums []int) {
	fmt.Println(len(nums), cap(nums))
	for i := 0; i < len(nums); i++ {
		nums[i] = i + 1
	}
	nums = append(nums, 7)
	fmt.Println("callme:", nums)
}
