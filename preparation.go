package main

import (
	"math/rand"
)

func (game *Game) Prepare(coop bool) {
	// Prepare 1
	game.Board = NewBoard()

	// Prepare 2
	for _, a := range game.Area {
		if a.Class == room_2 {
			a.Room = rooms2.Draw().(*Room)
		}
	}

	// Prepare 3
	for _, a := range game.Area {
		if a.Class == room_1 {
			a.Room = rooms1.Draw().(*Room)
		}
	}

	// Prepare 4
	for _, a := range game.Area {
		if a.Class == room_1 || a.Class == room_2 {
			a.ExplorationToken = explorationTokens.Draw().(*ExplorationToken)
		}
	}

	// Initialize coordinates card
	game.CoordinateCard = coordinates.Draw().(*Coordinates)

	// Initialize destination
	game.Destination = "B"

	// Initialize escape pods
	// order := rand.Perm(2)
	// alternate := []int{0, 1, order[0], order[1]}
	// numOfEscapePods := []int{2, 2, 3, 3, 4}[len(game.Players)-1]
	// for n := range numOfEscapePods {
	// 	search := func(c Card) bool {
	// 		ep := c.(*EscapePod)
	// 		return ep.number == n
	// 	}
	// 	game.EscapePods[alternate[n]] = append(game.EscapePods[alternate[n]], game.Take(escapePods, search))
	// }

	// initialize engines states
	game.EngineStatus = [3]bool{
		rand.Intn(2) == 0,
		rand.Intn(2) == 0,
		rand.Intn(2) == 0,
	}

	// Initialize intruder board
	game.Eggs = 5
	game.Weakness = []*Weakness{
		weaknesses.Draw().(*Weakness),
		weaknesses.Draw().(*Weakness),
		weaknesses.Draw().(*Weakness),
	}

	// Initialize intruder bag
	game.IntruderBag = NewIntruderBag(len(game.Players))

	// Initialize hyperdrive countdown
	game.hyperdriveCountdown = 15

	// Crew preparation step 18 A,B,C
	for _, p := range game.Players {
		p.Area, game.Area[A11].Players = game.Area[A11], append(game.Area[A11].Players, p)
	}

	// Crew preparation step 14
	helpDeck := newDeck(helpCards[:len(game.Players)])
	for _, p := range game.Players {
		p.HelpCard = helpDeck.Draw().(*HelpCard)
	}

	// Crew preparation step 16
	for _, p := range game.Players {
		if coop {
			p.Goals = append(p.Goals, game.GoalsCoop.Draw())
		} else {
			p.Goals = append(p.Goals, game.GoalsCorp.Draw(), game.GoalsPriv.Draw())
		}
	}

	// Crew preparation step 17
	for _, p := range game.Players {
		p.chooseCharacter(characters)
	}

	// Crew preparation step 18
	for _, p := range game.Players {
		p.Deck = actions[p.Character]
	}

	// Step 19
	for _, p := range game.Players {
		if p.Number == 1 {
			p.Jonesy = true
		}
	}

	// Preparation 20
	// hybernarium.Objects = append(hybernarium.Objects, game.BlueCorpseToken)
}
