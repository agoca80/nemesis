package main

import (
	"math/rand"
)

func (p *Dummy) NextAction(availableActions Actions) (action Action) {
	if p.HandSize() < 1 {
		return
	}

	action = &struct {
		Action
		Data map[string]interface{}
	}{
		Action: ActionBasic(basic_move),
		Data: map[string]interface{}{
			"corridor": p.chooseCorridor(),
		},
	}
}

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
