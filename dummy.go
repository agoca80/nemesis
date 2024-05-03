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

func (p *Player) chooseCorridor() *Gate {
	options := Gates{}
	for _, c := range p.Gates {
		if c.IsReachable() {
			options = append(options, c)
		}
	}
	return options[rand.Intn(len(options))]
}
