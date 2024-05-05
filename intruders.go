package main

const (
// weakness_intruder = iota
// weakness_corpse
// weakness_egg
)

type Intruder struct {
	*Area

	Kind   string
	Wounds int
	Dead   bool
}

type Intruders []*Intruder

func (i *Intruder) CanAttack() bool {
	return len(i.Area.Players) > 0
}

func (i *Intruder) IsInCombat() bool {
	return len(i.Area.Players) > 0
}

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
	return i.Kind
}

func (i *Intruder) IsAlive() bool {
	return !i.Dead
}

func (i *Intruder) Attack() {
	// This should be in player attack order
	player := i.Area.Players[0]
	for _, p := range i.Area.Players.Alive() {
		if p.HandSize() < player.HandSize() {
			player = p
		}
	}

	ResolveIntruderAttack(i, player)
}

func (i Intruders) Attack(p *Player) {
	for _, intruder := range i {
		ResolveIntruderAttack(intruder, p)
	}
}

func SpawnIntruder(kind string, area *Area) (i *Intruder) {
	Show(kind, "appears in", area, "!")
	i = &Intruder{
		Area: area,
		Kind: kind,
	}
	game.Intruders = append(game.Intruders, i)
	area.Intruders = append(area.Intruders, i)
	return
}

func NewIntruder(token *IntruderToken, area *Area) (i *Intruder) {
	i = SpawnIntruder(token.Kind, area)
	game.Retire(token)
	return
}

func (i *Intruder) Dies() {
	Pending(i, "Intruder has died!")
}

func (i *Intruder) Suffers(damage int) (dies bool) {
	if damage == 0 {
		return
	}

	if i.Kind == intruder_egg || i.Kind == intruder_larva {
		i.Dies()
		return true
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
		if c.Retreats() {
			Show(i, "in", i.Area, "retreats!")
			direction := events.Next().(*Event)
			i.Moves(direction.Corridor)
			return
		}
	}

	if dies = i.Wounds >= check; dies {
		i.Dies()
	}
	return
}
