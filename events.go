package main

func eventFailure(event *Event) {
	explored := Filter(ship.Area, (*Area).IsExplored)
	Each(explored, (*Area).Damage)
	events.Return(event)
	events.Shuffle()
}

func eventHatching(event *Event) {
	availableCrawlers := 3 - len(Filter(intruders, func(i *Intruder) bool {
		return i.Kind == intruder_crawler && i.Alive()
	}))

	for _, player := range players.Alive() {
		switch {
		case player.IsInfected && availableCrawlers > 0:
			player.Dies()
			newIntruder(intruder_crawler, player.Area)
			availableCrawlers--
		case player.IsInfected:
			player.Dies()
		case !player.IsInfected:
			for range 4 {
				card := player.Draw()
				player.IsInfected = card.Name() == "contamination" && card.(*Contamination).Infected
				player.Discard(card)
			}
		}
	}
}

func ResolveEvent(event *Event) {
	var effects = map[string]func(*Event){
		event_failure:  eventFailure,
		event_hatching: eventHatching,
	}

	if effect, ok := effects[event.name]; ok {
		Show("Resolving event", event.name)
		effect(event)
	} else {
		Show("PENDING effect for", event.name)
	}

	if ship.Destroyed() {
		Show("The ship has been destroyed!!!")
	}
}
