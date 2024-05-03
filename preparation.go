package main

import (
	"math/rand"
)

func (g *Game) Prepare(coop bool) {
	// Prepare 1
	g.Board = NewBoard()

	// Prepare 2
	for _, a := range g.Area {
		if a.Class == room_2 {
			a.Room = g.Rooms2.Draw().(*Room)
		}
	}

	// Prepare 3
	for _, a := range g.Area {
		if a.Class == room_1 {
			a.Room = g.Rooms1.Draw().(*Room)
		}
	}

	// Prepare 4
	for _, a := range g.Area {
		if a.Class == room_1 || a.Class == room_2 {
			a.ExplorationToken = g.ExplorationTokens.Draw().(*ExplorationToken)
		}
	}

	// Initialize coordinates card
	g.CoordinateCard = g.Coordinates.Draw().(*Coordinates)

	// Initialize destination
	g.Destination = "B"

	// Initialize escape pods
	// order := rand.Perm(2)
	// alternate := []int{0, 1, order[0], order[1]}
	// numOfEscapePods := []int{2, 2, 3, 3, 4}[len(g.Players)-1]
	// for n := range numOfEscapePods {
	// 	search := func(c Card) bool {
	// 		ep := c.(*EscapePod)
	// 		return ep.number == n
	// 	}
	// 	g.EscapePods[alternate[n]] = append(g.EscapePods[alternate[n]], g.Take(escapePods, search))
	// }

	// initialize engines states
	g.EngineStatus = [3]bool{
		rand.Intn(2) == 0,
		rand.Intn(2) == 0,
		rand.Intn(2) == 0,
	}

	// Initialize intruder board
	g.Eggs = 5
	g.Weakness = []*Weakness{
		g.Weaknesses.Draw().(*Weakness),
		g.Weaknesses.Draw().(*Weakness),
		g.Weaknesses.Draw().(*Weakness),
	}

	// Initialize intruder bag
	g.IntruderBag = NewIntruderBag(len(g.Players))

	// Initialize hyperdrive countdown
	g.hyperdriveCountdown = 15

	// Crew preparation step 18 A,B,C
	for _, p := range g.Players {
		p.Area, g.Area[A11].Players = g.Area[A11], append(g.Area[A11].Players, p)
	}

	// Crew preparation step 14
	helpCards := Cards{
		NewHelpCard(1),
		NewHelpCard(2),
		NewHelpCard(3),
		NewHelpCard(4),
		NewHelpCard(5),
	}
	helpDeck := NewDeck(helpCards[:len(g.Players)])
	for _, p := range g.Players {
		p.HelpCard = helpDeck.Draw().(*HelpCard)
	}

	// Crew preparation step 16
	for _, p := range g.Players {
		if coop {
			p.Goals = append(p.Goals, g.GoalsCoop.Draw())
		} else {
			p.Goals = append(p.Goals, g.GoalsCorp.Draw(), g.GoalsPriv.Draw())
		}
	}

	// Crew preparation step 17
	for _, p := range g.Players {
		p.chooseCharacter(g.Characters)
	}

	// Crew preparation step 18
	for _, p := range g.Players {
		p.Deck = NewDeck(actions[p.Character])
	}

	// Step 19
	for _, p := range g.Players {
		if p.Number == 1 {
			p.Jonesy = true
		}
	}

	// Preparation 20
	// hybernarium.Objects = append(hybernarium.Objects, g.BlueCorpseToken)
}
