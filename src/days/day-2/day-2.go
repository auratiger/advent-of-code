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

func partOne() {
	wd, _ := os.Getwd()
	filePath := wd + "/src/days/day-2/input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Can't open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	safeReports := 0

	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		splits := strings.Fields(line)
		intSplits, err := stringsToIntegers(splits)
		if err != nil {
			fmt.Println("Err converting splits")
		}

		// fmt.Println("line", splits)

		isValid := levelIsSafe(intSplits)

		if isValid {
			safeReports += 1

		}

	}

	fmt.Println("Safe reporst: ", safeReports)

}

func levelIsSafe(levels []int) bool {
	isIncreasing := true
	isSafe := true
	badLevel := 0

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

		if diff == 0 {
			badLevel++
			isSafe = false
			break
		}

		if diff < -3 || diff > 3 {
			isSafe = false
			break
		}

		if i == 1 && diff < 0 {
			isIncreasing = false
		} else if (isIncreasing && diff < 0) || (!isIncreasing && diff > 0) {
			isSafe = false
			break
		}
	}

	return isSafe
}

func partTwo() {
	wd, _ := os.Getwd()
	filePath := wd + "/src/days/day-2/input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Can't open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	safeReports := 0

	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		splits := strings.Fields(line)
		intSplits, err := stringsToIntegers(splits)
		if err != nil {
			fmt.Println("Err converting splits")
		}

		isSafe := levelIsSafe(intSplits)

		if !isSafe {
			for i := 0; i < len(intSplits); i++ {
				modifiedLevels := []int{}
				modifiedLevels = append(modifiedLevels, intSplits[:i]...)
				modifiedLevels = append(modifiedLevels, intSplits[i+1:]...)

				isSafe = levelIsSafe(modifiedLevels)

				if isSafe {
					break
				}
			}

		}

		if isSafe {
			safeReports += 1
		}

	}

	fmt.Println("Safe reporst: ", safeReports)

}

func main() {
	fmt.Println("Running")

	var start = time.Now()

	partOne()
	// partTwo()

	var end = time.Since(start)

	fmt.Println("Execution time: ", end.Seconds())

}
