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

func TestTakeTurn1(t *testing.T) {
	gboard := NewBoard(3, 3, 2, 3)
	gboard.TakeTurn(0, 1)
	want := "010\n000\n000\n"
	got := gboard.String()
	if got != want {
		t.Errorf("NewBoard(3, 3, 2, 3).TakeTurn(0, 1) ==\n%q\nwant\n%q", got, want)
	}
}