package main

import (
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
