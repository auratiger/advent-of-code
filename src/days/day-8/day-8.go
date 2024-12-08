package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func stringsToIntegers(lines []string) ([]int, error) {
	integers := make([]int, 0, len(lines))
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		integers = append(integers, n)
	}
	return integers, nil
}

func isAlphanumeric(c rune) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func partOne() {

	wd, _ := os.Getwd()
	filePath := wd + "/src/days/day-8/input.txt"
	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0

	antenaLocations := make(map[rune][][]int)
	matrix := [][]rune{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		runeArray := []rune(line)
		matrix = append(matrix, runeArray)
	}

	for y, _ := range matrix {
		for x, valX := range matrix[y] {
			if isAlphanumeric(valX) {
				antenaLocations[valX] = append(antenaLocations[valX], []int{y, x})
			}
		}
	}

	fmt.Println("antenals: ", antenaLocations)

	for antenaFrequency := range antenaLocations {

		antenas, ok := antenaLocations[antenaFrequency]

		if ok {

			for i := 0; i < len(antenas)-1; i++ {

				antenaOne := antenas[i]

				for j := i + 1; j < len(antenas); j++ {
					antenaTwo := antenas[j]

					fmt.Println("For freq: ", string(antenaFrequency), "locations: ", antenaOne, antenaTwo)

					vector := []int{antenaOne[0] - antenaTwo[0], antenaOne[1] - antenaTwo[1]}

					antinodeOneY := antenaOne[0] + vector[0]
					antinodeOneX := antenaOne[1] + vector[1]

					if antinodeOneY >= 0 && antinodeOneY < len(matrix) && antinodeOneX >= 0 && antinodeOneX < len(matrix[0]) &&
						matrix[antinodeOneY][antinodeOneX] != '#' {

						matrix[antinodeOneY][antinodeOneX] = '#'
						result++
					}

					antinodeOneY = antenaTwo[0] + vector[0]*-1
					antinodeOneX = antenaTwo[1] + vector[1]*-1

					if antinodeOneY >= 0 && antinodeOneY < len(matrix) && antinodeOneX >= 0 && antinodeOneX < len(matrix[0]) &&
						matrix[antinodeOneY][antinodeOneX] != '#' {

						matrix[antinodeOneY][antinodeOneX] = '#'
						result++
					}

				}
			}

		}

	}

	for y, _ := range matrix {
		for _, valX := range matrix[y] {
			fmt.Print(string(valX))
		}
		fmt.Println()
	}

	fmt.Println("result: ", result)

}

func partTwo() {
	wd, _ := os.Getwd()
	filePath := wd + "/src/days/day-8/input.txt"
	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0

	antenaLocations := make(map[rune][][]int)
	matrix := [][]rune{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		runeArray := []rune(line)
		matrix = append(matrix, runeArray)
	}

	for y, _ := range matrix {
		for x, valX := range matrix[y] {
			if isAlphanumeric(valX) {
				antenaLocations[valX] = append(antenaLocations[valX], []int{y, x})
			}
		}
	}

	fmt.Println("antenals: ", antenaLocations)

	for antenaFrequency := range antenaLocations {

		antenas, ok := antenaLocations[antenaFrequency]

		if ok {

			for i := 0; i < len(antenas)-1; i++ {

				antenaOne := antenas[i]

				for j := i + 1; j < len(antenas); j++ {
					antenaTwo := antenas[j]

					fmt.Println("For freq: ", string(antenaFrequency), "locations: ", antenaOne, antenaTwo)

					vector := []int{antenaOne[0] - antenaTwo[0], antenaOne[1] - antenaTwo[1]}

					antinodeOneY := antenaOne[0] + vector[0]
					antinodeOneX := antenaOne[1] + vector[1]

					matrix[antenaOne[0]][antenaOne[1]] = '#'
					matrix[antenaTwo[0]][antenaTwo[1]] = '#'

					for antinodeOneY >= 0 && antinodeOneY < len(matrix) && antinodeOneX >= 0 && antinodeOneX < len(matrix[0]) {

						matrix[antinodeOneY][antinodeOneX] = '#'
						antinodeOneY += vector[0]
						antinodeOneX += vector[1]
					}

					antinodeOneY = antenaTwo[0] + vector[0]*-1
					antinodeOneX = antenaTwo[1] + vector[1]*-1

					for antinodeOneY >= 0 && antinodeOneY < len(matrix) && antinodeOneX >= 0 && antinodeOneX < len(matrix[0]) {
						matrix[antinodeOneY][antinodeOneX] = '#'
						antinodeOneY += vector[0] * -1
						antinodeOneX += vector[1] * -1
					}

				}
			}

		}

	}

	for y, _ := range matrix {
		for x, valX := range matrix[y] {
			if matrix[y][x] == '#' {
				result++
			}
			fmt.Print(string(valX))
		}
		fmt.Println()
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
