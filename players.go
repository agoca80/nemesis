package main

import (
	"slices"
	"strconv"
)

func (game *Game) AvailableActions(p *player) (actions Actions) {
	actions = Actions{
		ActionBasic(basic_move),
		ActionBasic(basic_fire),
	}

}

type Controller interface {
	Choose(cards Cards) (selected, rejected Card)
	NextAction(Actions) Action
}

type player struct {
	*Area
	*Deck
	*HelpCard

	Controller

	Bruises   int
	Id        int
	Goals     Cards
	Hand      Cards
	Character string
	Jonesy    bool
	Wounds    Cards

	IsDrenched Issue
	IsInfected Issue
	Signaled   bool
	State      string
}

type Players []*player

var playerId = 0

func NewPlayer(kind string) *player {
	playerId++
	var newController func() Controller
	if kind == "human" {
		newController = newHuman
	} else {
		newController = newDummy
	}

	return &player{
		Controller: newController(),
		Id:         playerId,
		Goals:      Cards{},
		Hand:       Cards{},
		Wounds:     Cards{},
		State:      player_alive,
	}
}

func (p *player) ChooseCharacter() (rejected Card) {
	characters := Cards{
		characters.Random(),
		characters.Random(),
	}
	selected, rejected := p.Choose(characters)
	p.Character = selected.Name()
	Show("Player", p.Id, "picks", selected, "between", characters)
	return
}

func (p *player) String() string {
	return p.Character
}

func (p *player) HandCapacity() int {
	if p.HasWound(wound_body) {
		return 4
	} else {
		return 5
	}
}

func (p *player) HandSize() int {
	return len(p.Hand)
}

func (p *player) DrawActions() {
	for p.HandSize() < p.HandCapacity() {
		p.Hand = append(p.Hand, p.Draw())
	}
}

func (p *player) Passes() {
	Show(p, "passes")
	p.Flips()
	if p.Area.IsInFire {
		p.SuffersLightWound()
	}

	if p.HasWound(wound_hemorrhage) {
		p.SuffersLightWound()
	}
}

func (p *player) HasPassed() bool {
	return p.Flipped
}

func (p *player) SuffersContamination() {
	Show(p, "suffers contamination!")
	p.Discard(contamination.Draw())
}

func (p *player) SuffersLightWound() {
	if len(p.Wounds) == 3 {
		p.Dies()
		return
	}

	p.Bruises++
	if p.Bruises == 3 {
		p.Bruises = 0
		p.SufferSeriousWound()
	} else {
		Show(p, "suffers a light wound")
	}
}

func (p *player) SufferSeriousWound() {
	if len(p.Wounds) == 3 {
		p.Dies()
		return
	}

	card := wounds.Draw().(*SeriousWound)
	p.Wounds = append(p.Wounds, card)
	Show(p, "suffers", card.name, "!")
}

func (p *player) IsSurprised(token *IntruderToken) bool {
	return p.HandSize() < token.Alert
}

func (p *player) IsInCombat() bool {
	return len(p.Area.Intruders) > 0
}

func (p *player) HasWound(name string) bool {
	return slices.ContainsFunc(p.Wounds, func(w Card) bool {
		return w.Name() == name && !w.(*SeriousWound).isDressed
	})
}

func (p *player) RollNoise() (result string) {
	noiseDice := Symbols{
		silence,
		danger,
		n1,
		n1,
		n2,
		n2,
		n3,
		n3,
		n4,
		n4,
	}

	result = Roll(noiseDice)
	if p.IsDrenched && result == silence {
		result = danger
	}
	return
}

func (p Players) PassJonesy() (new int) {
	var old int
	for i := range p {
		if p[i].Jonesy {
			old = i
			break
		}
	}

	for i := range p {
		new = (old + i + 1) % len(p)
		if p[new].Alive() {
			p[old].Jonesy = false
			p[new].Jonesy = true
			break
		}
	}
	Show(p[old], "passes Jonesy to", p[new])
	return new
}

func (p *player) Dies() {
	p.State = player_dead
	Show(p, "dies!")
}

func (p *player) Alive() bool {
	return p.State == player_alive
}

func (p Players) Alive() (players Players) {
	for _, player := range p {
		if player.Alive() {
			players = append(players, player)
		}
	}
	return
}

func (p *player) GoingOn() bool {
	return p.State == player_alive && !p.HasPassed()
}

func (p Players) GoingOn() (players Players) {
	for _, player := range p {
		if player.GoingOn() {
			players = append(players, player)
		}
	}
	return
}

func (p *player) ResolveNoise() {
	encounter := false
	result := p.RollNoise()
	if result == ev_silence && p.IsDrenched {
		Show(p, "was silent but is drenched in mucus!")
		result = ev_danger
	}

	switch result {
	case ev_silence:
		Show(p, "is silent...")
	case ev_danger:
		Show(p, "is in danger!")
		p.Danger()
	default:
		n, _ := strconv.Atoi(result)
		corridor := p.Corridor(n)
		Show(p, "makes noise in corridor", corridor.Numbers)
		if corridor.Noise {
			for _, c := range p.Area.Corridors {
				c.Noise = false
			}
			encounter = true
		} else {
			corridor.Noise = true
		}
	}

	if encounter {
		game.ResolveEncounter(p)
	}
}
