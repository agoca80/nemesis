package main

import (
	"os"
)

const (
	show_ship    = true
	show_players = true
	wait         = false
)

var (
	game *Game

	ship    *Ship
	players Players
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
