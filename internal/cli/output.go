package cli

import (
	"fmt"
	"io"
)

func RenderBoard(writer io.Writer, board [][]string) {

	string := "\n"
	for row_num, row := range board {
		for col_num, cell := range row {
			if cell == "" {
				string += " "
			} else {
				string += cell
			}

			if col_num+1 < len(row) {
				string += "|"
			}

		}
		if row_num+1 < len(board) {
			string += "\n-+-+-"
		}
		string += "\n"
	}

	fmt.Fprint(writer, string)
}
