package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func (i *Intruder) Show() string {
	return fmt.Sprintf("%v,%d", i, i.Wounds)
}

func (c *Contamination) Show() string {
	return Issue(c.Infected).String()
}

func (a *Area) Show() string {
	return fmt.Sprintf("%v%v %s > %s\t%d,%s\t| %v\t| %v",
		Issue(a.IsBurning()),
		Issue(a.IsDamaged()),
		a,
		a.ShowCorridors(),
		a.Items,
		a.Describe(),
		ShowList(a.Players.Alive()),
		ShowList(a.Intruders),
	)
}

func (p Players) Show() {
	if !show_players {
		return
	}

	output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	for _, player := range p {
		fmt.Fprintf(output, "%v\n", player.Describe())
	}
	output.Flush()
	Show()
}

func (s *Ship) Show() {
	if !show_ship {
		return
	}

	output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	for _, a := range s.Area {
		if !a.IsReachable() {
			continue
		}

		if len(a.Intruders) == 0 && len(a.Players) == 0 {
			continue
		}

		fmt.Fprintf(output, "%v\n", a.Show())

	}
	fmt.Fprint(output)
	output.Flush()
	Show()
}

func Show(args ...interface{}) {
	fmt.Println(args...)
}
