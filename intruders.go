package main

import (
	"fmt"
	"slices"
)

const (
// weakness_intruder = iota
// weakness_corpse
// weakness_egg
)

var intruderIds = map[string]int{
	intruder_egg:     0,
	intruder_larva:   0,
	intruder_crawler: 0,
	intruder_adult:   0,
	intruder_breeder: 0,
	intruder_queen:   0,
}

type Intruder struct {
	*Area

	Id     int
	Kind   string
	Wounds int
	Dead   bool
}

type Intruders []*Intruder

func (i *Intruder) Leaves() {
	i.Area.RemIntruder(i)
	i.Area = nil
}

func (i *Intruder) Enters(area *Area) {
	i.Area = area
	area.Intruders = append(area.Intruders, i)
}

func (i *Intruder) Moves(number int) {
	corridor := i.Corridor(number)
	destination := corridor.End(i.Area)
	if corridor.Door == door_closed {
		Show(i, "in area", i.Area, "destroys", corridor.Numbers, "door!")
		corridor.Door = door_broken
		return
	} else {
		Show(i, "in area", i.Area, "moves to", destination, "through corridor", corridor.Numbers)
		i.Leaves()
		i.Enters(destination)
	}
}

func (i *Intruder) String() string {
	return fmt.Sprintf("%s%d", i.Kind, i.Id)
}

func (i *Intruder) IsAlive() bool {
	return !i.Dead
}

func (i *Intruder) InCombat() bool {
	return len(i.Area.Players.Alive()) > 0
}

func (intruder *Intruder) Attack() {
	if !intruder.InCombat() {
		return
	}

	// Choose the player with the smallest hand size
	players := intruder.Area.Players.Alive()
	smallest := players[0]
	for _, p := range players.Alive() {
		if p.HandSize() < smallest.HandSize() {
			smallest = p
		}
	}

	intruder.ResolveAttack(smallest)
}

func newIntruder(kind string, area *Area) (i *Intruder) {
	Show(kind, "appears in", area, "!")
	i = &Intruder{
		Id:   intruderIds[kind],
		Area: area,
		Kind: kind,
	}
	intruderIds[kind]++
	intruders = append(game.Intruders(), i)
	area.Intruders = append(area.Intruders, i)
	return
}

func spawnIntruder(token *IntruderToken, area *Area) (i *Intruder) {
	i = newIntruder(token.Kind, area)
	game.Retire(token)
	return
}

func (i *Intruder) Suffers(damage int) {
	if damage == 0 {
		return
	}

	if i.Kind == intruder_egg || i.Kind == intruder_larva {
		i.Dies()
		return
	}

	check := 0
	cards := []*Attack{}
	switch {
	case i.Kind == intruder_adult || i.Kind == intruder_crawler:
		cards = append(cards, attacks.Next().(*Attack))
		check = cards[0].Wounds
	case i.Kind == intruder_breeder || i.Kind == intruder_queen:
		cards = append(cards, attacks.Next().(*Attack))
		cards = append(cards, attacks.Next().(*Attack))
		check = cards[0].Wounds + cards[1].Wounds
	}

	i.Wounds += damage
	for _, c := range cards {
		Show(i, "draws", c)
		if c.Retreats() {
			Show(i, "in", i.Area, "retreats!")
			direction := events.Next().(*Event)
			i.Moves(direction.Corridor)
			return
		}
	}

	if i.Wounds >= check {
		i.Dies()
	}
}

func (i *Intruder) Dies() {
	Show(i, "squacks and dies!")
	i.Area.Objects = append(i.Area.Objects, &Object{
		Area: i.Area,
		Name: "carcass",
	})
	i.Leaves()
	i.Dead = true
	index := slices.Index(intruders, i)
	intruders = slices.Delete(intruders, index, index+1)
}

func (i *Intruder) ResolveAttack(p *Player) {
	effect := map[string]func(*Intruder, *Player){
		attack_claws:          attackClaws,
		attack_frenzy:         attackFrenzy,
		attack_tail:           attackTail,
		attack_transformation: attackTransformation,
		attack_bite:           attackByte,
		attack_mucosity:       attackMucosity,
		attack_recall:         attackRecall,
		attack_scratch:        attackScratch,
	}

	if !p.Alive() {
		Show(i, "was going to attack", p, "but it was already dead!")
		return
	}

	if i.Kind == intruder_larva {
		Show(i, "infects", p, "!")
		p.IsInfected = true
		p.SuffersContamination()
		RemIntruder(i)
		return
	}

	attack := attacks.Next().(*Attack)
	if !attack.Contains(i.Kind) {
		Show(i, "attacks", p, "but fails!")
		return
	}

	if fn, ok := effect[attack.name]; !ok {
		Show("PENDING", attack, "function not implemented.")
	} else {
		fn(i, p)
	}
}

func (i *Intruder) Alive() bool {
	return !i.Dead
}
