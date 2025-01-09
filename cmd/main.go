package main

import (
	"noughts-and-crosses/internal/game"
	"os"
)

func main() {

	g := game.NewGame(os.Stdin, os.Stdout)

	g.Run()

}
