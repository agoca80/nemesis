package main

import (
	"os"
)

const (
	show_board   = true
	show_players = true
	wait         = true
)

var (
	game *Game
)

func main() {
	cooperative := false
	players := 1

	game = newGame(players)
	game.Prepare(cooperative)
	game.Run()
	Show("Game over!!!")
}

func Pending(args ...interface{}) {
	Show("PENDING", args)
	os.Exit(1)
}
