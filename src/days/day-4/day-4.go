package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func partOne() {

	wd, _ := os.Getwd()
	filePath := wd + "/src/days/day-4/input.txt"
	file, _ := os.Open(filePath)

	scan := bufio.NewScanner(file)

	var arr []string

	for scan.Scan() {
		line := strings.TrimSpace(scan.Text())
		arr = append(arr, line)
	}

	result := 0

	yLen, xLen := len(arr), len(arr[0])

	for y, val := range arr {
		for x, letter := range val {
			if letter == 'X' {
				hasMas := findMas(&arr, y, yLen, x, xLen)
				result += hasMas
			}
		}
	}

	fmt.Println("result: ", result)
}

var directions = [8][2]int{
	{1, 1},   // Diagonal down-right
	{0, 1},   // Right
	{-1, 1},  // Diagonal up-right
	{-1, 0},  // Up
	{-1, -1}, // Diagonal up-left
	{0, -1},  // Left
	{1, -1},  // Diagonal down-left
	{1, 0},   // Down
}

var masArr = [3]rune{'M', 'A', 'S'}

func findMas(arr *[]string, y, yLen, x, xLen int) int {

	count := 0

	for _, val := range directions {
		has := hasMas(arr, y, yLen, x, xLen, &val)

		if has {
			count++
		}
	}

	return count
}

func hasMas(arr *[]string, y, yLen, x, xLen int, dir *[2]int) bool {

	totalY := y + (*dir)[0]*3
	if totalY >= yLen || totalY < 0 {
		return false
	}

	totalX := x + (*dir)[1]*3
	if totalX >= xLen || totalX < 0 {
		return false
	}

	for i := 0; i < 3; i++ {
		newY := y + (*dir)[0]*(i+1)
		newX := x + (*dir)[1]*(i+1)

		if rune((*arr)[newY][newX]) != masArr[i] {
			return false
		}
	}

	return true
}

var directions2 = [4][2]int{
	{1, 1},   // Diagonal down-right
	{-1, 1},  // Diagonal up-right
	{-1, -1}, // Diagonal up-left
	{1, -1},  // Diagonal down-left
}

func findXMas(arr *[]string, y, x int) bool {
	count := 0

	for _, dir := range directions2 {

		totalY := y + dir[0]
		totalX := x + dir[1]

		totalYReverse := y + dir[0]*-1
		totalXReverse := x + dir[1]*-1

		if (*arr)[totalY][totalX] == 'M' && (*arr)[totalYReverse][totalXReverse] == 'S' {
			count++
		}
	}

	return count == 2
}

func partTwo() {

	wd, _ := os.Getwd()
	filePath := wd + "/src/days/day-4/input.txt"
	file, _ := os.Open(filePath)

	scan := bufio.NewScanner(file)

	var arr []string

	for scan.Scan() {
		line := strings.TrimSpace(scan.Text())
		arr = append(arr, line)
	}

	result := 0

	yLen, xLen := len(arr), len(arr[0])

	for y, val := range arr {

		if y == 0 || y == yLen-1 {
			continue
		}

		for x, letter := range val {

			if x == 0 || x == xLen-1 {
				continue
			}

			if letter == 'A' {
				hasMas := findXMas(&arr, y, x)
				if hasMas {
					result += 1
				}
			}
		}
	}

	fmt.Println("result: ", result)
}

func main() {
	fmt.Println("Running")

	var start = time.Now()

	// partOne()
	partTwo()

	var end = time.Since(start)

	fmt.Println("Execution time: ", end.Seconds())

}
