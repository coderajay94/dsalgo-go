package main

import "fmt"

func main() {

	nums := [][]int{
		{7, 3, 1, 9},
		{3, 4, 6, 9},
		{6, 9, 6, 6},
		{9, 5, 8, 5}}

	fmt.Println(diagonalSum((nums)))

}

func diagonalSum(mat [][]int) int {

	sum := 0
	length := len(mat) - 1
	for i := 0; i < len(mat); i++ {

		if i == (length - i) {
			sum = sum + mat[i][i]
		} else {
			sum = sum + mat[i][i] + mat[i][length-i]
		}
		fmt.Println("indexes", i, length-i)
		fmt.Println("adding ", mat[i][i], mat[i][length-i])
	}
	return sum
}
