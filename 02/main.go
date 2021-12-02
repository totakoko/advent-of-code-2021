package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// part 1
type Submarine struct {
	x int
	y int
}

func (s *Submarine) Move(direction string, count int) {
	switch direction {
	case "forward":
		s.x += count
	case "up":
		s.y -= count
	case "down":
		s.y += count
	}
}

func extractMovementFromText(str string) (string, int) {
	parts := strings.Split(str, " ")
	count, _ := strconv.Atoi(parts[1])
	return parts[0], count
}

// part 2
type ComplexSubmarine struct {
	x   int
	y   int
	aim int
}

func (s *ComplexSubmarine) Move(direction string, count int) {
	switch direction {
	case "forward":
		s.x += count
		s.y += s.aim * count
	case "up":
		s.aim -= count
	case "down":
		s.aim += count
	}
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")

	submarine := &Submarine{}
	complexSubmarine := &ComplexSubmarine{}
	for _, command := range lines {
		if command != "" {
			direction, count := extractMovementFromText(command)
			submarine.Move(direction, count)
			complexSubmarine.Move(direction, count)
		}
	}
	fmt.Println("# result", submarine.x*submarine.y)
	fmt.Println("# result", complexSubmarine.x*complexSubmarine.y)
}
