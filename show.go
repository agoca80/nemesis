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

func (c *Gate) Show() string {
	return fmt.Sprintf(
		"%v%v%v%v",
		Noise(c.Noise),
		c.Numbers,
		c.Door,
		c.Area,
	)
}

func (c *Gates) Show() string {
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

	var description string
	if a.IsExplored() {
		description = a.name
	} else {
		description = "Unexplored"
	}

	return fmt.Sprintf("- %v%v %s,%d %-21s\t %s\t| %v\t| %v\n",
		Damage(a.IsDamaged),
		Fire(a.IsInFire),
		a,
		a.Items,
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

func (g *Game) Show() {
	g.Players.Show()
	g.Board.Show()
}

func Wait() {
	if wait {
		Show("Press Enter to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
