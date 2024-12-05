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

func partOne() {

	wd, _ := os.Getwd()
	filePath := wd + "/src/days/day-5/input.txt"
	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	m := map[int][]int{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			break
		}

		split := strings.Split(line, "|")

		left, _ := strconv.Atoi(split[0])
		right, _ := strconv.Atoi(split[1])

		m[left] = append(m[left], right)
	}

	result := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lineArr := strings.Split(line, ",")

		hasFailed := false

	Loop:
		for i, val := range lineArr {

			parsedVal, _ := strconv.Atoi(val)
			matchArr, ok := m[parsedVal]

			if ok {
				for _, xVal := range matchArr {
					checkString := lineArr[:i]

					parsedXVal := strconv.Itoa(xVal)

					if slices.Contains(checkString, parsedXVal) {
						hasFailed = true
						break Loop
					}
				}
			}

		}

		if !hasFailed {
			val, _ := strconv.Atoi(lineArr[len(lineArr)/2])

			result += val
		}

	}

	// for k, v := range m {
	// 	fmt.Println(k, "value is", v)
	// }

	fmt.Println("result: ", result)

}

func fixArr(lineArr *[]int, i int, matchArr *[]int) {

	temp := (*lineArr)[i]

	(*lineArr)[i] = (*lineArr)[i-1]
	(*lineArr)[i-1] = temp

	for index := i - 2; index >= 0; index-- {
		val := (*lineArr)[index]

		if slices.Contains(*matchArr, val) {
			(*lineArr)[index] = temp
			(*lineArr)[index+1] = val
		}

	}

}

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

func partTwo() {

	wd, _ := os.Getwd()
	filePath := wd + "/src/days/day-5/input.txt"
	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	m := map[int][]int{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			break
		}

		split := strings.Split(line, "|")

		left, _ := strconv.Atoi(split[0])
		right, _ := strconv.Atoi(split[1])

		m[left] = append(m[left], right)
	}

	result := 0
	hasFailed := false

	for scanner.Scan() {
		hasFailed = false
		line := strings.TrimSpace(scanner.Text())
		lineArr := strings.Split(line, ",")
		parsedLineArr, err := stringsToIntegers(lineArr)

		if err != nil {
			fmt.Println("parsing the array to integers has failed: ", err)
		}

		for i, val := range parsedLineArr {

			matchArr, ok := m[val]

			if ok {
				for _, xVal := range matchArr {
					checkRules := parsedLineArr[:i]

					if slices.Contains(checkRules, xVal) {
						fixArr(&parsedLineArr, i, &matchArr)
						hasFailed = true
						break
					}
				}
			}

		}

		if hasFailed {
			temp := (len(parsedLineArr)) / 2
			val := parsedLineArr[temp]
			result += val
		}

	}

	// for k, v := range m {
	// 	fmt.Println(k, "value is", v)
	// }

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
