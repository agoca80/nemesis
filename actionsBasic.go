package main

type ActionBasic string

func (a ActionBasic) Name() string {
	return string(a)
}

func (a ActionBasic) Cost() int {
	cost := map[ActionBasic]int{
		basic_move:     1,
		basic_fire:     1,
		basic_fight:    1,
		basic_pickup:   1,
		basic_exchange: 1,
		basic_prepare:  1,
		basic_sneak:    1,
	}

	return cost[a]
}

func (a ActionBasic) Resolve(actionData actionData) {
	player := actionData["player"].(*Player)
	resolve := map[string]func(actionData){
		basic_move: player.ResolveMove,
		basic_fire: player.ResolveFire,
	}

	resolve[string(a)](actionData)
}
