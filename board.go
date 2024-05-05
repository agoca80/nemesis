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

type Corridor struct {
	Id    int
	AreaX *Area
	AreaY *Area
	Door  string
	Noise Issue
	Numbers
}

func (c *Corridor) End(a *Area) *Area {
	if a == c.AreaX {
		return c.AreaY
	}
	return c.AreaX
}

func (c *Corridor) IsReachable() bool {
	return c.Door == door_open && c.AreaX.IsReachable() && c.AreaY.IsReachable()
}

type Corridors []*Corridor

type Gates []*struct {
	X int
	Y int
	N []int
}

type Board struct {
	Area []*Area
}

func NewBoard() (b *Board) {
	b = &Board{
		Area: areas,
	}

	for _, area := range b.Area {
		if area.Class != room_1 && area.Class != room_2 {
			area.Room = &Room{
				card: newCard(area.Class),
			}
		}
	}

	corridorId := 0
	for _, gate := range gates {
		corridorId++
		areaX, areaY := b.Area[gate.X], b.Area[gate.Y]
		c := &Corridor{
			Id:      corridorId,
			AreaX:   areaX,
			AreaY:   areaY,
			Door:    door_open,
			Numbers: gate.N,
		}
		areaX.Corridors = append(areaX.Corridors, c)
		areaY.Corridors = append(areaY.Corridors, c)
	}

	return
}

func (b *Board) Damages() (damaged int) {
	for _, area := range b.Area {
		if area.IsDamaged {
			damaged++
		}
	}
	return
}

func (b *Board) Fires() (result int) {
	for _, area := range b.Area {
		if area.IsInFire {
			result++
		}
	}
	return
}

func (b *Board) Destroyed() bool {
	return b.Damages() > 8 || b.Fires() > 8
}
