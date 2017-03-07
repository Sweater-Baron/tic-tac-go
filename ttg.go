package ttgBoard

var PLAYER_1 = 1
var PLAYER_2 = 2

type Board struct {
	BoardState [3][3]int
	WhoseTurn int
}

// Returns true only if the given space is not already occupied.
func (board Board) isValidTurn(xpos, ypos int) bool {
	return board.BoardState[ypos][xpos] == 0
}

// Changes which player's turn it is.
func (board Board) swapWhoseTurn() {
	switch turn := board.WhoseTurn {
	case PLAYER_1:
		board.WhoseTurn = PLAYER_2
	default:
		board.WhoseTurn = PLAYER_1
	}
}

// Attempts to make specified move. If invalid move, returns false.
func (board Board) TakeTurn(xpos, ypos int) bool {
	if ! board.isValidTurn(xpos, ypos) {
		return false
	}
	board.BoardState[ypos][xpos] = board.WhoseTurn
	board.swapWhoseTurn()
	return true
}