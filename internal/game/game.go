package game

import (
	"bufio"
	"fmt"
	"io"
	"noughts-and-crosses/internal/board"
	"noughts-and-crosses/internal/cli"
)

const PlayerTurnError = GameLoopError("PlayerTurn")

type GameLoopError string

func (e GameLoopError) Error() string {
	return string(e)
}

type Game struct {
	Board        *board.Board
	Player       string
	InputStream  io.Reader
	OutputStream io.Writer
	GameOver     bool
}

func NewGame(input io.Reader, output io.Writer) *Game {

	g := &Game{
		Board:        &board.Board{},
		Player:       "X",
		InputStream:  input,
		OutputStream: output,
		GameOver:     false,
	}
	return g
}

func (g *Game) SwitchPlayer() {
	if g.Player == "X" {
		g.Player = "O"
	} else {
		g.Player = "X"
	}
}

func (g *Game) Run() {

	fmt.Fprintln(g.OutputStream, "Starting a new Game")
	g.Board.InitBoard()

	g.GameLoop()

}

func (g *Game) GameLoop() {

	reader := bufio.NewReader(g.InputStream)
	for !g.GameOver {
		cli.RenderBoard(g.OutputStream, g.Board.Grid)
		fmt.Fprintln(g.OutputStream)

		err := g.PlayerTurn(reader)

		if err == nil {
			isWin := g.Board.CheckForWin()

			if isWin {
				fmt.Fprintf(g.OutputStream, "Player %s wins!\n", g.Player)
				g.GameOver = true
			} else if g.Board.IsBoardFull() {
				fmt.Fprintln(g.OutputStream, "It is a draw!")
				g.GameOver = true
			} else {
				g.SwitchPlayer()
			}
		} else {
			fmt.Fprintf(g.OutputStream, "\n\nInvalid move! Please put in the formay of x,y. Take your turn again.\n")
		}

	}
	cli.RenderBoard(g.OutputStream, g.Board.Grid)
}

func (g *Game) PlayerTurn(reader *bufio.Reader) error {

	fmt.Fprintf(g.OutputStream, "Player %s, please enter your move!: ", g.Player)
	num1, num2, err := cli.GetPlayerMove(reader)

	if err != nil {
		return PlayerTurnError
	}

	move := board.GridCell{
		Row: num1,
		Col: num2,
	}

	isValid := g.Board.ValidateMove(move)

	if !isValid {
		return PlayerTurnError
	}

	g.Board.PlaceMove(move, g.Player)
	return nil
}
