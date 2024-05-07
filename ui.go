package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

func (i *Intruder) Show() string {
	return fmt.Sprintf("%v,%d", i, i.Wounds)
}

func (c *Contamination) Show() string {
	return Issue(c.Infected).String()
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

func ShowList[T fmt.Stringer](list []T) string {
	strs := []string{}
	for _, item := range list {
		strs = append(strs, item.String())
	}
	return strings.Join(strs, " ")
}

func (a *Area) Show() string {
	return fmt.Sprintf(" %v%v %s,%d %-21s\t> %v\t%v\t| %v\t| %v",
		Issue(a.IsBurning()),
		Issue(a.IsDamaged()),
		a,
		a.Items,
		a.Describe(),
		a.Corridors,
		a.Neighbors(),
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
		if a.IsReachable() {
			fmt.Fprintf(output, "%v\n", a.Show())
		}
	}
	fmt.Fprint(output)
	output.Flush()
	Show()
}

func Show(args ...interface{}) {
	fmt.Println(args...)
}
