package main

import (
	"os"
)

const (
	show_board   = true
	show_players = true
	wait         = false
)

var (
	game *Game
)

func main() {
	cooperative := false
	players := 5

	game = newGame(players)
	game.Prepare(cooperative)
	game.Run()
	finish()
}

func Pending(args ...interface{}) {
	Show("PENDING", args)
	os.Exit(1)
}
