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
	for _, c := range p.Area.Corridors {
		if c.IsReachable() {
			options = append(options, c)
		}
	}
	return options[rand.Intn(len(options))]
}

func (player *Player) NewAction() *Action {
	if player.HandSize() < 1 {
		return nil
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
		corridor := player.chooseCorridor()
		name, data = basic_move, map[string]interface{}{
			"corridor": corridor,
		}
	}

	return &Action{
		Cost:   Cards{player.Hand[rand.Intn(len(player.Hand))]},
		Player: player,
		Name:   name,
		Data:   data,
	}
}
