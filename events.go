package main

func eventFailure(event *EventCard) {
	for _, a := range game.Area[A01:A21] {
		if a.IsExplored() {
			a.IsDamaged = true
		}
	}
	game.Events.Return(event)
	game.Events.Shuffle()
}

func ResolveEvent(event *EventCard) {
	var effects = map[string]func(*EventCard){
		event_failure: eventFailure,
	}

	if effect, ok := effects[event.name]; ok {
		Show("Resolving event", event.name)
		effect(event)
	} else {
		Show("PENDING effect for", event.name)
	}

	if game.Destroyed() {
		Pending("The ship has been destroyed!!!")
	}
}
