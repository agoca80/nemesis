package main

import (
	"math/rand"
	"slices"
	"strings"
)

var (
	intruders Intruders
	players   Players
	ship      *Ship
)

type Game struct {
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

func newGame(numPlayers int) (game *Game) {
	players = Players{}

	game = &Game{
		GoalsCoop: newDeck(goalsCoop),
		GoalsCorp: newDeck(goalsCorp[:numPlayers+4]),
		GoalsPriv: newDeck(goalsPriv[:numPlayers+4]),
		HelpDeck:  newDeck(helpCards[:numPlayers]),
	}

	for range numPlayers {
		players = append(players, NewPlayer())
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
		intruder.ResolveAttack(p)
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
		area.Damage()
	case ev_danger:
		area.Danger()
	case ev_door:
		if corridor.Door == door_open {
			corridor.Door = door_closed
		}
	case ev_fire:
		area.Burning()
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
	index, a := slices.Index(intruders, i), i.Area
	intruders = slices.Delete(intruders, index, index+1)
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

func gameOver() bool {
	return ship.Destroyed() || len(players.Alive()) == 0 || game.hyperdriveCountdown == 0
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
	for !gameOver() {
		Show(strings.Repeat("-", 58))
		Show(strings.ToUpper(string(s)))
		Show()
		s = step[s]()
	}
}

func (game *Game) Intruders() Intruders {
	return intruders
}
