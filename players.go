package main

import (
	"fmt"
	"slices"
	"strconv"
)

func (player *Player) Name() string {
	return player.Character.Name()
}

func (player *Player) ChooseCharacter(options Cards) (selected, rejected Card) {
	selected, rejected = player.Choose(options)
	player.Character = selected.(*Character)
	Debug(fmt.Sprintf("Player %d picks %-9s and rejects %-9s", player.Id, selected, rejected))
	return
}

type PlayerInput interface {
	Choose(Cards) (selected, rejected Card)
}

type Player struct {
	PlayerInput

	*Area
	*Character
	*Deck
	*HelpCard

	Bruises int
	Id      int
	Goals   Cards
	Hand    Cards
	Jonesy  bool
	Wounds  Cards

	HasSlime   Issue
	IsInfected Issue
	Signaled   bool
	State      string
}

type Players []*Player

var playerId = 0

func newPlayer(human bool) (player *Player) {
	playerId++
	player = &Player{
		Id:     playerId,
		Goals:  Cards{},
		Hand:   Cards{},
		Wounds: Cards{},
		State:  player_alive,
	}

	if human {
		player.PlayerInput = newHuman()
	} else {
		player.PlayerInput = newDummy()
	}

	return player
}

func (p *Player) String() string {
	return p.Character.Name()
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
	p.Discard(contamination.Draw())
}

func (p *Player) SuffersLightWound() {
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

func (p *Player) SufferSeriousWound() {
	if len(p.Wounds) == 3 {
		p.Dies()
		return
	}

	card := wounds.Draw().(*SeriousWound)
	p.Wounds = append(p.Wounds, card)
	Show(p, "suffers", card.name, "!")
}

func (p *Player) IsSurprised(token *IntruderToken) bool {
	return p.HandSize() < token.Alert
}

func (p *Player) IsInCombat() bool {
	return len(p.Area.Intruders) > 0
}

func (p *Player) HasWound(name string) bool {
	return slices.ContainsFunc(p.Wounds, func(w Card) bool {
		return w.Name() == name && !w.(*SeriousWound).isDressed
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
	if p.HasSlime && result == silence {
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
	p.State = player_dead
	Show(p, "dies!")
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
	if result == ev_silence && p.HasSlime {
		Show(p, "was silent but is covered in slime!")
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
		ResolveEncounter(p)
	}
}

func (player *Player) MovesTo(dstArea *Area) (noiseRoll bool) {
	noiseRoll = dstArea.IsEmpty()
	srcArea := player.Area
	srcArea.RemPlayer(player)
	player.Area, dstArea.Players = dstArea, append(dstArea.Players, player)
	Show(player, "moves to", dstArea)
	return
}
