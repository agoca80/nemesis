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
	HelpDeck  *Deck
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
		HelpDeck:  newDeck(helpCards[:players]),
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

	intruder := spawnIntruder(token, p.Area)
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
	Show(player, "reveals", token, "in the", area.name)
	if event == ev_silence && player.IsDrenched {
		Show(player, "tried to be silent but they were too noisy...")
		event = ev_danger
	}

	switch event {
	case ev_damaged:
		area.IsDamaged = true
	case ev_danger:
		area.Danger()
	case ev_door:
		if corridor.Door == door_open {
			corridor.Door = door_closed
		}
	case ev_fire:
		area.IsInFire = true
	case ev_mucus:
		player.IsDrenched = true
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

func (game *Game) Run() {
	var step = map[Step]func() Step{
		step_draw:      game.draw,
		step_token:     game.token,
		step_turn:      game.stepTurn,
		step_counters:  game.counters,
		step_attack:    game.stepAttacks,
		step_fire:      game.fireDamage,
		step_event:     game.event,
		step_evolution: game.evolution,
	}

	s := Step(step_draw)
	for !game.Over() {
		Show(strings.Repeat("-", 58))
		Show(strings.ToUpper(string(s)))
		Show()
		s = step[s]()
	}
}
