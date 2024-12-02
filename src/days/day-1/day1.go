package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func partOne() {

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}

	filePath := wd + "/src/days/day-1/input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Can't open the file %s: %v\n", filePath, err)
		return
	}
	defer file.Close()

	var leftArr []int64
	var rightArr []int64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		splits := strings.Fields(line)
		if len(splits) != 2 {
			fmt.Printf("Invalid line format: %s\n", line)
			continue
		}

		leftVal, err := strconv.ParseInt(splits[0], 10, 64)
		if err != nil {
			fmt.Printf("Error parsing left value in line '%s': %v\n", line, err)
			continue
		}

		rightVal, err := strconv.ParseInt(splits[1], 10, 64)
		if err != nil {
			fmt.Printf("Error parsing right value in line '%s': %v\n", line, err)
			continue
		}

		leftArr = append(leftArr, leftVal)
		rightArr = append(rightArr, rightVal)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
	}

	sort.Slice(leftArr, func(i, j int) bool { return leftArr[i] < leftArr[j] })
	sort.Slice(rightArr, func(i, j int) bool { return rightArr[i] < rightArr[j] })

	// fmt.Println("Left Array:", leftArr)
	// fmt.Println("Right Array:", rightArr)

	var result int64

	for index, left := range leftArr {
		var right = rightArr[index]

		if left < right {
			result += right - left
		} else {
			result += left - right
		}
	}

	fmt.Println("result: ", result)

}

func partTwo() {
	var wd, err = os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}

	filePath := wd + "/src/days/day-1/input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Can't open the file %s: %v\n", filePath, err)
		return
	}
	defer file.Close()

	var leftArr []int64
	var rightMap = map[int64]int64{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		splits := strings.Fields(line)
		if len(splits) != 2 {
			fmt.Printf("Invalid line format: %s\n", line)
			continue
		}

		leftVal, err := strconv.ParseInt(splits[0], 10, 64)
		if err != nil {
			fmt.Printf("Error parsing left value in line '%s': %v\n", line, err)
			continue
		}

		rightVal, err := strconv.ParseInt(splits[1], 10, 64)
		if err != nil {
			fmt.Printf("Error parsing right value in line '%s': %v\n", line, err)
			continue
		}

		leftArr = append(leftArr, leftVal)

		var count, ok = rightMap[rightVal]
		if ok {
			rightMap[rightVal] = count + 1
		} else {
			rightMap[rightVal] = 1
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
	}

	var result int64

	for _, left := range leftArr {
		result += left * rightMap[left]
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
