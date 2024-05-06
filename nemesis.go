package main

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

type Ship struct {
	Area              []*Area
	hibernatoriumOpen bool
}

func newShip() (b *Ship) {
	b = &Ship{
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

func (b *Ship) Damages() (damaged int) {
	for _, area := range b.Area {
		if area.IsDamaged {
			damaged++
		}
	}
	return
}

func (b *Ship) Fires() (result int) {
	for _, area := range b.Area {
		if area.IsInFire {
			result++
		}
	}
	return
}

func (b *Ship) Destroyed() bool {
	return b.Damages() > 8 || b.Fires() > 8
}
