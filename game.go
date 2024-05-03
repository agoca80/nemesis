package main

import (
	"math/rand"
	"slices"
	"strings"
)

type Game struct {
	Intruders
	Players

	*Board
	*IntruderBag

	// Decks
	GoalsCoop *Deck
	GoalsCorp *Deck
	GoalsPriv *Deck

	CoordinateCard      *Coordinates
	EscapePods          []Cards
	Destination         string
	Eggs                int
	EngineStatus        [3]bool
	hyperdriveCountdown int
	Weakness            []*Weakness
}

func newGame(players int) (game *Game) {
	game = &Game{
		Players: Players{},

		GoalsCoop: newDeck(goalsCoop),
		GoalsCorp: newDeck(goalsCorp[:players+4]),
		GoalsPriv: newDeck(goalsPriv[:players+4]),
	}

	for range players {
		game.Players = append(game.Players, NewPlayer())
	}

	return
}

func (game *Game) ResolveEncounter(p *Player) {
	token := game.FetchToken()
	kind := token.name
	if kind == token_blank {
		Show("Unexpectedly, nothing happened")
		return
	}

	intruder := NewIntruder(token, p.Area)
	if p.IsSurprised(token) {
		Show(p, "is surprised by", intruder)
		ResolveIntruderAttack(intruder, p)
	}
}

func (game *Game) ResolveExploration(player *Player, corridor *Corridor) (event string) {
	area := player.Area
	token := area.ExplorationToken
	area.ExplorationToken = nil

	event = token.Event
	if event == ev_silence && player.IsDrenched {
		Show(player, "tried to be silent but they were too noisy...")
		event = ev_danger
	}

	switch event {
	case ev_damaged:
		area.IsDamaged = true
	case ev_danger:
		Show(area, "is dangerous!")
		area.Danger()
	case ev_door:
		if corridor.Door != door_broken {
			Show(corridor, "door closes")
			corridor.Door = door_closed
		} else {
			Show(corridor, "tried to close but is broken")
		}
	case ev_fire:
		area.IsInFire = true
	case ev_mucus:
		Show(area, "is full of mucus")
		player.IsDrenched = true
	case ev_silence:
		Show(area, "is silent")
	}

	if area.name != room_nest && area.name != room_slime {
		area.Items = token.Items
	}

	return event
}

func Roll(dice []string) string {
	face := rand.Intn(len(dice))
	return dice[face]
}

func RemIntruder(i *Intruder) {
	index, a := slices.Index(game.Intruders, i), i.Area
	game.Intruders = slices.Delete(game.Intruders, index, index+1)
	a.RemIntruder(i)
}

func (i *Intruder) FireDamage(damage int) {
	Show(i, "in", i.Area, "is damaged by fire!")

	for _, w := range game.Weakness {
		if w.Revealed && w.name == weakness_fire {
			damage++
		}
	}

	i.Suffers(damage)
}

func (g *Game) Over() bool {
	return g.Destroyed() || len(g.Players.Alive()) == 0 || g.hyperdriveCountdown == 0
}

func (g *Game) Run() {
	var step = map[Step]func() Step{
		step_draw:      g.draw,
		step_token:     g.token,
		step_turn:      g.stepTurn,
		step_counters:  g.counters,
		step_attack:    g.stepAttacks,
		step_fire:      g.fireDamage,
		step_event:     g.event,
		step_evolution: g.evolution,
	}

	s := Step(step_draw)
	for !g.Over() {
		Show(strings.Repeat("-", 80))
		Show(strings.ToUpper(string(s)))
		s = step[s]()
	}

	game.Players.Show()
	Show("Game over!!!")
}
