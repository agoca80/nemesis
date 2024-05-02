package main

import "slices"

func (g *Game) draw() Step {
	for _, p := range g.Players {
		p.DrawActions()
	}
	return step_token
}

func (g *Game) token() Step {
	if g.hyperdriveCountdown < 15 {
		g.PassJonesy()
	}

	var first int
	for i, p := range g.Players {
		if p.Jonesy {
			first = i
			break
		}
	}

	players := Players{}
	for i := range g.Players {
		next := (first + i) % len(g.Players)
		players = append(players, g.Players[next])
	}
	g.Players = players
	players.Alive().Show()
	return step_turn
}

func (game *Game) turn() Step {
	players := game.Players.Alive()
	for _, p := range players {
		p.Unflips()
	}

	for {
		for _, player := range players {
			Show()
			Show("TURN", player)
			game.AskAction(player)
			if game.GoingOn() && player.GoingOn() {
				game.AskAction(player)
			}
		}

		if game.gameOver() {
			break
		}

		players = players.GoingOn()
		if len(players) == 0 {
			break
		}
	}

	return step_counters
}

func (g *Game) counters() Step {
	g.hyperdriveCountdown--
	Show("Hyperdrive countdown:", g.hyperdriveCountdown)

	if g.hyperdriveCountdown == 8 {
		Pending("The hibernatorium chambers are open again!")
	}

	return step_attack
}

func (g *Game) attacks() Step {
	for _, i := range g.Intruders {
		if i.IsInCombat() {
			i.Attack()
		}
	}

	return step_fire
}

func (g *Game) fireDamage() Step {
	for _, a := range g.Area {
		if !a.IsInFire {
			continue
		}

		for _, intruder := range a.Intruders {
			Show(intruder, "in", a, "is damaged by fire!")
			g.FireDamage(intruder, 1)
		}
	}
	return step_event
}

func (g *Game) event() Step {
	event := g.Events.Draw().(*EventCard)
	Show("Event card is", event.name)
	for _, i := range g.Intruders {
		if slices.Contains(event.Symbols, i.Kind) && !i.IsInCombat() {
			i.Moves(event.Corridor)
		}
	}

	g.ResolveEvent(event)

	return step_evolution
}

func (g *Game) evolution() Step {
	rollNoise := func() {
		Show("All players roll noise in turn order")
		for _, p := range g.Players.Alive() {
			if !p.IsInCombat() {
				p.RollNoise()
			}
		}
	}
	token := g.FetchToken()
	kind := token.Kind
	switch kind {
	case token_blank:
		Show("More adults are lurking on the ship")
		g.Return(token_adult)
	case token_larva:
		Show("A larva grows into an adult")
		g.Retire(token)
		g.Return(token_adult)
	case token_crawler:
		Show("A crawler grows into a breeder")
		g.Retire(token)
		g.Return(token_breeder)
	case token_adult:
		rollNoise()
	case token_breeder:
		rollNoise()
	case token_queen:
		var nest *Area
		for _, a := range g.Area {
			if a.Name() == room_nest {
				break
			}
		}

		if nest == nil || len(nest.Players) == 0 {
			Show("The queen lays another egg!")
			g.Eggs++
		} else {
			Show("The queen show its might!")
			intruder := g.NewIntruder(token, nest)
			intruder.Attack()
		}
	}

	return step_draw
}
