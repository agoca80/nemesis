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

func ResolveEvent(event *Event) {
	var effects = map[string]func(*Event){
		event_failure: eventFailure,
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
