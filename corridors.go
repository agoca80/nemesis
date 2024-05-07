package main

import "fmt"

type Corridor struct {
	Id    int
	AreaX *Area
	AreaY *Area
	Door  string
	Noise Issue
	Numbers
}

func (c *Corridor) Danger() {
	c.Noise = true
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

func (corridors Corridors) Show() (result string) {
	doors := ""
	noise := ""
	for _, corridor := range corridors {
		doors += corridor.Door
		noise += corridor.Noise.String()
	}
	result = fmt.Sprintf(
		"%v\t%v",
		noise,
		doors,
	)
	return
}
