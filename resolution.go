package main

import "slices"

type Action interface {
	Name() string
	Cost() int
	Resolve(map[string]interface{})
}

func (player *Player) MovesTo(dstArea *Area) (moiseRoll bool) {
	moiseRoll = dstArea.IsEmpty()
	srcArea := player.Area
	srcArea.RemPlayer(player)
	player.Area, dstArea.Players = dstArea, append(dstArea.Players, player)
	Show(player, "moves to", dstArea)
	return
}

func (game *Game) ResolveMove(player *Player, corridor *Corridor) {
	if player.IsInCombat() {
		Show(player, "tries to leave", player.Area)
		player.Area.Intruders.Attack(player)
	}

	if !player.Alive() {
		return
	}

	noiseRoll := player.MovesTo(corridor.Area)

	var event string
	if !player.Area.IsExplored() {
		event = game.ResolveExploration(player, corridor)
	}

	if game.Destroyed() {
		return
	}

	if noiseRoll && event != ev_danger && event != ev_silence {
		game.ResolveNoise(player)
	}
}

func (a ActionBasic) Resolve(data map[string]interface{}) {
	player := data["player"].(*Player)
	switch string(a) {
	case basic_move:
		corridor := data["corridor"].(*Corridor)
		player.ResolveMove(player, corridor)
	default:
		Pending(a, "not implemented")
	}
}

func (p *Player) GetAction() (actionData map[string]interface{}) {
	if p.HandSize() < 1 {
		return
	}

	return map[string]interface{}{
		"player":   p,
		"action":   ActionBasic(basic_move),
		"corridor": p.chooseCorridor(),
		"cost": Cards{
			p.Hand[0],
		},
	}
}

func (p *Player) Pay(card Card) {
	index := slices.Index(p.Hand, card)
	if index == -1 {
		panic("WTF")
	}
	p.Hand = slices.Delete(p.Hand, index, index+1)
	p.Discard(card)
}

func (game *Game) GetAction(player *Player) {
	actionData := player.GetAction()
	if actionData == nil {
		player.Passes()
		return
	}

	cost := actionData["cost"].(Cards)
	for _, card := range cost {
		player.Pay(card)
	}

	action := actionData["action"].(Action)
	action.Resolve(actionData)
	player.Show()
	game.Show()
	Wait()
}
