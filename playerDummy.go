package main

import (
	"math/rand"
)

func dummyChoose(options Cards) (selected, rejected Card) {
	shuffle := rand.Perm(2)
	return options[shuffle[0]], options[shuffle[1]]
}

func dummyChooseCorridor(p *Player) *Corridor {
	options := Corridors{}
	for _, c := range p.Area.Corridors {
		if c.IsReachable() {
			options = append(options, c)
		}
	}
	return options[rand.Intn(len(options))]
}

func dummyAction(player *Player) (action *Action) {
	if player.HandSize() < 1 {
		return
	}

	var name string
	var data map[string]interface{}
	switch {
	case player.IsInCombat():
		intruder := player.Area.Intruders[rand.Intn(len(player.Area.Intruders))]
		name, data = basic_fire, map[string]interface{}{
			"intruder": intruder,
		}
	case !player.IsInCombat():
		corridor := dummyChooseCorridor(player)
		name, data = basic_move, map[string]interface{}{
			"corridor": corridor,
		}
	}

	action = &Action{
		Cost:   Cards{player.Hand[rand.Intn(len(player.Hand))]},
		Player: player,
		Name:   name,
		Data:   data,
	}

	return action
}
