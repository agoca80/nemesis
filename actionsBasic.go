package main

type ActionBasic struct {
	name string
}

func newActionBasic(name string) *ActionBasic {
	return &ActionBasic{name: name}
}

func (a *ActionBasic) Name() string {
	return a.name
}

func (a *ActionBasic) Cost() int {
	cost := map[string]int{
		basic_move:     1,
		basic_fire:     1,
		basic_fight:    1,
		basic_pickup:   1,
		basic_exchange: 1,
		basic_prepare:  1,
		basic_sneak:    1,
	}

	return cost[a.name]
}

func (a *ActionBasic) Resolve(player *Player, d data) {
	resolve := map[string]func(data){
		basic_move: player.ResolveMove,
		basic_fire: player.ResolveFire,
	}

	resolve[a.name](d)
}
