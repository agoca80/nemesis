package main

import "slices"

func (player *player) MovesTo(dstArea *Area) (moiseRoll bool) {
	moiseRoll = dstArea.IsEmpty()
	srcArea := player.Area
	srcArea.RemPlayer(player)
	player.Area, dstArea.Players = dstArea, append(dstArea.Players, player)
	Show(player, "moves to", dstArea)
	return
}

func (player *player) ResolveMove(corridor *Corridor) {
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

	if game.Destroyed() {
		return
	}

	if noiseRoll && event != ev_danger && event != ev_silence {
		player.ResolveNoise()
	}
}

func (a ActionBasic) Resolve(data map[string]interface{}) {
	player := data["player"].(*player)
	switch string(a) {
	case basic_move:
		corridor := data["corridor"].(*Corridor)
		player.ResolveMove(corridor)
	default:
		Pending(a, "not implemented")
	}
}

func (p *player) Pay(card Card) {
	index := slices.Index(p.Hand, card)
	if index == -1 {
		panic("WTF")
	}
	p.Hand = slices.Delete(p.Hand, index, index+1)
	p.Discard(card)
}
