package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

func Show(args ...interface{}) {
	fmt.Println(args...)
}

func (a *Area) Show() string {
	actors := []string{}
	for _, p := range a.Players {
		actors = append(actors, p.Character)
	}
	for _, i := range a.Intruders {
		actors = append(actors, i.Kind)
	}

	var description string
	if a.IsExplored() {
		description = a.name
	} else {
		description = "Unexplored"
	}

	neighbors := []string{}
	for _, n := range a.Neighbors() {
		neighbors = append(neighbors, n.String())
	}

	return fmt.Sprintf("- %v%v %s,%d -> %s\t%-21s\t %s\t| %v\t| %v\n",
		Damage(a.IsDamaged),
		Fire(a.IsInFire),
		a,
		a.Items,
		strings.Join(neighbors, ","),
		description,
		a.ExplorationToken,
		a.Gates,
		strings.Join(actors, " "),
	)
}

func (p *Player) Show() string {
	return fmt.Sprintf("%-9v %v LightWounds %v SeriousWounds %v Hand %v",
		p.Character,
		p.State,
		p.LightWounds,
		p.SeriousWounds,
		p.Hand,
	)
}

func (p Players) Show() {
	if !show_players {
		return
	}

	Show()
	for _, player := range p {
		fmt.Println(player.Show())
	}
}

func (b *Board) Show() {
	if !show_board {
		return
	}

	Show()
	output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	for _, a := range b.Area {
		if !a.IsReachable() {
			continue
		}
		fmt.Fprintf(output, "%v", a.Show())
	}
	fmt.Fprint(output)
	output.Flush()
	Show()
}

func (game *Game) Show() {
	game.Players.Show()
	game.Board.Show()
}

func Wait() {
	if wait {
		Show("Press Enter to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
