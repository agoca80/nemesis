package main

import (
	"fmt"
	"slices"
)

type Player struct {
	*Area
	*Deck
	*HelpCard

	Bruises   int
	Character Card
	Goals     Cards
	Hand      Cards
	Id        string
	Jonesy    bool
	Wounds    Cards

	IsDrenched Issue
	IsInfected bool
	IsHuman    bool
	Signaled   bool
	State      string
}

type Players []*Player

var playerId = 0

func NewPlayer(isHuman bool) *Player {
	playerId++
	return &Player{
		Id:      fmt.Sprintf("p%d", playerId),
		Goals:   Cards{},
		Hand:    Cards{},
		IsHuman: isHuman,
		Wounds:  Cards{},
		State:   player_alive,
	}
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
	actionCards := Filter(p.Hand, func(c Card) bool {
		return c.Name() != "contamination"
	})
	return len(actionCards)
}

func (p *Player) DrawActions() {
	for p.HandSize() < p.HandCapacity() {
		p.Hand = append(p.Hand, p.Draw())
	}
}

func (p *Player) Passes() {
	Show(p, "passes")
	p.Flips()
	if p.Area.IsBurning() {
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

	card := wounds.Draw().(*Wound)
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
		return w.Name() == name && !w.(*Wound).isDressed
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
	p.State = player_dead
	Show(p, "dies!")
}

func (p *Player) Alive() bool {
	return p.State == player_alive
}

func (p *Player) GoingOn() bool {
	return p.State == player_alive && !p.HasPassed()
}

func (player *Player) NextAction() {
	nextAction := map[bool]func(*Player) *Action{
		false: dummyAction,
		true:  humanAction,
	}

	if gameOver() || !player.GoingOn() {
		return
	}

	action := nextAction[player.IsHuman](player)
	if action == nil {
		player.Passes()
		return
	}

	resolveAction(action)
	Show()
	ship.Show()
	player.Describe()
	Wait()
}

func (p *Player) RollDamage() (result string) {
	damageDice := Symbols{
		damage_blank,
		damage_crawler,
		damage_crawler,
		damage_adult,
		damage_single,
		damage_double,
	}

	return Roll(damageDice)
}

func (player *Player) MovesTo(dstArea *Area) (moiseRoll bool) {
	moiseRoll = dstArea.IsEmpty()
	srcArea := player.Area
	srcArea.RemPlayer(player)
	player.Area, dstArea.Players = dstArea, append(dstArea.Players, player)
	Show(player, "moves to", dstArea)
	return
}

func (p *Player) Describe() string {
	return fmt.Sprintf("%v\t(%v)\t%v\t%v+%v\tHand %v",
		p.Character,
		Issue(p.IsInfected),
		p.State,
		p.Bruises,
		p.Wounds,
		p.Hand,
	)
}

func (p Players) Alive() Players {
	return Filter(p, (*Player).Alive)
}

func (p Players) GoingOn() Players {
	return Filter(p, (*Player).GoingOn)
}

func (player *Player) Choose(options Cards) (selected, rejected Card) {
	switch player.IsHuman {
	case true:
		return chooseCharacter(options)
	default:
		return dummyChoose(options)
	}
}

func (player *Player) AvailableActions() (actions Actions) {
	actions = Actions{}
	switch {
	case player.IsInCombat():
		actions = append(actions, Actions{
			&Action{Name: basic_escape},
			&Action{Name: basic_fight},
			&Action{Name: basic_fire},
		}...)
	case !player.IsInCombat():
		actions = append(actions, Actions{
			&Action{Name: basic_move},
			&Action{Name: basic_careful},
		}...)
	}

	return
}
