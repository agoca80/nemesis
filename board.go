package main

import (
	"fmt"
	"slices"
	"strings"
)

type Numbers []int

type Area struct {
	*Room

	Corridors
	Intruders
	Players

	Class     string
	Id        int
	IsDamaged bool
	IsInFire  bool
}

type Areas []*Area

type Corridor struct {
	*Area
	*Tunnel
}

type Corridors []*Corridor

type Tunnel struct {
	Id int
	Numbers
	Door  string
	Noise bool
}

type Tunnels []*Tunnel

type Board struct {
	Area   []*Area
	Tunnel []*Tunnel
}

func (n Numbers) String() string {
	total := 0
	for _, number := range n {
		total = total*10 + number
	}
	return fmt.Sprintf("%02d", total)
}

func (a *Area) RemPlayer(p *Player) {
	index := slices.Index(a.Players, p)
	a.Players = slices.Delete(a.Players, index, index+1)
}

func (a *Area) RemIntruder(i *Intruder) {
	index := slices.Index(a.Intruders, i)
	a.Intruders = slices.Delete(a.Intruders, index, index+1)
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

func (a *Area) String() string {
	return fmt.Sprintf("A%02d", a.Id)
}

func (a *Area) Corridor(n int) *Corridor {
	for _, c := range a.Corridors {
		if slices.Contains(c.Numbers, n) {
			return c
		}
	}
	return nil
}

func area(id int, class string) *Area {
	return &Area{
		Id:        id,
		Class:     class,
		Players:   Players{},
		Intruders: Intruders{},
	}
}

func corridor(a *Area, t *Tunnel) *Corridor {
	return &Corridor{
		Area:   a,
		Tunnel: t,
	}
}

func (c *Corridor) String() string {
	return c.Area.String()
}

func (c Corridors) String() (result string) {
	corridors := Symbols{}
	for _, corridor := range c {
		corridors = append(corridors, corridor.String())
	}
	return strings.Join(corridors, "\t")
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
