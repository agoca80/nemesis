package main

import (
	"os"
)

const (
	debug        = true
	show_board   = true
	show_players = true
	wait         = false
)

var (
	game        *Game
	intruderBag *IntruderBag
)

func main() {
	cooperative := false

	game = newGame()
	game.Prepare(cooperative)
	game.Run()
	finish()
}

func Pending(args ...interface{}) {
	Show("PENDING", args)
	os.Exit(1)
}
