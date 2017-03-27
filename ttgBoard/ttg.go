package ttgBoard

import (
	"strconv"
	"strings"
)

type gameBoard struct {
	BoardState [][]int
	WhoseTurn, NumPlayers, Cols, Rows, WinNumber int
}

// Returns a new instance of a gameBoard.
func NewBoard(numRows, numCols, numPlayers, winNumber int) *gameBoard {
	board := new(gameBoard)
	board.WhoseTurn = 1
	board.Rows = numRows
	board.Cols = numCols
	board.NumPlayers = numPlayers
	board.WinNumber = winNumber
	board.BoardState = make([][]int, numRows, numRows)
	for i := range board.BoardState {
		board.BoardState[i] = make([]int, numCols, numCols)
	}
	return board
}

// Returns true only if the given row and column are not out of range.
func (board gameBoard) isInRange(row, col int) bool {
	return row >= 0 && row < board.Rows && col >= 0 && col < board.Cols
}

// Returns true only if the given space exists and is not already occupied.
func (board gameBoard) isValidTurn(row, col int) bool {
	if ! board.isInRange(row, col) {
		return false
	}
	return board.BoardState[row][col] == 0
}

// Increments which player's turn it is.
func (board gameBoard) changeWhoseTurn() {
	player := board.WhoseTurn
	player = player % board.NumPlayers
	player += 1
	board.WhoseTurn = player
}

// Attempts to make specified move. If invalid move, returns false.
func (board gameBoard) TakeTurn(row, col int) bool {
	if ! board.isValidTurn(row, col) {
		return false
	}
	board.BoardState[row][col] = board.WhoseTurn
	board.changeWhoseTurn()
	return true
}

// Returns the game board represented as a string
func (board gameBoard) String() string {
	var results []string
	for _, row := range board.BoardState {
		for _, cellVal := range row {
			results = append(results, strconv.Itoa(cellVal))
		}
		results = append(results, "\n")
	}
	return strings.Join(results, "")
}