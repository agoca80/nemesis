package main

import "fmt"

type ActionCard struct {
	*card
	cost int
}

type AttackCard struct {
	*card
	Wounds int
	Symbols
}

type ContaminationCard struct {
	*card
	bool
}

type Coordinates struct {
	*card
	coordinates string
}

type EventCard struct {
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

func NewAttackCard(wounds int, name string, symbols ...string) *AttackCard {
	return &AttackCard{
		card:    newCard(name),
		Wounds:  wounds,
		Symbols: symbols,
	}
}

func contaminationCard(infected bool) *ContaminationCard {
	return &ContaminationCard{
		card: newCard("contamination"),
		bool: infected,
	}
}

func CoordinatesCard(coordinates string) *Coordinates {
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

func NewEventCard(corridor int, name string, symbols ...string) *EventCard {
	return &EventCard{
		card:     newCard(name),
		Corridor: corridor,
		Symbols:  symbols,
	}
}

func NewExplorationToken(items int, event string) *ExplorationToken {
	return &ExplorationToken{
		card:  newCard("exploration token"),
		Items: items,
		Event: event,
	}
}

func (et *ExplorationToken) String() (str string) {
	if et != nil {
		str = fmt.Sprintf("%d-%s", et.Items, et.Event)
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

func NewGoal(players int, name string) *Goal {
	return &Goal{
		card: newCard(name),
		eval: func(game *Game) bool {
			Show("PENDING Goal.Eval")
			return false
		},
	}
}

func NewRoom(name, color string, computer bool) *Room {
	return &Room{
		card:     newCard(name),
		Color:    color,
		Computer: computer,
	}
}

type SeriousWound struct {
	*card
	isDressed bool
}

type SeriousWounds []*SeriousWound

func NewSeriousWound(name string) *SeriousWound {
	return &SeriousWound{
		card: newCard(name),
	}
}

func NewWeakness(name string) *Weakness {
	return &Weakness{
		card: newCard(name),
	}
}
