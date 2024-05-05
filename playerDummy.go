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

func (p *Player) NewAction() (action *action) {
	if p.HandSize() < 1 {
		return
	}

	switch {
	case p.IsInCombat():
		Show("Debug", p)
		action = &action{
			Action: ActionBasic(basic_fire),
			data: actionData{
				"intruder": p.Area.Intruders[rand.Intn(len(p.Area.Intruders))],
				"cost": Cards{
					p.Hand[0],
				},
			},
		}
	case !p.IsInCombat():
		action = &action{
			Action: ActionBasic(basic_move),
			data: actionData{
				"corridor": p.chooseCorridor(),
				"cost": Cards{
					p.Hand[0],
				},
			},
		}
	}

	action.player = p
	return action
}
