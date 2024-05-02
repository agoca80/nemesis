package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

const (
	show_board = false
)

func main() {
	cooperative := false
	players := 3

	g := NewGame(players)
	g.Prepare(cooperative)
	g.Run()
	Show("Game over!!!")
}

func Pending(args ...interface{}) {
	Show("PENDING", args)
	os.Exit(1)
}

func WaitForKey() {
	// Show("Press Enter to continue...")
	// bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func (c *Corridor) Show() string {
	return fmt.Sprintf(
		"%v%v%v%v",
		Noise(c.Noise),
		c.Numbers,
		c.Door,
		c.Area,
	)
}

func (c *Corridors) Show() string {
	corridors := []string{}
	for _, c := range *c {
		corridors = append(corridors, c.Show())
	}
	return strings.Join(corridors, " ")
}

func (a *Area) Show() string {
	actors := []string{}
	for _, p := range a.Players {
		actors = append(actors, p.Character)
	}
	for _, i := range a.Intruders {
		actors = append(actors, i.Kind)
	}

	return fmt.Sprintf("- %v%v %s,%d %s \t %v \t %v\n",
		Damage(a.IsDamaged),
		Fire(a.IsInFire),
		a,
		a.Items,
		a.ExplorationToken,
		a.Corridors.Show(),
		strings.Join(actors, " "),
	)
}

func (p Players) Show() {
	for _, player := range p {
		fmt.Println("-", player.Character, player.State, player.LightWounds, player.SeriousWounds)
	}
}

func (b *Board) Show() {
	output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	for _, a := range b.Area {
		if !a.IsReachable() {
			continue
		}
		fmt.Fprintf(output, "%v", a.Show())
	}
	fmt.Fprintln(output)
	output.Flush()
}

func (g *Game) Show() {
	if show_board {
		g.Board.Show()
	}
	WaitForKey()
}

func Show(args ...interface{}) {
	fmt.Println(args...)
}
