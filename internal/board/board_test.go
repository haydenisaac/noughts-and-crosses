package board

import (
	"reflect"
	"testing"
)

func setupBoard() Board {
	return Board{
		Grid: [][]string{
			{"", "", ""},
			{"", "", ""},
			{"", "", ""},
		},
		Size: 3,
	}
}

func TestBoardInit(t *testing.T) {
	board := Board{}

	if board.Grid != nil {
		t.Errorf("Board already initialised - %v", board.Grid)
	}

	board.InitBoard()

	empty_grid := [][]string{{"", "", ""}, {"", "", ""}, {"", "", ""}}

	if !reflect.DeepEqual(board.Grid, empty_grid) {
		t.Errorf("empty board not initialised - got %v, want %v", board.Grid, empty_grid)
	}
}

func TestValidateMove(t *testing.T) {

	t.Run("ValidMove", func(t *testing.T) {
		board := setupBoard()
		move := GridCell{Row: 1, Col: 1}
		if !board.ValidateMove(move) {
			t.Error("Expected true for valid move on an empty grid")
		}
	})

	t.Run("InvalidMove", func(t *testing.T) {
		board := setupBoard()
		board.Grid = [][]string{
			{"X", "", ""},
			{"", "", ""},
			{"", "", ""},
		}
		move := GridCell{Row: 0, Col: 0}
		if board.ValidateMove(move) {
			t.Errorf("Expected an invalid error for move %v on board %v", move, board.Grid)
		}
	})

	t.Run("OutOfBounds", func(t *testing.T) {
		board := setupBoard()
		moves := []GridCell{
			{Row: -1, Col: -1},
			{Row: 3, Col: 0},
			{Row: 0, Col: 3},
		}
		for _, move := range moves {
			if board.ValidateMove(move) {
				t.Error("Expected to be false as the move is off of the board")
			}
		}
	})
}

func TestWinConditions(t *testing.T) {
	t.Run("HorizontalWin", func(t *testing.T) {
		board := setupBoard()
		board.Grid[1] = []string{"X", "X", "X"}

		if !board.CheckForWin() {
			t.Errorf("Expected to find win on board %v", board.Grid)
		}
	})

	t.Run("NoWin", func(t *testing.T) {
		board := setupBoard()
		if board.CheckForWin() {
			t.Errorf("Expected no win on board %v", board.Grid)
		}
	})

	t.Run("VerticalWin", func(t *testing.T) {
		board := setupBoard()
		board.Grid[0][1] = "X"
		board.Grid[1][1] = "X"
		board.Grid[2][1] = "X"

		if !board.CheckForWin() {
			t.Errorf("Expected to find a vertical win on board %v", board.Grid)
		}
	})

	t.Run("RightDiagonalWin", func(t *testing.T) {
		board := setupBoard()
		board.Grid[0][0] = "X"
		board.Grid[1][1] = "X"
		board.Grid[2][2] = "X"

		if !board.CheckForWin() {
			t.Errorf("Expected to find a diagonal win on board %v", board.Grid)
		}
	})

	t.Run("LeftDiagonalWin", func(t *testing.T) {
		board := setupBoard()
		board.Grid[0][2] = "X"
		board.Grid[1][1] = "X"
		board.Grid[2][0] = "X"

		if !board.CheckForWin() {
			t.Errorf("Expected to find a diagonal win on board %v", board.Grid)
		}
	})
}

func TestDrawCondition(t *testing.T) {

	t.Run("FullBoard", func(t *testing.T) {
		board := Board{
			Grid: [][]string{
				{"X", "O", "X"},
				{"X", "O", "X"},
				{"O", "X", "O"},
			},
			Size: 3,
		}

		if !board.IsBoardFull() {
			t.Errorf("Expected to get a full board with %v", board.Grid)
		}
	})

	t.Run("EmptyBoard", func(t *testing.T) {
		board := setupBoard()
		if board.IsBoardFull() {
			t.Errorf("Expected to not get a full board with %v", board.Grid)
		}
	})
}

func TestPlaceMove(t *testing.T) {

	t.Run("ValidMove", func(t *testing.T) {
		board := setupBoard()

		cell := GridCell{Row: 0, Col: 0}

		board.PlaceMove(cell, "X")

		updated_board := [][]string{
			{"X", "", ""},
			{"", "", ""},
			{"", "", ""},
		}

		if !reflect.DeepEqual(board.Grid, updated_board) {
			t.Errorf("Board not as expected. wanted %v got %v", updated_board, board.Grid)
		}
	})

	t.Run("InvalidMove", func(t *testing.T) {
		board := setupBoard()
		board.Grid = [][]string{
			{"X", "", ""},
			{"", "", ""},
			{"", "", ""},
		}
		move := GridCell{Row: 0, Col: 0}
		err := board.PlaceMove(move, "X")
		if err != InvalidMoveError {
			t.Errorf("Expected an invalid error for move %v on board %v", move, board.Grid)
		}
	})

	t.Run("OutOfBounds", func(t *testing.T) {
		board := setupBoard()
		moves := []GridCell{
			{Row: -1, Col: -1},
			{Row: 3, Col: 0},
			{Row: 0, Col: 3},
		}
		for _, move := range moves {
			err := board.PlaceMove(move, "X")
			if err != InvalidMoveError {
				t.Error("Expected to be raise an error as the move is off of the board")
			}
		}
	})
}
