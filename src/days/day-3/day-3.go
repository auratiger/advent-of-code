package main

import (
	"fmt"
	"os"
	"regexp"
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
	filePath := wd + "/src/days/day-3/input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Can't open file")
	}
	defer file.Close()

	scanner, _ := os.ReadFile(filePath)
	str := string(scanner)

	var result uint64 = 0

	r, _ := regexp.Compile("mul\\((\\d{1,3}),(\\d{1,3})\\)")

	res := r.FindAllStringSubmatch(str, 1000)

	// fmt.Println(res)

	for _, val := range res {

		left, _ := strconv.ParseUint(val[1], 10, 64)
		right, _ := strconv.ParseUint(val[2], 10, 64)

		result += left * right
	}

	fmt.Println("Safe reporst: ", result)

}

func partTwo() {
	wd, _ := os.Getwd()
	filePath := wd + "/src/days/day-3/input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Can't open file")
	}
	defer file.Close()

	scanner, _ := os.ReadFile(filePath)
	str := string(scanner)

	var result uint64 = 0

	// fmt.Println(str)

	r, _ := regexp.Compile("mul\\((\\d{1,3}),(\\d{1,3})\\)|do\\(\\)|don't\\(\\)")

	res := r.FindAllStringSubmatch(str, 1000)

	// fmt.Println(res)

	isEnabled := true

	for _, val := range res {

		symbol := val[0]

		if symbol == "do()" {
			isEnabled = true
		} else if symbol == "don't()" {
			isEnabled = false
		} else {
			if isEnabled {
				left, _ := strconv.ParseUint(val[1], 10, 64)
				right, _ := strconv.ParseUint(val[2], 10, 64)

				result += left * right
			}
		}

	}

	fmt.Println("Safe reporst: ", result)

}

func partOneNoRegex() {

	wd, _ := os.Getwd()
	filePath := wd + "/src/days/day-3/input.txt"
	content, _ := os.ReadFile(filePath)
	str := string(content)

	split := strings.Split(str, "mul")

	result := 0

	for _, val := range split {
		if val[0] != '(' {
			continue
		}

		hasClosingBracket := strings.IndexByte(val, ')')

		if hasClosingBracket == -1 {
			continue
		}

		vars := val[:hasClosingBracket]
		varsSplit := strings.Split(vars[1:], ",")

		if len(varsSplit) != 2 {
			// fmt.Println("split not 2", varsSplit)
			continue
		}

		leftVar, err := strconv.Atoi(varsSplit[0])

		if err != nil {
			// fmt.Println("left not int: ", leftVar)
			continue
		}

		rigthVar, err := strconv.Atoi(varsSplit[1])

		if err != nil {
			// fmt.Println("rigth not int: ", rigthVar)
			continue
		}

		result += leftVar * rigthVar

	}

	fmt.Println("result: ", result)

}

func partTwoNoRegex() {

	wd, _ := os.Getwd()
	filePath := wd + "/src/days/day-3/input.txt"
	content, _ := os.ReadFile(filePath)
	str := string(content)

	split := strings.Split(str, "mul")

	result := 0
	active := true
	bufferedActive := true

	for _, val := range split {
		active = bufferedActive

		if active {
			hasDont := strings.Contains(val, "don't()")

			if hasDont {
				bufferedActive = false
			}
		} else {
			hasDo := strings.Contains(val, "do()")

			if hasDo {
				bufferedActive = true
			}

		}

		if !active {
			continue
		}

		if val[0] != '(' {
			continue
		}

		hasClosingBracket := strings.IndexByte(val, ')')
		if hasClosingBracket == -1 {
			continue
		}

		vars := val[1:hasClosingBracket]
		varsSplit := strings.Split(vars, ",")

		if len(varsSplit) != 2 {
			// fmt.Println("split not 2", varsSplit)
			continue
		}

		leftVar, err := strconv.Atoi(varsSplit[0])

		if err != nil {
			// fmt.Println("left not int: ", leftVar)
			continue
		}

		rigthVar, err := strconv.Atoi(varsSplit[1])

		if err != nil {
			// fmt.Println("rigth not int: ", rigthVar)
			continue
		}

		result += leftVar * rigthVar

	}

	fmt.Println("result: ", result)

}

func main() {
	fmt.Println("Running")

	var start = time.Now()

	// partOne()
	// partTwo()
	// partOneNoRegex()
	partTwoNoRegex()

	var end = time.Since(start)

	fmt.Println("Execution time: ", end.Seconds())

}
