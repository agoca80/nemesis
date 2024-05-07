package main

import (
	"slices"
	"strconv"
)

func (player *Player) ResolveMove(corridor *Corridor) {
	if player.IsInCombat() {
		Show(player, "tries to leave", player.Area)
	}
	for _, intruder := range player.Area.Intruders {
		intruder.ResolveAttack(player)
	}

	if !player.Alive() {
		return
	}

	destination := corridor.End(player.Area)
	noiseRoll := player.MovesTo(destination)

	var event string
	if !player.Area.IsExplored() {
		event = game.ResolveExploration(player, corridor)
	}

	if ship.Destroyed() {
		return
	}

	if c := player.Area.Carcasses(); c > 0 {
		Show("There are", c, "carcasses in the area")
	}

	if noiseRoll && event != ev_danger && event != ev_silence {
		player.ResolveNoise()
	}
}

func (a ActionBasic) Resolve(data map[string]interface{}) {
	player := data["player"].(*Player)
	switch string(a) {
	case basic_move:
		corridor := data["corridor"].(*Corridor)
		player.ResolveMove(corridor)
	default:
		Pending(a, "not implemented")
	}
}

func (p *Player) Pay(card Card) {
	index := slices.Index(p.Hand, card)
	p.Hand = slices.Delete(p.Hand, index, index+1)
	p.Discard(card)
}

func resolveAction(action *Action) {
	player := action.Player
	for _, card := range action.Cost {
		player.Pay(card)
	}

	switch action.Name {
	case basic_move:
		corridor := action.Data["corridor"].(*Corridor)
		player.ResolveMove(corridor)
	case basic_fire:
		intruder := action.Data["intruder"].(*Intruder)
		player.ResolveFire(intruder)
	default:
		panic("WTF")
	}
}

func (player *Player) ResolveFire(intruder *Intruder) {
	var damage int
	var roll = player.RollDamage()
	switch roll {
	case damage_double:
		damage = 2
	case damage_single:
		damage = 1
	case intruder_adult:
		if symbols(intruder_adult, intruder_crawler, intruder_larva, intruder_egg).Contains(intruder.Kind) {
			damage = 1
		}
	case intruder_crawler:
		if symbols(intruder_crawler, intruder_larva, intruder_egg).Contains(intruder.Kind) {
			damage = 1
		}
	}

	Show(player, "fires against", intruder, ": rolls", roll, ", deals", damage, "damage")
	intruder.Suffers(damage)
}

func (p *Player) ResolveNoise() {
	encounter := false
	result := p.RollNoise()
	if result == ev_silence && p.IsDrenched {
		Show(p, "was silent but is drenched in mucus!")
		result = ev_danger
	}

	switch result {
	case ev_silence:
		Show(p, "is silent...")
	case ev_danger:
		Show(p, "is in danger!")
		p.Danger()
	default:
		n, _ := strconv.Atoi(result)
		corridor := p.Corridor(n)
		Show(p, "makes noise in corridor", corridor.Numbers)
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
		game.ResolveEncounter(p)
	}
}
