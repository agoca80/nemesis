package main

const (
// weakness_intruder = iota
// weakness_corpse
// weakness_egg
)

type Intruder struct {
	*Area
	*Game

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
	if corridor.Door == door_closed {
		Show(i, "in area", i.Area, "destroys", corridor.Numbers, "door!")
		corridor.Door = door_broken
		return
	} else {
		Show(i, "in area", i.Area, "moves to", corridor.Area, "through corridor", corridor.Numbers)
		i.Leaves()
		i.Enters(corridor.Area)
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

	i.ResolveIntruderAttack(i, player)
}

func (i Intruders) Attack(p *Player) {
	for _, intruder := range i {
		intruder.ResolveIntruderAttack(intruder, p)
	}
}

func (g *Game) SpawnIntruder(kind string, area *Area) (i *Intruder) {
	Show(kind, "appears in", area, "!")
	i = &Intruder{
		Area: area,
		Game: g,
		Kind: kind,
	}
	area.Intruders = append(area.Intruders, i)
	g.Intruders = append(g.Intruders, i)
	return
}

func (g *Game) NewIntruder(token *IntruderToken, area *Area) (i *Intruder) {
	i = g.SpawnIntruder(token.Kind, area)
	g.Retire(token)
	return
}

func (i *Intruder) Dies() {
	Pending(i, "Intruder has died!")
}

func (i *Intruder) Suffers(damage int) (dies bool) {
	i.Wounds += damage
	Show("PENDING intruder suffers")
	return i.Wounds > 1
}
