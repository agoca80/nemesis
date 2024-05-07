package main

import (
	"cmp"
	"fmt"
	"slices"
)

type Area struct {
	isDamaged bool
	isInFire  bool

	Id    int
	Class string

	Corridors
	Intruders
	Objects
	Players

	*Room
}

type Areas []*Area

func newArea(id int, class string) *Area {
	return &Area{
		Id:        id,
		Class:     class,
		Intruders: Intruders{},
		Objects:   Objects{},
		Players:   Players{},
	}
}

func (a *Area) Describe() string {
	if a.IsExplored() {
		return a.name
	}
	return "Unexplored"
}

func (a *Area) Carcasses() (total int) {
	Each(a.Objects, func(o *Object) {
		if o.Name == object_carcass {
			total++
		}
	})
	return
}

func (a *Area) Danger() {
	Each(a.Corridors, (*Corridor).Danger)
}

func (a *Area) IsExplored() bool {
	return a.ExplorationToken == nil
}

func (a *Area) IsEmpty() bool {
	return len(a.Players) == 0 && len(a.Intruders) == 0
}

func (a *Area) IsReachable() bool {
	return A00 < a.Id && a.Id < S01
}

func (a *Area) Neighbors() (neighbors Areas) {
	reachable := Filter(a.Corridors, (*Corridor).IsReachable)
	for _, corridor := range reachable {
		neighbors = append(neighbors, corridor.End(a))
	}

	return neighbors
}

func (a *Area) RemIntruder(i *Intruder) {
	index := slices.Index(a.Intruders, i)
	a.Intruders = slices.Delete(a.Intruders, index, index+1)
}

func (a *Area) RemPlayer(p *Player) {
	index := slices.Index(a.Players, p)
	a.Players = slices.Delete(a.Players, index, index+1)
}

func (a *Area) String() string {
	return fmt.Sprintf("%02d", a.Id)
}

func (a *Area) Corridor(n int) *Corridor {
	index := slices.IndexFunc(a.Corridors, func(c *Corridor) bool {
		return slices.Contains(c.Numbers, n)
	})
	return a.Corridors[index]
}

func (a *Area) IsDamaged() bool {
	return a.isDamaged
}

func (a *Area) Damage() {
	a.isDamaged = true
}

func (a *Area) Burning() {
	a.isInFire = true
}

func (a *Area) IsBurning() bool {
	return a.isInFire
}

func (a *Area) ShowCorridors() (str string) {
	ends := Symbols{}
	numbers := Symbols{}
	doors := ""
	noise := ""
	for number := 1; number < 5; number++ {
		corridor := a.Corridor(number)
		switch {
		case corridor.IsReachable():
			doors += corridor.Door
			noise += Issue(corridor.Noise).String()
			numbers = append(numbers, corridor.Numbers.String())
			ends = append(ends, corridor.End(a).String())
		default:
			doors += "T"
			noise += Issue(corridor.Noise).String()
			numbers = append(numbers, corridor.Numbers.String())
			ends = append(ends, "ST")
		}
	}

	return fmt.Sprintf("%s %s %s %s", doors, ends, numbers, noise)
}

// func (a *Area) ShowCorridors() (str string) {
// 	corridors := make(Corridors, len(a.Corridors))
// 	copy(corridors, a.Corridors)
// 	slices.SortFunc(corridors, func(c1, c2 *Corridor) int {
// 		return cmp.Compare(c1.End(a).Id, c2.End(a).Id)
// 	})

// 	ends := Symbols{}
// 	numbers := Symbols{}
// 	doors := ""
// 	noise := ""
// 	for _, corridor := range corridors {
// 		doors += corridor.Door
// 		noise += Issue(corridor.Noise).String()
// 		numbers = append(numbers, corridor.Numbers.String())
// 		ends = append(ends, corridor.End(a).String())
// 	}

// 	return fmt.Sprintf("%s\t%s\t%s\t%s", ends, doors, numbers, noise)
// }

type Direction struct {
	*Area
	*Corridor
}

type Directions []*Direction

func (a *Area) Directions() (d Directions) {
	for _, corridor := range a.Corridors {
		if corridor.IsReachable() && corridor.Door != door_closed {
			d = append(d, &Direction{corridor.End(a), corridor})
		}
	}
	slices.SortFunc(d, func(d1, d2 *Direction) int {
		return cmp.Compare(d1.Corridor.Numbers.String(), d2.Corridor.Numbers.String())
	})
	return
}

func (d *Direction) String() string {
	return fmt.Sprintf("%v>%v", d.Area, d.Corridor.Numbers)
}
