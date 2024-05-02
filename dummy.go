package main

import (
	"math/rand"
)

func (p *Player) chooseCharacter(characters *Deck) {
	shuffle := rand.Perm(2)
	options := []Card{
		characters.Draw(),
		characters.Draw(),
	}
	p.Character = options[shuffle[0]].Name()
	characters.Return(options[shuffle[1]])
	characters.Shuffle()
}

func (p *Player) chooseCorridor() *Corridor {
	options := Corridors{}
	for _, c := range p.Corridors {
		if c.IsReachable() {
			options = append(options, c)
		}
	}
	return options[rand.Intn(len(options))]
}

func (p *Player) AskAction() *ActionPlayer {
	if p.HandSize() < 1 {
		p.Passes()
		return nil
	} else {
		p.Discard(p.Hand[0])
		p.Hand = p.Hand[1:]
	}

	return &ActionPlayer{
		Player:   p,
		Action:   ActionBasic(basic_move),
		Corridor: p.chooseCorridor(),
	}
}
