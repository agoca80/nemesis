package main

import (
	"cmp"
	"fmt"
	"math/rand"
	"slices"
	"strings"
)

func (game *Game) Prepare(coop bool) {
	Show(strings.Repeat("-", 58))
	Show(strings.ToUpper("Preparation"))
	Show()

	// Prepare 1
	ship = newShip()

	// Prepare 2
	for _, a := range ship.Area {
		if a.Class == room_2 {
			a.Room = rooms2.Draw().(*Room)
		}
	}

	// Prepare 3
	for _, a := range ship.Area {
		if a.Class == room_1 {
			a.Room = rooms1.Draw().(*Room)
		}
	}

	// Prepare 4
	for _, a := range ship.Area {
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
	game.IntruderBag = NewIntruderBag(len(players))

	// Initialize hyperdrive countdown
	game.hyperdriveCountdown = 15

	// Crew preparation step 18 A,B,C
	for _, p := range players {
		p.Area, ship.Area[A11].Players = ship.Area[A11], append(ship.Area[A11].Players, p)
	}

	// Crew preparation step 14
	helpDeck := newDeck(helpCards[:len(players)])
	for _, player := range players {
		player.HelpCard = helpDeck.Draw().(*HelpCard)
		Show(player, "takes", player.HelpCard)
	}
	Show()

	// Crew preparation step 16
	for _, p := range players {
		if coop {
			p.Goals = append(p.Goals, game.GoalsCoop.Draw())
		} else {
			p.Goals = append(p.Goals, game.GoalsCorp.Draw(), game.GoalsPriv.Draw())
		}
	}

	// Crew preparation step 17
	sorted := players[:]
	slices.SortFunc(sorted, func(a, b *Player) int {
		return cmp.Compare(a.Number, b.Number)
	})
	for _, player := range sorted {
		options := []Card{
			characters.Draw(),
			characters.Draw(),
		}
		selected, rejected := player.Choose(options)
		Show(fmt.Sprintf("%s picks %-9s, rejects %s", player, selected, rejected))
		player.Character = selected.Name()
		characters.Return(rejected)
		characters.Shuffle()
	}
	Show()

	// Crew preparation step 18
	for _, p := range players {
		p.Deck = actions[p.Character]
	}

	// Step 19
	for _, p := range players {
		if p.Number == 1 {
			p.Jonesy = true
		}
	}

	// Preparation 20
	// hybernarium.Objects = append(hybernarium.Objects, game.BlueCorpseToken)
}
