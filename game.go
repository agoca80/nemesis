package main

import (
	"math/rand"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Game struct {
	Intruders
	Players

	*Board
	*IntruderBag

	// Decks
	Attacks       *Deck
	Contamination *Deck
	Events        *Deck
	Weaknesses    *Deck
	Wounds        *Deck

	Coordinate          *Coordinates
	EscapePods          []Cards
	Destination         string
	Eggs                int
	EngineStatus        [3]bool
	hyperdriveCountdown int
	Weakness            []*Weakness
}

func (game *Game) ResolveEncounter(p *Player) {
	token := game.FetchToken()
	kind := token.name
	if kind == token_blank {
		Show("Unexpectedly, nothing happened")
		return
	}

	intruder := game.NewIntruder(token, p.Area)
	if p.IsSurprised(token) {
		Show(p, "is surprised by", intruder)
		game.ResolveIntruderAttack(intruder, p)
	}
}

func (g *Game) ResolveMove(action *ActionPlayer) {
	player := action.Player
	corridor := action.Corridor
	area := action.Player.Area
	if player.IsInCombat() {
		Show(player, "tries to leave", player.Area)
		player.Area.Intruders.Attack(player)
	}

	if !player.Alive() {
		return
	}

	// Show(player, "moves from", player.Area, "to", corridor.Area, "through", corridor.Numbers)
	area.RemPlayer(player)
	area = corridor.Area
	wasEmpty := area.IsEmpty()
	player.Area, area.Players = area, append(area.Players, player)
	Show(player, "enters", area)

	if !wasEmpty {
		return
	}

	var event string
	if !area.IsExplored() {
		event = g.ResolveExploration(player, corridor)
	}

	if g.Destroyed() {
		Pending("The ship has been destroyed!!!")
	}

	if event != ev_danger && event != ev_silence {
		g.ResolveNoise(player)
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

func (g *Game) ResolveNoise(p *Player) {
	encounter := false
	result := p.RollNoise()
	if result == ev_silence && p.IsDrenched {
		Show(p, "was silent but is drenched in mucus")
		result = ev_danger
	}

	switch result {
	case ev_silence:
		Show("There is no noise")
	case ev_danger:
		Show(p.Area, "is in danger!")
		p.Danger()
	default:
		n, _ := strconv.Atoi(result)
		corridor := p.Corridor(n)
		Show("Noise in corridor", p.Area, corridor.Numbers)
		if corridor.Noise {
			for _, c := range p.Area.Corridors {
				c.Noise = false
			}
			encounter = true
		} else {
			corridor.Noise = true
		}
	}

	if encounter {
		g.ResolveEncounter(p)
	}
}

func Roll(dice []string) string {
	face := rand.Intn(len(dice))
	return dice[face]
}

type ActionPlayer struct {
	Action
	*Player
	*Corridor
}

func (g *Game) AskAction(p *Player) (nextAction bool) {
	action := p.AskAction()
	if action == nil {

		return
	}

	actions := map[string]func(*ActionPlayer){
		basic_move: g.ResolveMove,
	}

	if fn, ok := actions[action.Name()]; !ok {
		Show("PENDING ACTION", action)
		os.Exit(1)
	} else {
		Show()
		fn(action)
	}

	return true
}

func (g *Game) RemIntruder(i *Intruder) {
	index, a := slices.Index(g.Intruders, i), i.Area
	g.Intruders = slices.Delete(g.Intruders, index, index+1)
	a.RemIntruder(i)
}

func (g *Game) FireDamage(i *Intruder, damage int) {
	for _, w := range g.Weakness {
		if w.Revealed && w.name == weakness_fire {
			damage++
		}
	}

	i.Suffers(damage)
}

func (g *Game) gameOver() bool {
	return g.Destroyed() || len(g.Players.Alive()) == 0 || g.hyperdriveCountdown == 0
}

func (g *Game) GoingOn() bool {
	return !g.gameOver()
}

func (g *Game) Run() {
	var step = map[Step]func() Step{
		step_draw:      g.draw,
		step_token:     g.token,
		step_turn:      g.turn,
		step_counters:  g.counters,
		step_attack:    g.attacks,
		step_fire:      g.fireDamage,
		step_event:     g.event,
		step_evolution: g.evolution,
	}

	s := Step(step_draw)
	for g.GoingOn() {
		Show(strings.Repeat("-", 80))
		Show("STEP", s)
		s = step[s]()
	}

	Pending("Game over!!!")
}
