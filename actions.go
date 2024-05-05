package main

type actionData map[string]interface{}

type action struct {
	Action
	player *Player
	data   actionData
}

func (player *Player) ResolveMove(data actionData) {
	if player.IsInCombat() {
		Show(player, "tries to leave", player.Area)
		game.Intruders.Attack(player)
	}

	if !player.Alive() {
		return
	}

	corridor := data["corridor"].(*Corridor)
	destination := corridor.End(player.Area)
	noiseRoll := player.MovesTo(destination)

	var event string
	if !player.Area.IsExplored() {
		event = game.ResolveExploration(player, corridor)
	}

	if game.Destroyed() {
		return
	}

	if noiseRoll && event != ev_danger && event != ev_silence {
		player.ResolveNoise()
	}
}

func (player *Player) ResolveFire(actionData actionData) {
	intruder := actionData["intruder"].(*Intruder)

	var damage int
	var roll = player.RollDamage()
	switch {
	case roll == damage_double:
		damage = 2
	case roll == damage_single:
		damage = 1
	case roll == intruder.Kind:
		damage = 1
	case roll == intruder_adult && intruder.Kind == intruder_crawler:
		damage = 1
	case roll == intruder_adult && intruder.Kind == intruder_larva:
		damage = 1
	case roll == intruder_adult && intruder.Kind == intruder_egg:
		damage = 1
	case roll == intruder_blank:
		damage = 0
	}

	Show(player, "opens fire against", intruder, ", rolls", roll, ", deals", damage, "damage")
	if intruder.Suffers(damage) {
		Show(intruder, "dies!")
	}
}
