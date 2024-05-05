package main

import "slices"

type Action interface {
	Name() string
	Cost() int
	Resolve(map[string]interface{})
}

func (player *Player) MovesTo(dstArea *Area) (moiseRoll bool) {
	moiseRoll = dstArea.IsEmpty()
	srcArea := player.Area
	srcArea.RemPlayer(player)
	player.Area, dstArea.Players = dstArea, append(dstArea.Players, player)
	Show(player, "moves to", dstArea)
	return
}

func (p *Player) Pay(card Card) {
	index := slices.Index(p.Hand, card)
	if index == -1 {
		panic("WTF")
	}
	p.Hand = slices.Delete(p.Hand, index, index+1)
	p.Discard(card)
}
