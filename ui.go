package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

func (i *Intruder) Show() string {
	return fmt.Sprintf("%v,%d", i, i.Wounds)
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

func ShowList[T fmt.Stringer](list []T) string {
	strs := []string{}
	for _, item := range list {
		strs = append(strs, item.String())
	}
	return strings.Join(strs, " ")
}

func (a *Area) Show() string {
	var description string
	if a.IsExplored() {
		description = a.name
	} else {
		description = "Unexplored"
	}

	return fmt.Sprintf(" %v%v %s,%d %-21s\t> %v\t%v\t| %v\t| %v\t| %v",
		a.IsInFire,
		a.IsDamaged,
		a,
		a.Items,
		description,
		a.Corridors,
		a.Neighbors(),
		ShowList(a.Players),
		ShowList(a.Objects),
		ShowList(a.Intruders),
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

func Wait() {
	if wait {
		Show("Press Enter to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}

func Prompt(message string) {
	fmt.Print("PROMPT ", message, " > ")
}

func Choose(cards Cards) (selected, rejected Card) {
	var stdin = bufio.NewReader(os.Stdin)
	var choice string
	for _, c := range cards {
		Show(c.Id(), c.Name())
	}
	for {
		Prompt("Choose a card")
		input, _ := fmt.Fscanln(stdin, &choice)
		switch {
		case input == 0:
			fallthrough
		case choice == cards[0].Id():
			return cards[0], cards[1]
		case choice == cards[1].Id():
			return cards[1], cards[0]
		}
	}
}
