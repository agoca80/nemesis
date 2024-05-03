package main

import (
	"slices"
	"strconv"
)

type Player struct {
	*Area
	*Deck
	*HelpCard

	Id            int
	Hand          Cards
	Character     string
	Jonesy        bool
	LightWounds   int
	SeriousWounds SeriousWounds

	IsDrenched bool
	IsInfested bool
	Goals      Cards
	Signaled   bool
	State      string
}

type Players []*Player

var playerId = 0

func NewPlayer() *Player {
	playerId++
	return &Player{
		Id:            playerId,
		Goals:         Cards{},
		Hand:          Cards{},
		SeriousWounds: SeriousWounds{},
		State:         player_alive,
	}
}

func (p *Player) String() string {
	return p.Character
}

func (p *Player) HandCapacity() int {
	if p.HasWound(wound_body) {
		return 4
	} else {
		return 5
	}
}

func (p *Player) HandSize() int {
	return len(p.Hand)
}

func (p *Player) DrawActions() {
	for p.HandSize() < p.HandCapacity() {
		p.Hand = append(p.Hand, p.Draw())
	}
}

func (p *Player) Passes() {
	Show(p, "passes")
	p.Flips()
	if p.Area.IsInFire {
		p.SuffersLightWound()
	}

	if p.HasWound(wound_hemorrhage) {
		p.SuffersLightWound()
	}
}

func (p *Player) HasPassed() bool {
	return p.Flipped
}

func (p *Player) SuffersContamination() {
	Show(p, "suffers contamination!")
	p.Discard(game.Contamination.Draw())
}

func (p *Player) SuffersLightWound() {
	if len(p.SeriousWounds) == 3 {
		p.Dies()
		return
	}

	p.LightWounds++
	if p.LightWounds == 3 {
		p.LightWounds = 0
		p.SufferSeriousWound()
	} else {
		Show(p, "suffers a light wound")
	}
}

func (p *Player) SufferSeriousWound() {
	if len(p.SeriousWounds) == 3 {
		p.Dies()
		return
	}

	card := game.Wounds.Draw().(*SeriousWound)
	p.SeriousWounds = append(p.SeriousWounds, card)
	Show(p, "suffers", card.name, "!")
}

func (p *Player) IsSurprised(token *IntruderToken) bool {
	return p.HandSize() < token.Alert
}

func (p *Player) IsInCombat() bool {
	return len(p.Area.Intruders) > 0
}

func (p *Player) HasWound(name string) bool {
	return slices.ContainsFunc(p.SeriousWounds, func(w *SeriousWound) bool {
		return w.Name() == name && !w.isDressed
	})
}

func (p *Player) RollNoise() (result string) {
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

func (p *Player) Dies() {
	Show(p, "dies!")
	p.State = player_dead
}

func (p *Player) Alive() bool {
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

func (p *Player) GoingOn() bool {
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

func (p *Player) ResolveNoise() {
	encounter := false
	result := p.RollNoise()
	if result == ev_silence && p.IsDrenched {
		Show(p, "was silent but is drenched in mucus")
		result = ev_danger
	}

	switch result {
	case ev_silence:
		Show(p, "makes no noise")
	case ev_danger:
		Show(p.Area, "is in danger!")
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
