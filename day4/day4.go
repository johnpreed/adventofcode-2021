package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// struct of position defined as number and whether it is marked or not
type position struct {
	number int
	marked bool
}

// bingoboard struct
type bingoboard struct {
	board [][]position
}

func main() {
	// read input file
	lines := readFile("input.txt")

	// create bingo game input from lines, the first line is an array of ints that represents the numbers to be used to mark game boards. the 5 lines after each blank line represents a new bingo game board. each bingo game board is a 5x5 grid of numbers.
	var bingoInput []int
	var bingoBoards []bingoboard
	for i := 0; i < len(lines); i++ {
		// if it is the first line fill in the bingo input from the comma delimited integers on line 1
		if i == 0 {
			bingoInput = getBingoInput(lines[i])
			continue
		} else if lines[i] == "" {
			// skip blank lines
			continue			
		}

		// create a new bingo board from the next 5 lines
		bingoBoards = append(bingoBoards, createBingoBoard(lines[i:i+5]))

		// set next position after bingo board
		i += 5
	}

	// create a map of winning bingo board indexes and boolean values
	winningBingoBoards := make(map[int]bool)

	// keep a count of winning boards
	winningBingoBoardCount := 0

	// foreach bingo input mark each bingo board, check if it's a bingo, and break if it is a bingo
	for _, bingoInput := range bingoInput {
		// print bingo input with label "drew number"
		fmt.Printf("drew number %d\n", bingoInput)
		
		for index, bingoBoard := range bingoBoards {
			// if board has already won, skip it. foreach winning boad, if the value in the array equals the index of the winning board, skip it
			if winningBingoBoards[index] {
				continue
			}
			
			isBingo := markBingoBoard(bingoBoard, bingoInput)
			if isBingo {
				// if board is a bingo, add it to the winning bingo boards map
				winningBingoBoards[index] = true
				winningBingoBoardCount++

				// calculate the score from sum the unmarked positions on the bingo board
				var score int
				for i := 0; i < len(bingoBoard.board); i++ {
					for j := 0; j < len(bingoBoard.board[i]); j++ {
						if !bingoBoard.board[i][j].marked {
							score += bingoBoard.board[i][j].number
						}
					}
				}
				score *= bingoInput
				
				// print the bingo board index with a winning board label
				fmt.Printf("BINGO! for board number: %d. score: %d\n", index+1, score)
				
				// stop when all bingo boards have won
				if winningBingoBoardCount == len(bingoBoards) {
					// print all boards have won with a counter
					fmt.Printf("all boards have won. total winning boards: %d\n", winningBingoBoardCount)
					return
				}
			}
		}
	}

}

// mark bingo board with bingo input and return true if board is a bingo
func markBingoBoard(bingoBoard bingoboard, bingoInput int) bool {
	// iterate through the rows and set the position to marked if it matches the bingo input
	for i := 0; i < len(bingoBoard.board); i++ {
		for j := 0; j < len(bingoBoard.board[i]); j++ {
			if bingoBoard.board[i][j].number == bingoInput {
				bingoBoard.board[i][j].marked = true
			}
		}
	}

	// check if bingo board is a bingo
	return isBingo(bingoBoard)
}

// get bingo input from comma delimited string of integers
func getBingoInput(line string) []int {
	var bingoInput []int
	for _, s := range strings.Split(line, ",") {
		i, _ := strconv.Atoi(s)
		bingoInput = append(bingoInput, i)
	}
	return bingoInput
}

// create bingo board from lines
func createBingoBoard(lines []string) bingoboard {
	var bingoBoard bingoboard
	for i := 0; i < len(lines); i++ {
		var row []position

		// trim the line of leading and trailing whitespace
		line := strings.TrimSpace(lines[i])

		// replace double spaces with single spaces
		line = strings.Replace(line, "  ", " ", -1)
		
		// split the line into integers, delimited by whitespace. whitespace can have multiple spaces between integers.
		for _, s := range strings.Split(line, " ") {
			i, _ := strconv.Atoi(s)
			row = append(row, position{i, false})
		}

		// add row to bingo board
		bingoBoard.board = append(bingoBoard.board, row)
	}

	// panic if the board is not 5x5 using array length of x and y
	if len(bingoBoard.board) != 5 {
		panic("board is not 5x5")
	}

	for i := 0; i < len(bingoBoard.board); i++ {
		if len(bingoBoard.board[i]) != 5 {
			panic("board is not 5x5")
		}
	}

	return bingoBoard
}

// check if a bingo board is a bingo
func isBingo(bingoBoard bingoboard) bool {
	// check rows
	for i := 0; i < len(bingoBoard.board); i++ {
		if isRowBingo(bingoBoard.board[i]) {
			// print found bingo row
			fmt.Printf("found bingo row %d\n", i+1)
			return true
		}
	}

	// check columns
	for i := 0; i < len(bingoBoard.board[0]); i++ {
		if isColumnBingo(bingoBoard.board, i) {
			// print found bingo column
			fmt.Printf("found bingo column %d\n", i+1)
			return true
		}
	}

	// check diagonals
	// if isDiagonalBingo(bingoBoard.board) {
	// 	return true
	// }

	return false
}

// check if a row is a bingo
func isRowBingo(row []position) bool {
	for i := 0; i < len(row); i++ {
		if !row[i].marked {
			return false
		}
	}
	return true
}

// check if a column is a bingo
func isColumnBingo(board [][]position, column int) bool {
	for i := 0; i < len(board); i++ {
		if !board[i][column].marked {
			return false
		}
	}
	return true
}


func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}