package main

const (
	debug        = true
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
}
