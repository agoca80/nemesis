package main

import "fmt"

type ActionCard struct {
	*card
	cost int
}

type Attack struct {
	*card
	Wounds int
	Symbols
}

type Contamination struct {
	*card
	Infected Issue
}

func newContamination(infected bool) *Contamination {
	return &Contamination{
		card:     newCard("contamination"),
		Infected: Issue(infected),
	}
}

func (c *Contamination) Reveal() string {
	if c.Infected {
		return "infected"
	} else {
		return "clean"
	}
}

type Coordinates struct {
	*card
	coordinates string
}

type Event struct {
	*card
	Corridor int
	Symbols
}

type ExplorationToken struct {
	*card
	Items int    // Number of items in the room
	Event string // Exploration event kind
}

type Goal struct {
	*card
	Number int
	Kind   string
	eval   func(*Game) bool
}

type IntruderToken struct {
	*card
	Alert int
	Kind  string
}

type IntruderTokens []*IntruderToken

type Item struct {
	*card
	Color     string
	SingleUse bool
	Cost      int
}

type Room struct {
	*card
	Color            string
	Computer         bool
	ExplorationToken *ExplorationToken
	Items            int
}

type Weakness struct {
	*card
	Revealed bool
}

func newAttack(wounds int, name string, symbols ...string) *Attack {
	return &Attack{
		card:    newCard(name),
		Wounds:  wounds,
		Symbols: symbols,
	}
}

func newCoordinates(coordinates string) *Coordinates {
	return &Coordinates{
		card:        newCard("coordinates"),
		coordinates: coordinates,
	}
}

func (c *Coordinates) String() (str string) {
	destinationName := []string{
		"Deep Space",
		"Earth",
		"Mars",
		"Venus",
	}
	str = fmt.Sprintf("A: %-10s B: %-10s C: %-10s D: %-10s",
		destinationName[c.coordinates[0]-'A'],
		destinationName[c.coordinates[1]-'A'],
		destinationName[c.coordinates[2]-'A'],
		destinationName[c.coordinates[3]-'A'],
	)

	return
}

func newEvent(corridor int, name string, symbols ...string) *Event {
	return &Event{
		card:     newCard(name),
		Corridor: corridor,
		Symbols:  symbols,
	}
}

func newExplorationToken(items int, event string) *ExplorationToken {
	return &ExplorationToken{
		card:  newCard("exploration token"),
		Items: items,
		Event: event,
	}
}

func (et *ExplorationToken) String() (str string) {
	if et != nil {
		str = fmt.Sprintf("%s-%d", et.Event, et.Items)
	}
	return
}

func NewIntruderToken(alert int, kind string) *IntruderToken {
	return &IntruderToken{
		card:  newCard(kind),
		Alert: alert,
		Kind:  kind,
	}
}

func (it *IntruderToken) String() string {
	return it.name
}

func newGoal(players int, name string) *Goal {
	return &Goal{
		card:   newCard(name),
		Number: players,
		eval: func(game *Game) bool {
			Show("PENDING Goal.Eval")
			return false
		},
	}
}

func newRoom(name, color string, computer bool) *Room {
	return &Room{
		card:     newCard(name),
		Color:    color,
		Computer: computer,
	}
}

type Wound struct {
	*card
	isDressed bool
}

type Wounds []*Wound

func newWound(name string) *Wound {
	return &Wound{
		card: newCard(name),
	}
}

func newWeakness(name string) *Weakness {
	return &Weakness{
		card: newCard(name),
	}
}
