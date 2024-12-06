package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

}

func main() {
	fmt.Println("Running")

	var start = time.Now()

	partOne()
	// partTwo()

	var end = time.Since(start)

	fmt.Println("Execution time: ", end.Seconds())

}
