package game

import (
	"noughts-and-crosses/internal/board"
	"os"
	"reflect"
	"testing"
)

func TestNewGame(t *testing.T) {

	g := NewGame(os.Stdin, os.Stdout)

	if reflect.DeepEqual(g.Board, [][]string{{"", "", ""}, {"", "", ""}, {"", "", ""}}) {
		t.Error("Not setup the board")

	}

	if g.Player != "X" {
		t.Error("player not initialised correctly")
	}

	if g.InputStream != os.Stdin {
		t.Error("Input stream not setup to os")
	}
	if g.OutputStream != os.Stdout {
		t.Error("Output steam not setup to os")
	}
}

func TestSwitchPlayer(t *testing.T) {
	g := Game{&board.Board{}, "X", os.Stdin, os.Stdout, false}

	g.SwitchPlayer()
	if g.Player != "O" {
		t.Error("player didn't switch")
	}
	g.SwitchPlayer()
	if g.Player != "X" {
		t.Error("player didn't switch")
	}

}
