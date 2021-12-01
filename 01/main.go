package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	report := []int{}
	for _, line := range lines {
		number, err := strconv.Atoi(line)
		if err == nil {
			report = append(report, number)
		}
	}

	fmt.Println("# increases", countIncreases(report))
	fmt.Println("# sliding window increases", countSlidingWindowIncreases(report))
}

func countIncreases(report []int) int {
	increasesCount := 0
	previousDepth := report[0]
	for _, depth := range report[1:] {
		if depth > previousDepth {
			increasesCount++
		}
		previousDepth = depth
	}
	return increasesCount
}

func countSlidingWindowIncreases(report []int) int {
	increasesCount := 0
	previousSum := report[0] + report[1] + report[2]
	for i := 1; i < len(report)-2; i++ {
		sum := report[i] + report[i+1] + report[i+2]
		if sum > previousSum {
			increasesCount++
		}
		previousSum = sum
	}
	return increasesCount
}
