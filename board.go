package main

import (
	"fmt"
	"slices"
	"strings"
)

type Area struct {
	Id        int
	Class     string
	IsDamaged bool
	IsInFire  bool

	Gates
	Intruders
	Players

	*Room
}

type Areas []*Area

func newArea(id int, class string) *Area {
	return &Area{
		Id:        id,
		Class:     class,
		Players:   Players{},
		Intruders: Intruders{},
	}
}

func (a *Area) Danger() {
	for _, c := range a.Gates {
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

func (a *Area) RemIntruder(i *Intruder) {
	index := slices.Index(a.Intruders, i)
	a.Intruders = slices.Delete(a.Intruders, index, index+1)
}

func (a *Area) RemPlayer(p *Player) {
	index := slices.Index(a.Players, p)
	a.Players = slices.Delete(a.Players, index, index+1)
}

func (a *Area) String() string {
	return fmt.Sprintf("%03d", a.Id)
}

func (a *Area) Corridor(n int) *Gate {
	for _, c := range a.Gates {
		if slices.Contains(c.Numbers, n) {
			return c
		}
	}
	return nil
}

type Board struct {
	Area []*Area
}

func NewBoard() (b *Board) {
	b = &Board{
		Area: areas,
	}

	corridorId := 0
	connectAreas := func(c *Corridor) {
		c.Id, corridorId = corridorId, corridorId+1
		areaX, areaY := b.Area[c.AreaX], b.Area[c.AreaY]
		areaX.Gates = append(areaX.Gates, newGate(areaY, c))
		areaY.Gates = append(areaY.Gates, newGate(areaX, c))
	}

	for _, area := range b.Area {
		if area.Class != room_1 && area.Class != room_2 {
			area.Room = &Room{
				card: newCard(area.Class),
			}
		}
	}

	for _, tunnel := range tunnels {
		connectAreas(tunnel)
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

type Corridor struct {
	Id    int
	AreaX int
	AreaY int
	Door  string
	Noise bool
	Numbers
}

type Corridors []*Corridor

func newCorridor(areaX, areaY int, numbers ...int) *Corridor {
	return &Corridor{
		AreaX:   areaX,
		AreaY:   areaY,
		Numbers: numbers,
		Door:    door_open,
	}
}

func (c *Corridor) String() string {
	return fmt.Sprintf(
		"%v%v%v",
		Noise(c.Noise),
		c.Numbers,
		c.Door,
	)
}

type Gate struct {
	*Area
	*Corridor
}

type Gates []*Gate

func newGate(a *Area, t *Corridor) *Gate {
	return &Gate{
		Area:     a,
		Corridor: t,
	}
}

func (g *Gate) String() string {
	return g.Corridor.String()
}

func (g Gates) String() (result string) {
	gates := []string{}
	for _, gate := range g {
		gates = append(gates, gate.Numbers.String())
	}
	return strings.Join(gates, "\t")
}
