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

// Returns true only if the given space is not already occupied.
func (board gameBoard) isValidTurn(xpos, ypos int) bool {
	return board.BoardState[ypos][xpos] == 0
}

// Increments which player's turn it is.
func (board gameBoard) changeWhoseTurn() {
	player := board.WhoseTurn
	player = player % board.NumPlayers
	player += 1
	board.WhoseTurn = player
}

// Attempts to make specified move. If invalid move, returns false.
func (board gameBoard) TakeTurn(xpos, ypos int) bool {
	if ! board.isValidTurn(xpos, ypos) {
		return false
	}
	board.BoardState[ypos][xpos] = board.WhoseTurn
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