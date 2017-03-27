package ttgBoard

import "testing"

func TestNewBoard(t *testing.T) {
	gboard := NewBoard(3, 3, 2, 3)
	want := "000\n000\n000\n"
	got := gboard.String()
	if got != want {
		t.Errorf("NewBoard(3, 3, 2, 3) ==\n%q\nwant\n%q", got, want)
	}
}

func TestChangeWhoseTurn(t *testing.T) {
	gboard := NewBoard(3, 3, 2, 3)
	gboard.ChangeWhoseTurn()
	want := 2
	got := gboard.WhoseTurn
	if got != want {
		t.Errorf("NewBoard(3, 3, 2, 3).ChangeWhoseTurn(): WhoseTurn == %d, want: %d", got, want)
	}
}

func TestTakeTurn(t *testing.T) {
	gboard := NewBoard(3, 3, 2, 3)
	gboard.TakeTurn(0, 1)
	want := "010\n000\n000\n"
	got := gboard.String()
	if got != want {
		t.Errorf("NewBoard(3, 3, 2, 3).TakeTurn(0, 1) ==\n%q\nwant\n%q", got, want)
	}
}

func TestTakeMultipleTurns(t *testing.T) {
	gboard := NewBoard(3, 3, 2, 3)
	gboard.TakeTurn(0, 1)
	gboard.TakeTurn(1,1)
	gboard.TakeTurn(0, 0)
	want := "110\n020\n000\n"
	got := gboard.String()
	if got != want {
		t.Errorf("NewBoard(3, 3, 2, 3).TakeTurn(0, 1) ==\n%q\nwant\n%q", got, want)
	}
}

func TestTakeInvalidTurn(t *testing.T) {
	gboard := NewBoard(3, 3, 2, 3)
	gboard.TakeTurn(0, 1)
	firstState := gboard.String()
	// Next turn is invalid and should not change the game board:
	gboard.TakeTurn(0, 1)
	secondState := gboard.String()
	if firstState != secondState {
		t.Errorf("State changed after taking invalid turn:\nfrom %q\n to %q", firstState, secondState)
	}
}