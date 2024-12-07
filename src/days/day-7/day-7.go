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

func operate(a, b int, operator string) int {
	if operator == "*" {
		return a * b
	} else {
		return a + b
	}
}

func recurs(nums *[]int, a, index, end int, operator string, target int) int {

	if index == end {
		return operate(a, (*nums)[index], operator)
	}

	res1 := recurs(nums, operate(a, (*nums)[index], operator), index+1, end, "*", target)
	if res1 == target {
		return res1
	}

	res2 := recurs(nums, operate(a, (*nums)[index], operator), index+1, end, "+", target)
	if res2 == target {
		return res2
	}

	return -1
}

func parse(nums *[]int, target int) bool {

	res1 := recurs(nums, (*nums)[0], 1, len(*nums)-1, "*", target)
	if res1 != -1 && res1 == target {
		return true
	}

	res2 := recurs(nums, (*nums)[0], 1, len(*nums)-1, "+", target)
	if res2 != -1 && res2 == target {
		return true
	}

	return false
}

func partOne() {

	wd, _ := os.Getwd()
	filePath := wd + "/src/days/day-7/input.txt"
	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		split := strings.Split(line, ":")

		target, _ := strconv.Atoi(split[0])
		intArr, _ := stringsToIntegers(strings.Split(strings.TrimSpace(split[1]), " "))

		gotValue := parse(&intArr, target)

		if gotValue {
			// fmt.Println("here", target)
			result += target
		}
	}

	fmt.Println("result: ", result)

}

func operate2(a, b int, operator string) int {
	if operator == "*" {
		return a * b
	} else if operator == "+" {
		return a + b
	} else {
		res, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
		return res
	}
}

func recurs2(nums *[]int, a, index, end int, operator string, target int) int {

	if index == end {
		return operate2(a, (*nums)[index], operator)
	}

	res1 := recurs2(nums, operate2(a, (*nums)[index], operator), index+1, end, "*", target)
	if res1 == target {
		return res1
	}

	res2 := recurs2(nums, operate2(a, (*nums)[index], operator), index+1, end, "+", target)
	if res2 == target {
		return res2
	}

	res3 := recurs2(nums, operate2(a, (*nums)[index], operator), index+1, end, "||", target)
	if res3 == target {
		return res3
	}

	return -1
}

func parse2(nums *[]int, target int) bool {

	res1 := recurs2(nums, (*nums)[0], 1, len(*nums)-1, "*", target)
	if res1 != -1 && res1 == target {
		return true
	}

	res2 := recurs2(nums, (*nums)[0], 1, len(*nums)-1, "+", target)
	if res2 != -1 && res2 == target {
		return true
	}

	res3 := recurs2(nums, (*nums)[0], 1, len(*nums)-1, "||", target)
	if res3 != -1 && res3 == target {
		return true
	}

	return false
}

func partTwo() {
	wd, _ := os.Getwd()
	filePath := wd + "/src/days/day-7/input.txt"
	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		split := strings.Split(line, ":")

		target, _ := strconv.Atoi(split[0])
		intArr, _ := stringsToIntegers(strings.Split(strings.TrimSpace(split[1]), " "))

		gotValue := parse2(&intArr, target)

		if gotValue {
			// fmt.Println("here", target)
			result += target
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
