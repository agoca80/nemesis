package main

import "slices"

func (g *Game) draw() Step {
	for _, p := range players.Alive() {
		p.DrawActions()
	}
	return step_token
}

func (g *Game) token() Step {
	if game.hyperdriveCountdown < 15 {
		players.PassJonesy()
	}

	var first int
	for i, p := range players {
		if p.Jonesy {
			first = i
			break
		}
	}

	players = append(players[first:], players[:first]...)
	players.Alive().Show()
	return step_turn
}

func (g *Game) stepTurn() Step {
	for _, p := range players.Alive() {
		p.Unflips()
	}

	for {
		if gameOver() {
			break
		}

		goingOn := players.GoingOn()
		if len(goingOn) == 0 {
			break
		}

		for _, player := range goingOn {
			player.NextAction()
			player.NextAction()
		}
	}

	return step_counters
}

func (g *Game) counters() Step {
	game.hyperdriveCountdown--
	Show("Hyperdrive countdown:", game.hyperdriveCountdown)

	if game.hyperdriveCountdown == 8 {
		Pending("The hibernatorium chambers are open again!")
	}

	return step_attack
}

func (g *Game) stepAttacks() Step {
	for _, i := range game.Intruders {
		if i.IsInCombat() {
			i.Attack()
		}
	}

	return step_fire
}

func (g *Game) fireDamage() Step {
	for _, a := range ship.Area {
		if !a.IsInFire {
			continue
		}

		for _, intruder := range a.Intruders {
			intruder.FireDamage(1)
		}
	}
	return step_event
}

func (g *Game) event() Step {
	event := events.Draw().(*Event)
	Show("Event card is", event.name)
	for _, i := range game.Intruders {
		if slices.Contains(event.Symbols, i.Kind) && !i.IsInCombat() {
			i.Moves(event.Corridor)
		}
	}

	ResolveEvent(event)

	return step_evolution
}

func (g *Game) evolution() Step {
	rollNoise := func() {
		Show("All non-combating players roll noise in turn order")
		for _, p := range players.Alive() {
			if !p.IsInCombat() {
				p.RollNoise()
			}
		}
	}
	token := game.FetchToken()
	kind := token.Kind
	switch kind {
	case token_blank:
		Show("More adults are lurking on the ship")
		game.Return(token_adult)
	case token_larva:
		Show("A larva grows into an adult")
		game.Retire(token)
		game.Return(token_adult)
	case token_crawler:
		Show("A crawler grows into a breeder")
		game.Retire(token)
		game.Return(token_breeder)
	case token_adult:
		rollNoise()
	case token_breeder:
		rollNoise()
	case token_queen:
		var nest *Area
		for _, a := range ship.Area {
			if a.Name() == room_nest {
				nest = a
				break
			}
		}

		if nest == nil || len(nest.Players) == 0 {
			Show("The queen lays another egg!")
			game.Eggs++
		} else {
			Show("The queen show its might!")
			intruder := spawnIntruder(token, nest)
			intruder.Attack()
		}
	}

	return step_draw
}

func finish() {
	Show()
	Show("GAME OVER!!!")
	Show()
	ship.Show()
	players.Show()
}
