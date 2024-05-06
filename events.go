package main

func eventFailure(event *Event) {
	for _, a := range ship.Area[A01:A21] {
		if a.IsExplored() {
			a.IsDamaged = true
			break
		}
	}
	events.Return(event)
	events.Shuffle()
}

func eventHatching(event *Event) {
	availableCrawlers := 3
	for _, i := range game.Intruders {
		if i.Kind == intruder_crawler {
			availableCrawlers--
		}
	}

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
