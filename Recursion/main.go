package main

import "fmt"

func main() {
	//user can go only down or right
	fmt.Println("Maze Problem")
	//fmt.Println(solveMaze(3, 3))
	//solveMazePath("", 3, 3)
	fmt.Println(solveMazePathWithSlice("", 3, 3))
}

func solveMaze(row, column int) int {

	//even if one of row or column reached to 1, we know only one path exists now
	if row == 1 || column == 1 {
		return 1
	}

	left := solveMaze(row-1, column)
	right := solveMaze(row, column-1)
	return left + right
}

func solveMazePath(path string, row, column int) {

	if row == 1 && column == 1 {
		fmt.Println(path)
		return
	}

	if row > 1 {
		solveMazePath(path+"D", row-1, column)
	}
	if column > 1 {
		solveMazePath(path+"R", row, column-1)
	}

}

func solveMazePathWithSlice(path string, row, column int) []string {

	res := []string{}
	if row == 1 && column == 1 {
		res = append(res, path)
		return res
	}

	if row > 1 {
		res = append(res, solveMazePathWithSlice(path+"D", row-1, column)...)
	}
	if column > 1 {
		res = append(res, solveMazePathWithSlice(path+"R", row, column-1)...)
	}
	return res
}
