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
