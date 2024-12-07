package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

var directions = [4][2]int{
	{-1, 0}, // Up
	{0, 1},  // Right
	{1, 0},  // Down
	{0, -1}, // Left
}

func partOne() {

	wd, _ := os.Getwd()
	filePath := wd + "/src/days/day-6/input.txt"
	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	matrix := [][]string{}
	position := []int{}
	dirIndex := 0
	dir := directions[dirIndex]

	for y := 0; scanner.Scan(); y++ {
		line := strings.Split(strings.TrimSpace(scanner.Text()), "")
		// fmt.Println(line)

		found := slices.Index(line, "^")

		if found != -1 {
			position = []int{y, found}
		}

		matrix = append(matrix, line)
	}

	yLen := len(matrix)
	xLen := len(matrix[0])

	steps := 0

	for {
		newYPos := position[0] + dir[0]
		newXPos := position[1] + dir[1]

		if matrix[position[0]][position[1]] != "X" {
			steps++
		}

		matrix[position[0]][position[1]] = "X"

		if newYPos >= yLen || newYPos < 0 || newXPos >= xLen || newXPos < 0 {
			break
		}

		if matrix[newYPos][newXPos] == "#" {
			// change direction
			dirIndex = (dirIndex + 1) % len(directions)
			dir = directions[dirIndex]
		}

		position[0] = position[0] + dir[0]
		position[1] = position[1] + dir[1]
	}

	for _, val := range matrix {
		fmt.Println(val)
	}

	fmt.Println("Steps: ", steps)

}

func partTwo() {
	wd, _ := os.Getwd()
	filePath := wd + "/src/days/day-6/input.txt"
	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	matrix := [][]string{}
	guardPosition := []int{}
	guardDirectionIndex := 0
	guardDirection := directions[guardDirectionIndex]

	for y := 0; scanner.Scan(); y++ {
		line := strings.Split(strings.TrimSpace(scanner.Text()), "")
		// fmt.Println(line)

		found := slices.Index(line, "^")
		if found != -1 {
			guardPosition = []int{y, found}
		}
		matrix = append(matrix, line)
	}

	matrixYLength := len(matrix)
	matrixXLength := len(matrix[0])

	obstaclesFound := 0
	obstacleTurns := make(map[string][][]int)

	result := 0

	for {
		newYPos := guardPosition[0] + guardDirection[0]
		newXPos := guardPosition[1] + guardDirection[1]

		matrix[guardPosition[0]][guardPosition[1]] = "X"

		if newYPos >= matrixYLength || newYPos < 0 || newXPos >= matrixXLength || newXPos < 0 {
			break
		}

		if obstaclesFound > 2 {
			nextGuardDirectionIndex := directions[(guardDirectionIndex+1)%len(directions)]
			newGuardDirectionString := strconv.Itoa(nextGuardDirectionIndex[0]) + strconv.Itoa(nextGuardDirectionIndex[1])

			items, ok := obstacleTurns[newGuardDirectionString]

			isNextDirectionHorizontal := nextGuardDirectionIndex[1] == -1 || nextGuardDirectionIndex[1] == 1

			// fmt.Println("here", items, ok, isNextDirectionHorizontal, guardPosition)

			if ok {
				for _, item := range items {

					if isNextDirectionHorizontal {

						if item[0] == guardPosition[0] {
							fmt.Println("hello", guardPosition[0], guardPosition[1])

							isValid := true

							if isValid {
								result++
							}
						}

					} else {
						if item[1] == guardPosition[1] {
							fmt.Println("hello", guardPosition[0], guardPosition[1])

							isValid := true

							if isValid {
								result++
							}
						}

					}

				}
			}
		}

		if matrix[newYPos][newXPos] == "#" {
			// change direction
			// fmt.Println("coor: ", position[0], position[1], " | dir: ", dir)

			guardDirectionString := strconv.Itoa(guardDirection[0]) + strconv.Itoa(guardDirection[1])
			obstacleTurns[guardDirectionString] = append(obstacleTurns[guardDirectionString], []int{guardPosition[0], guardPosition[1]})

			obstaclesFound++

			guardDirectionIndex = (guardDirectionIndex + 1) % len(directions)
			guardDirection = directions[guardDirectionIndex]
		}

		guardPosition[0] = guardPosition[0] + guardDirection[0]
		guardPosition[1] = guardPosition[1] + guardDirection[1]
	}

	// fmt.Println(obstacleTurns)
	//
	// for _, val := range matrix {
	// 	fmt.Println(val)
	// }

	fmt.Println("Result: ", result)

}

func partTwoTwo() {
	wd, _ := os.Getwd()
	filePath := wd + "/src/days/day-6/input.txt"
	input, _ := os.ReadFile(filePath)

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	grid := make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(line)
	}
	output := 0

	guardRow := -1
	guardCol := -1
	guardDirection := 0

	for row, _ := range grid {
		if guardRow >= 0 {
			break
		}
		for col, _ := range grid[row] {
			if grid[row][col] == '^' {
				guardRow = row
				guardCol = col
				break
			}
		}
	}

	directions := [][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] != '.' {
				continue
			}

			grid[row][col] = '#'

			visitedLocations := make(map[[3]int]bool)
			currentRow := guardRow
			currentCol := guardCol
			currentDirection := guardDirection

			loopDetected := false

			for {
				guardState := [3]int{currentRow, currentCol, currentDirection}
				if visitedLocations[guardState] {
					loopDetected = true
					break
				}

				visitedLocations[guardState] = true

				nextGuardRow := currentRow + directions[currentDirection][0]
				nextGuardCol := currentCol + directions[currentDirection][1]

				if nextGuardRow < 0 || nextGuardRow >= len(grid) || nextGuardCol < 0 || nextGuardCol >= len(grid[0]) {
					break
				}

				if grid[nextGuardRow][nextGuardCol] == '#' {
					currentDirection = (currentDirection + 1) % 4
				} else {
					currentRow = nextGuardRow
					currentCol = nextGuardCol
				}
			}

			if loopDetected {
				output++
			}

			grid[row][col] = '.'
		}
	}

	fmt.Println("Output Day 6 Part 2", output)
}

func main() {
	fmt.Println("Running")

	var start = time.Now()

	// partOne()
	// partTwo()
	partTwoTwo()

	var end = time.Since(start)

	fmt.Println("Execution time: ", end.Seconds())

}
