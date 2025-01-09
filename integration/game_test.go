package integration

import (
	"bytes"
	"noughts-and-crosses/internal/game"
	"testing"
)

func TestGameLoop(t *testing.T) {

	t.Run("X win", func(t *testing.T) {
		inputs := "0,0\n1,1\n0,1\n2,2\n0,2\n"
		mockInput := bytes.NewBufferString(inputs)
		mockOutput := &bytes.Buffer{}

		g := game.NewGame(mockInput, mockOutput)
		g.Board.InitBoard()
		g.GameLoop()

		output := mockOutput.String()
		if !bytes.Contains([]byte(output), []byte("Player X wins!")) {
			t.Errorf("Expected 'Player X wins!' message in output, but got:\n%s", output)
		}

		if !g.GameOver {
			t.Errorf("Should have been a game over")
		}
	})

	t.Run("Y win", func(t *testing.T) {
		inputs := "0,0\n1,1\n2,1\n1,2\n0,2\n1,0\n"
		mockInput := bytes.NewBufferString(inputs)
		mockOutput := &bytes.Buffer{}

		g := game.NewGame(mockInput, mockOutput)
		g.Board.InitBoard()
		g.GameLoop()

		output := mockOutput.String()
		if !bytes.Contains([]byte(output), []byte("Player O wins!")) {
			t.Errorf("Expected 'Player O wins!' message in output, but got:\n%s", output)
		}

		if !g.GameOver {
			t.Errorf("Should have been a game over")
		}

	})

	t.Run("Error", func(t *testing.T) {
		inputs := "3,3\n1,1\n1,2\n0,0\n1,0\n2,2\n"
		mockInput := bytes.NewBufferString(inputs)
		mockOutput := &bytes.Buffer{}

		g := game.NewGame(mockInput, mockOutput)
		g.Board.InitBoard()
		g.GameLoop()

		output := mockOutput.String()
		if !bytes.Contains([]byte(output), []byte("Invalid move!")) {
			t.Errorf("Expected 'Invalid Move!' message in output, but got:\n%s", output)
		}
	})

}
