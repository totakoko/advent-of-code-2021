package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	input := strings.Split(string(content), "\n")
	input = input[:len(input)-1]

	fmt.Println("# Part 1", part1(input))
	fmt.Println("# Part 2", part2(input))
}

func part1(input []string) int {
	uniqueNumbersCount := 0
	for _, line := range input {
		parts := strings.Split(line, " | ")
		notes := strings.Split(parts[1], " ")

		for _, signal := range notes {
			switch len(signal) {
			case 2, 3, 4, 7: // 1 7 4 8
				uniqueNumbersCount++
			}
		}
	}
	return uniqueNumbersCount
}

func part2(input []string) int {
	uniqueNumbersCount := 0
	for _, line := range input {
		uniqueNumbersCount += part2line(line)
	}
	return uniqueNumbersCount
}

func part2line(line string) int {
	parts := strings.Split(line, " | ")
	digits := strings.Split(parts[0], " ")

	one := findN(digits, 2)
	four := findN(digits, 4)
	seven := findN(digits, 3)
	eight := findN(digits, 7)

	fmt.Println("line", line)
	fmt.Println("1", one)
	fmt.Println("4", four)
	fmt.Println("7", seven)
	fmt.Println("8", eight)

	for _, digit := range digits {
		if !includesAllLetters(digit, one) {
			fmt.Println("segment could be 2 5 6", digit)
		}
	}

	// = len(6) and not all letters from 1
	six := findMatching(digits, func(digit string) bool {
		return len(digit) == 6 && !includesAllLetters(digit, one)
	})
	fmt.Println("6", six)

	for _, digit := range digits {
		if len(digit) == 6 {
			fmt.Println("segment could be 0 6 9", digit)
		}
	}

	nine := findMatching(digits, func(digit string) bool {
		return len(digit) == 6 && includesAllLetters(digit, four)
	})
	fmt.Println("9", nine)

	zero := findMatching(digits, func(digit string) bool {
		return len(digit) == 6 && digit != nine && digit != six
	})
	fmt.Println("0", zero)

	three := findMatching(digits, func(digit string) bool {
		return len(digit) == 5 && includesAllLetters(digit, seven)
	})
	fmt.Println("3", three)

	// 5 shares 3 segments with 4
	five := findMatching(digits, func(digit string) bool {
		count := 0
		for _, letter := range four {
			if bytes.IndexByte([]byte(digit), byte(letter)) != -1 {
				count++
			}
		}
		return len(digit) == 5 && count == 3 && digit != three
	})
	fmt.Println("5", five)

	// 2 shares 2 segments with 4
	two := findMatching(digits, func(digit string) bool {
		count := 0
		for _, letter := range four {
			if bytes.IndexByte([]byte(digit), byte(letter)) != -1 {
				count++
			}
		}
		return len(digit) == 5 && count == 2
	})
	fmt.Println("2", two)

	decodedDigitSignals := []string{zero, one, two, three, four, five, six, seven, eight, nine}

	outputNumberString := ""
	notes := strings.Split(parts[1], " ")
	for _, note := range notes {
		for i, digit := range decodedDigitSignals {
			if len(note) == len(digit) && includesAllLetters(note, digit) {
				fmt.Printf("Found matching signal at %d : %s  (note = %s)\n", i, digit, note)
				outputNumberString += strconv.Itoa(i)
				break
			}
		}
	}

	outputNumber, _ := strconv.Atoi(outputNumberString)
	return outputNumber
}

func findN(digits []string, size int) string {
	return findMatching(digits, func(digit string) bool {
		return len(digit) == size
	})
}

func includesAllLetters(digit string, letters string) bool {
	for _, letter := range []byte(letters) {
		if bytes.IndexByte([]byte(digit), letter) == -1 {
			return false
		}
	}
	return true
}

func findMatching(digits []string, filter func(digit string) bool) string {
	for _, digit := range digits {
		if filter(digit) {
			return digit
		}
	}
	return ""
}
