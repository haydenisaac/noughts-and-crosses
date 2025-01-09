package cli

import (
	"bytes"
	"testing"
)

func TestRenderBoard(t *testing.T) {

	buffer := bytes.Buffer{}
	board := [][]string{
		{"X", "O", "X"},
		{"X", "O", "X"},
		{"O", "X", "O"},
	}

	RenderBoard(&buffer, board)

	got := buffer.String()
	want := "\nX|O|X\n-+-+-\nX|O|X\n-+-+-\nO|X|O\n"

	if got != want {
		t.Errorf("Board display doesn't match, got %s and wanted %s", got, want)
	}

}
