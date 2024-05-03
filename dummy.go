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

func (p *Player) NewAction() (actionData map[string]interface{}) {
	if p.HandSize() < 1 {
		return
	}

	return map[string]interface{}{
		"player":   p,
		"action":   ActionBasic(basic_move),
		"corridor": p.chooseCorridor(),
		"cost": Cards{
			p.Hand[0],
		},
	}
}
