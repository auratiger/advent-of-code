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

func fixArr(lineArr *[]string, i int, matchArr *[]int) {

	temp, _ := strconv.Atoi((*lineArr)[i])
	for index := i - 1; index >= 0; index-- {
		val, _ := strconv.Atoi((*lineArr)[index])

		if slices.Contains(*matchArr, val) {
			(*lineArr)[index] = strconv.Itoa(temp)
			(*lineArr)[index+1] = strconv.Itoa(val)
		}

	}

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

		for i, val := range lineArr {

			parsedVal, _ := strconv.Atoi(val)
			matchArr, ok := m[parsedVal]

			if ok {
				for _, xVal := range matchArr {
					checkString := lineArr[:i]

					parsedXVal := strconv.Itoa(xVal)

					if slices.Contains(checkString, parsedXVal) {
						// fmt.Println("fail: ", lineArr, parsedVal)
						fixArr(&lineArr, i, &matchArr)
						hasFailed = true
						break
					}
				}
			}

		}

		// fmt.Println("wha happen: ", lineArr)

		if hasFailed {
			temp := (len(lineArr)) / 2
			// fmt.Println("temp: ", temp, lineArr)
			val, _ := strconv.Atoi(lineArr[temp])
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
