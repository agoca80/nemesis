package main

import (
	"fmt"
	"slices"
)

type Area struct {
	Id        int
	Class     string
	IsDamaged Issue
	IsInFire  Issue

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

func (a *Area) Carcasses() (total int) {
	for _, o := range a.Objects {
		if o.Name == object_carcass {
			total++
		}
	}
	return
}

func (a *Area) Danger() {
	for _, c := range a.Corridors {
		c.Noise = true
	}
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
	neighbors = Areas{}
	for _, corridor := range a.Corridors {
		end := corridor.End(a)
		if end.IsReachable() {
			neighbors = append(neighbors, corridor.End(a))
		}
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
	for _, c := range a.Corridors {
		if slices.Contains(c.Numbers, n) {
			return c
		}
	}
	return nil
}
