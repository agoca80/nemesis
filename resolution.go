package main

import "slices"

func (player *Player) MovesTo(dstArea *Area) (moiseRoll bool) {
	moiseRoll = dstArea.IsEmpty()
	srcArea := player.Area
	srcArea.RemPlayer(player)
	player.Area, dstArea.Players = dstArea, append(dstArea.Players, player)
	Show(player, "moves to", dstArea)
	return
}

func (player *Player) ResolveMove(corridor *Corridor) {
	if player.IsInCombat() {
		Show(player, "tries to leave", player.Area)
		player.Area.Intruders.Attack(player)
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
	if index == -1 {
		panic("WTF")
	}
	p.Hand = slices.Delete(p.Hand, index, index+1)
	p.Discard(card)
}

func resolveAction(action *Action) {
	player := action.Player
	for _, card := range action.Cost {
		player.Discard(card)
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
