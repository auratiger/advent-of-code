package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
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

	fmt.Println(str)

	r, _ := regexp.Compile("mul\\((\\d{1,3}),(\\d{1,3})\\)")

	res := r.FindAllStringSubmatch(str, 1000)

	fmt.Println(res)

	for _, val := range res {

		left, _ := strconv.ParseUint(val[1], 10, 64)
		right, _ := strconv.ParseUint(val[2], 10, 64)

		result += left * right
	}

	fmt.Println("Safe reporst: ", result)

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
