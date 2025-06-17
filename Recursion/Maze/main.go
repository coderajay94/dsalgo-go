package main

import "fmt"

func main() {

	maze := [][]bool{
		{true, true, true},
		{true, true, true},
		{true, true, true},
	}

	//solveMaze(maze, 0, 0, "")
	result := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	solveMazeWithPath(maze, 0, 0, result, 1)
}

func solveMazeWithPath(maze [][]bool, row, column int, result [][]int, steps int) {
	//fmt.Println(row, column, path)
	//res := [][]int{}

	if row == len(maze)-1 && column == len(maze[0])-1 {
		result[row][column] = steps
		for i := 0; i < len(result); i++ {
			fmt.Println(result[i])
		}
		fmt.Println()
		return
	}

	if maze[row][column] == false {
		return
	}

	maze[row][column] = false
	result[row][column] = steps

	if row < len(maze)-1 {
		solveMazeWithPath(maze, row+1, column, result, steps+1)
	}

	if column < len(maze[0])-1 {
		solveMazeWithPath(maze, row, column+1, result, steps+1)
	}

	if row > 0 {
		solveMazeWithPath(maze, row-1, column, result, steps+1)
	}

	if column > 0 {
		solveMazeWithPath(maze, row, column-1, result, steps+1)
	}

	maze[row][column] = true
	result[row][column] = 0
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
