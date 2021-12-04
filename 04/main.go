package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	board := strings.Split(string(content), "\n")
	board = board[:len(board)-1]

	fmt.Println("# Winning bingo score", getWinningBingoScore(board))
	fmt.Println("# Losing bingo score", getLosingBingoScore(board))
}

func getWinningBingoScore(game []string) int {
	calledNumbers, boards := buildGameData(game)

	// for each called number, mark it on each board
	for _, calledNumber := range calledNumbers {
		fmt.Println("Calling number", calledNumber)
		for boardIndex, board := range boards {
			if board.MarkNumber(calledNumber) {
				fmt.Println("Board", boardIndex, "wins")
				return board.GetUnmarkedCellsSum() * calledNumber
			}
		}
	}

	panic("A board should have won already...")
}

func getLosingBingoScore(game []string) int {
	calledNumbers, boards := buildGameData(game)
	boardsCount := len(boards)
	winningBoardsCount := 0
	boardHasWon := make([]bool, boardsCount)

	// for each called number, mark it on each board
	for _, calledNumber := range calledNumbers {
		fmt.Println("Calling number", calledNumber)
		for boardIndex, board := range boards {
			if boardHasWon[boardIndex] {
				continue
			}
			if board.MarkNumber(calledNumber) {
				fmt.Println("Board", boardIndex, "wins")
				// mark the board as inactive
				boardHasWon[boardIndex] = true
				winningBoardsCount++
				if winningBoardsCount == boardsCount {
					fmt.Println("Board", boardIndex, "wins (last)")
					return board.GetUnmarkedCellsSum() * calledNumber
				}
			}
		}
	}

	panic("A board should have won already...")
}

type Board struct {
	Cells [][]Cell
}

// Returns true if the called number make the board win
func (board *Board) MarkNumber(calledNumber int) bool {
	for lineIndex, lineCells := range board.Cells {
		for columnIndex, cell := range lineCells {
			if cell.Number == calledNumber {
				fmt.Println("> Marking", calledNumber)
				lineCells[columnIndex].Marked = true

				if board.isWinningCell(lineIndex, columnIndex) {
					return true
				}
			}
		}
	}
	return false
}

func (board *Board) GetUnmarkedCellsSum() int {
	sum := 0
	for _, lineCells := range board.Cells {
		for _, cell := range lineCells {
			if !(&cell).Marked {
				sum += cell.Number
			}
		}
	}
	return sum
}

func (board *Board) isWinningCell(lineIndex int, columnIndex int) bool {
	horizontalMarkedCellsCount := 0
	verticalMarkedCellsCount := 0
	for i := 0; i < len(board.Cells); i++ {
		if board.Cells[lineIndex][i].Marked {
			horizontalMarkedCellsCount++
		}
		if board.Cells[i][columnIndex].Marked {
			verticalMarkedCellsCount++
		}
	}
	return horizontalMarkedCellsCount == 5 || verticalMarkedCellsCount == 5
}

type Cell struct {
	Number int
	Marked bool
}

func buildGameData(game []string) ([]int, []Board) {
	// extract the called numbers (first line)
	calledNumbers := []int{}
	for _, numberStr := range strings.Split(game[0], ",") {
		number, _ := strconv.Atoi(numberStr)
		calledNumbers = append(calledNumbers, number)
	}

	// extract each board
	boards := []Board{}
	for i := 2; i < len(game); i += 6 {
		board := Board{
			Cells: [][]Cell{},
		}
		board.Cells = append(board.Cells, convertBoardLineToCells(game[i]))
		board.Cells = append(board.Cells, convertBoardLineToCells(game[i+1]))
		board.Cells = append(board.Cells, convertBoardLineToCells(game[i+2]))
		board.Cells = append(board.Cells, convertBoardLineToCells(game[i+3]))
		board.Cells = append(board.Cells, convertBoardLineToCells(game[i+4]))
		boards = append(boards, board)
	}

	return calledNumbers, boards
}

func convertBoardLineToCells(line string) []Cell {
	numbers := []Cell{}
	for _, numberStr := range strings.Split(line, " ") {
		if numberStr != "" {
			number, _ := strconv.Atoi(numberStr)
			numbers = append(numbers, Cell{
				Number: number,
			})
		}
	}
	return numbers
}
