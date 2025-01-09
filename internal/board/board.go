package board

type Board struct {
	Grid [][]string
	Size int
}

type GridCell struct {
	Row int
	Col int
}

const InvalidMoveError = BoardError("Invalid Move")

type BoardError string

func (e BoardError) Error() string {
	return string(e)
}

func (b *Board) InitBoard() {
	b.Grid = [][]string{{"", "", ""}, {"", "", ""}, {"", "", ""}}
	b.Size = 3
}

func (b *Board) ValidateMove(move GridCell) bool {
	return b.isInGrid(move)
}

func (b *Board) PlaceMove(move GridCell, symbol string) error {

	if !b.isInGrid(move) {
		return InvalidMoveError
	}
	b.Grid[move.Row][move.Col] = symbol
	return nil
}

func (b *Board) CheckForWin() bool {
	return b.checkRowsForWin() ||
		b.checkColumnsForWin() ||
		b.checkRightDiagonalForWin() ||
		b.checkLeftDiagonalForWin()
}

func (b *Board) checkRowsForWin() bool {

	for _, row := range b.Grid {
		if checkArray(row) {
			return true
		}
	}
	return false
}

func (b *Board) checkColumnsForWin() bool {
	for col_num := 0; col_num < b.Size; col_num++ {
		col := []string{}
		for row_num := 0; row_num < b.Size; row_num++ {
			col = append(col, b.Grid[row_num][col_num])
		}
		if checkArray(col) {
			return true
		}
	}
	return false
}

func (b *Board) checkRightDiagonalForWin() bool {

	slice := []string{}
	for num := 0; num < b.Size; num++ {
		slice = append(slice, b.Grid[num][num])
	}
	return checkArray(slice)
}

func (b *Board) checkLeftDiagonalForWin() bool {
	slice := []string{}
	max_col := b.Size - 1
	for num := 0; num < b.Size; num++ {
		slice = append(slice, b.Grid[num][max_col-num])
	}
	return checkArray(slice)
}

func (b *Board) IsBoardFull() bool {
	for row := 0; row < b.Size; row++ {
		for col := 0; col < b.Size; col++ {
			if b.Grid[row][col] == "" {
				return false
			}
		}
	}
	return true
}

func (b *Board) isInGrid(move GridCell) bool {
	if move.Row < 0 || move.Col < 0 || move.Row >= b.Size || move.Col >= b.Size {
		return false
	}
	return b.Grid[move.Row][move.Col] == ""
}

func checkArray(slice []string) bool {
	first := slice[0]
	if first == "" {
		return false
	}
	for _, cell := range slice {
		if first != cell {
			return false
		}
	}
	return true
}
