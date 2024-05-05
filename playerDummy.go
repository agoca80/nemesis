package main

import (
	"math/rand"
)

type InputDummy struct {
}

func newDummy() *InputDummy {
	return &InputDummy{}
}

func (input *InputDummy) Choose(cards Cards) (selected, rejected Card) {
	shuffle := rand.Perm(2)
	selected, rejected = cards[shuffle[0]], cards[shuffle[1]]
	return
}

func (p *Player) chooseCorridor() *Corridor {
	options := Corridors{}
	for _, c := range p.Area.Corridors {
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
