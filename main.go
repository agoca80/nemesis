package main

import (
	"os"
)

const (
	show_board   = true
	show_players = true
	wait         = true
)

func main() {
	cooperative := false
	players := 1

	g := newGame(players)
	g.Prepare(cooperative)
	g.Run()
	Show("Game over!!!")
}

func Pending(args ...interface{}) {
	Show("PENDING", args)
	os.Exit(1)
}
