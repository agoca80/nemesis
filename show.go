package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

func (i *Intruder) Show() string {
	return fmt.Sprintf("%v%d,%d", i.Kind, i.Id, i.Wounds)
}

func (c *Contamination) Show() string {
	return c.Infected.String()
}

func (corridors Corridors) String() (result string) {
	doors := ""
	noise := ""
	for _, corridor := range corridors {
		doors += corridor.Door
		noise += corridor.Noise.String()
	}
	result = fmt.Sprintf(
		"%v\t%v",
		noise,
		doors,
	)
	return
}

func (a *Area) Show() string {
	actors := []string{}
	for _, p := range a.Players {
		actors = append(actors, p.Character)
	}
	for _, i := range a.Intruders {
		actors = append(actors, i.Show())
	}

	var description string
	if a.IsExplored() {
		description = a.name
	} else {
		description = "Unexplored"
	}

	return fmt.Sprintf(" %v%v %s,%d %-21s\t> %v\t%v\t| %v",
		a.IsInFire,
		a.IsDamaged,
		a,
		a.Items,
		description,
		a.Corridors,
		a.Neighbors(),
		strings.Join(actors, " "),
	)
}

func (p *Player) Show() string {
	return fmt.Sprintf("%v\t(%v)\t%v\t%v+%v\tHand %v",
		p.Character,
		p.IsInfected,
		p.State,
		p.Bruises,
		p.Wounds,
		p.Hand,
	)
}

func (p Players) Show() {
	if !show_players {
		return
	}

	output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	for _, player := range p {
		fmt.Fprintf(output, "%v\n", player.Show())
	}
	output.Flush()
	Show()
}

func (b *Board) Show() {
	if !show_board {
		return
	}

	Show(strings.Repeat("-", 58))
	output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	for _, a := range b.Area {
		if !a.IsReachable() {
			continue
		}
		fmt.Fprintf(output, "%v\n", a.Show())
	}
	fmt.Fprint(output)
	output.Flush()
	Show(strings.Repeat("-", 58))
	Show()
}

func (game *Game) Show() {
	game.Players.Show()
	game.Board.Show()
}

func Show(args ...interface{}) {
	fmt.Println(args...)
}

func Wait() {
	if wait {
		Show("Press Enter to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
