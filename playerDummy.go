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

func (player *Player) NewAction() (a *action) {
	if player.HandSize() < 1 {
		return
	}

	a = &action{}
	a.Player = player

	// switch {
	// case player.IsInCombat():
	// 	Show("Debug", player)
	// 	a = &action{
	// 		Action: newActionBasic(basic_fire),
	// 		data: data{
	// 			"intruder": player.Area.Intruders[rand.Intn(len(player.Area.Intruders))],
	// 			"cost": Cards{
	// 				player.Hand[0],
	// 			},
	// 		},
	// 	}
	// case !player.IsInCombat():
	// 	a = &action{
	// 		Action: newActionBasic(basic_move),
	// 		data: data{
	// 			"corridor": player.chooseCorridor(),
	// 			"cost": Cards{
	// 				player.Hand[0],
	// 			},
	// 		},
	// 	}
	// }

	return
}
