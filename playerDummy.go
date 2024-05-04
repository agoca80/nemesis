package main

import (
	"math/rand"
)

type Dummy struct {
	*player
}

func newDummy() Controller {
	return &Dummy{}
}

func (p *Dummy) Choose(cards Cards) (selected, rejected Card) {
	shuffle := rand.Perm(2)
	selected, rejected = cards[shuffle[0]], cards[shuffle[1]]
	return
}

func (p *Dummy) chooseCorridor() *Corridor {
	options := Corridors{}
	for _, c := range p.Area.Corridors {
		if c.IsReachable() {
			options = append(options, c)
		}
	}
	return options[rand.Intn(len(options))]
}

func (p *Dummy) NextAction() (action map[string]interface{}) {
	if p.HandSize() < 1 {
		return
	}

	return map[string]interface{}{
		"action":   ActionBasic(basic_move),
		"corridor": p.chooseCorridor(),
		"cost": Cards{
			p.Hand[0],
		},
	}
}
