package main

import (
	"fmt"
	"io/ioutil"
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
	graph := parseInput(input)
	paths := [][]string{}
	NewPath().VisitCave(graph.Start, &paths)
	return len(paths)
}

type Graph struct {
	Start *Cave
	End   *Cave

	Caves map[string]*Cave
}

type Cave struct {
	ID        string
	Big       bool
	Neighbors []*Cave
}

func NewCave(ID string) *Cave {
	return &Cave{
		ID:        ID,
		Big:       ID[0] >= 'A' && ID[0] <= 'Z',
		Neighbors: []*Cave{},
	}
}

func parseInput(input []string) Graph {
	caves := map[string]*Cave{}
	for _, line := range input {
		parts := strings.Split(line, "-")
		var existingCave0, existingCave1 *Cave
		var exists bool
		if existingCave0, exists = caves[parts[0]]; !exists {
			existingCave0 = NewCave(parts[0])
			caves[parts[0]] = existingCave0
		}
		if existingCave1, exists = caves[parts[1]]; !exists {
			existingCave1 = NewCave(parts[1])
			caves[parts[1]] = existingCave1
		}
		existingCave0.Neighbors = append(existingCave0.Neighbors, existingCave1)
		existingCave1.Neighbors = append(existingCave1.Neighbors, existingCave0)
	}

	startCave := caves["start"]
	endCave := caves["start"]
	return Graph{
		Start: startCave,
		End:   endCave,
		Caves: caves,
	}
}

type Path struct {
	VisitedCaves []string
}

func NewPath() *Path {
	return &Path{
		VisitedCaves: []string{},
	}
}

func (path *Path) Clone() *Path {
	return &Path{
		VisitedCaves: path.VisitedCaves, // careful here as the slice data is not copied. However it works fine oO
	}
}

// we never have 2 big caves connected
func (path *Path) Contains(cave *Cave) bool {
	for _, caveId := range path.VisitedCaves {
		if caveId == cave.ID {
			return true
		}
	}
	return false
}

func (path *Path) VisitCave(cave *Cave, completePaths *[][]string) {
	path.VisitedCaves = append(path.VisitedCaves, cave.ID)

	if cave.ID == "end" {
		*completePaths = append(*completePaths, path.VisitedCaves)
		return
	}

	for _, neighbor := range cave.Neighbors {
		if neighbor.Big || !path.Contains(neighbor) {
			path.Clone().VisitCave(neighbor, completePaths) // copy self path
		}
	}
}

func part2(input []string) int {
	graph := parseInput(input)
	paths := [][]string{}
	NewPart2Path().VisitCave(graph.Start, &paths)

	// deduplicate paths
	set := map[string]bool{}
	for _, path := range paths {
		key := strings.Join(path, ",")
		set[key] = true
	}
	return len(set)
}

type Part2Path struct {
	VisitedCaves []string

	SmallCaveVisitID    string
	SmallCaveVisitCount int
}

// we never have 2 big caves connected
func NewPart2Path() *Part2Path {
	return &Part2Path{
		VisitedCaves:        []string{},
		SmallCaveVisitID:    "",
		SmallCaveVisitCount: 0,
	}
}

// we never have 2 big caves connected
func (path *Part2Path) Clone() *Part2Path {
	VisitedCavesCopy := make([]string, len(path.VisitedCaves))
	copy(VisitedCavesCopy, path.VisitedCaves)
	return &Part2Path{
		VisitedCaves:        VisitedCavesCopy,
		SmallCaveVisitID:    path.SmallCaveVisitID,
		SmallCaveVisitCount: path.SmallCaveVisitCount,
	}
}

func (path *Part2Path) Contains(cave *Cave) bool {
	for _, caveId := range path.VisitedCaves {
		if caveId == cave.ID {
			return true
		}
	}
	return false
}

func (path *Part2Path) VisitCave(cave *Cave, completePaths *[][]string) {
	path.VisitedCaves = append(path.VisitedCaves, cave.ID)

	if cave.ID == "end" {
		*completePaths = append(*completePaths, path.VisitedCaves)
		return
	}

	for _, neighbor := range cave.Neighbors {
		if neighbor.ID == "start" {
			continue
		}
		if neighbor.Big || neighbor.ID == "end" {
			path.Clone().VisitCave(neighbor, completePaths)
			continue
		}

		// small cave, not start, not end

		if path.Contains(neighbor) {
			// the neighbor may be the elected small cave that can be visited twice
			if path.SmallCaveVisitID == neighbor.ID && path.SmallCaveVisitCount == 1 {
				p := path.Clone()
				p.SmallCaveVisitCount++
				p.VisitCave(neighbor, completePaths)
			}
		} else {
			// the small cave can be visited for the first time
			path.Clone().VisitCave(neighbor, completePaths)

			// if no small cave elected for two visits yet, elect the neighbor
			if path.SmallCaveVisitID == "" {
				p := path.Clone()
				p.SmallCaveVisitID = neighbor.ID
				p.SmallCaveVisitCount = 1
				p.VisitCave(neighbor, completePaths)
			}
		}
	}
}
