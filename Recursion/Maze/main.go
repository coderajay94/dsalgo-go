package main

import "fmt"

func main() {

	maze := [][]bool{
		{true, true, true},
		{true, true, true},
		{true, true, true},
	}

	solveMaze(maze, 0, 0, "")
}

func solveMaze(maze [][]bool, row, column int, path string) {
	//fmt.Println(row, column, path)

	if row == len(maze)-1 && column == len(maze[0])-1 {
		fmt.Println(path)
		return
	}

	if maze[row][column] == false {
		return
	}
	maze[row][column] = false

	if row < len(maze)-1 {
		solveMaze(maze, row+1, column, path+" Down")
	}

	if column < len(maze[0])-1 {
		solveMaze(maze, row, column+1, path+" Right")
	}

	if row > 0 {
		solveMaze(maze, row-1, column, path+" Up")
	}

	if column > 0 {
		solveMaze(maze, row, column-1, path+" Down")
	}

	// if row < len(maze)-1 && column < len(maze[0])-1 {
	// 	solveMaze(maze, row+1, column+1, path+" DiagonalDown")
	// }

	// if row > 0 && column > 0 {
	// 	solveMaze(maze, row-1, column-1, path+" DiagonalUp")
	// }

	maze[row][column] = true
}
