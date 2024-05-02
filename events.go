package main

func (g *Game) eventFailure(event *EventCard) {
	for _, a := range g.Area[A01:A21] {
		if a.IsExplored() {
			a.IsDamaged = true
		}
	}
	g.Events.Return(event)
	g.Events.Shuffle()
}

func (g *Game) ResolveEvent(event *EventCard) {
	var effects = map[string]func(*EventCard){
		event_failure: g.eventFailure,
	}

	if effect, ok := effects[event.name]; ok {
		Show("Resolving event", event.name)
		effect(event)
	} else {
		Show("PENDING effect for", event.name)
	}

	if g.Destroyed() {
		Pending("The ship has been destroyed!!!")
	}
}
