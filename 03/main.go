package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	report := strings.Split(string(content), "\n")
	report = report[:len(report)-1]

	fmt.Println("# Power Usage", computePowerUsage(report))
	fmt.Println("# Life support rating", computeOxygenGeneratorRating(report)*computeCO2ScubberRating(report))
}

func processDiagnosticReport(report []string) (int, int) {
	onesCountByColumn := make([]int, len(report[0]))
	for _, line := range report {
		for index, number := range line {
			if number == '1' {
				onesCountByColumn[index]++
			}
		}
	}

	halfLinesCount := len(report) / 2
	gammaStr := ""
	epsilonStr := ""
	for _, onesCount := range onesCountByColumn {
		if onesCount > halfLinesCount {
			gammaStr += "1"
			epsilonStr += "0"
		} else {
			gammaStr += "0"
			epsilonStr += "1"
		}
	}
	gamma, err := strconv.ParseInt(gammaStr, 2, 0)
	if err != nil {
		panic(err)
	}
	epsilon, err := strconv.ParseInt(epsilonStr, 2, 0)
	if err != nil {
		panic(err)
	}

	return int(gamma), int(epsilon)
}

func computePowerUsage(report []string) int {
	gamma, epsilon := processDiagnosticReport(report)
	return gamma * epsilon
}

func computeOxygenGeneratorRating(report []string) int {
	matchingLines := report

	// for each column, process each line then restrict them
	columnsCount := len(report[0])
	for columnIndex := 0; columnIndex < columnsCount; columnIndex++ {
		onesCount := 0
		for _, line := range matchingLines {
			if line[columnIndex] == '1' {
				onesCount++
			}
		}

		mostCommonValue := byte('0')
		if onesCount >= len(matchingLines)-onesCount {
			mostCommonValue = '1'
		}

		newMatchingLines := []string{}
		for _, line := range matchingLines {
			if line[columnIndex] == mostCommonValue {
				newMatchingLines = append(newMatchingLines, line)
			}
		}

		if len(newMatchingLines) == 1 {
			oxygen, err := strconv.ParseInt(newMatchingLines[0], 2, 0)
			if err != nil {
				panic(err)
			}
			return int(oxygen)
		}
		matchingLines = newMatchingLines
	}
	// should never be executed
	return 0
}

func computeCO2ScubberRating(report []string) int {
	matchingLines := report

	// for each column, process each line then restrict them
	columnsCount := len(report[0])
	for columnIndex := 0; columnIndex < columnsCount; columnIndex++ {
		onesCount := 0
		for _, line := range matchingLines {
			if line[columnIndex] == '1' {
				onesCount++
			}
		}

		mostCommonValue := byte('1')
		if onesCount >= len(matchingLines)-onesCount {
			mostCommonValue = '0'
		}

		newMatchingLines := []string{}
		for _, line := range matchingLines {
			if line[columnIndex] == mostCommonValue {
				newMatchingLines = append(newMatchingLines, line)
			}
		}

		if len(newMatchingLines) == 1 {
			oxygen, err := strconv.ParseInt(newMatchingLines[0], 2, 0)
			if err != nil {
				panic(err)
			}
			return int(oxygen)
		}
		matchingLines = newMatchingLines
	}
	// should never be executed
	return 0
}
