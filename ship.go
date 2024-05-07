package main

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

func (b *Ship) Destroyed() bool {
	fires := len(Filter(b.Area, (*Area).IsBurning))
	damages := len(Filter(b.Area, (*Area).IsDamaged))
	return fires > 8 || damages > 8
}
