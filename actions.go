package main

type ActionBasic string

func (a ActionBasic) Name() string {
	return string(a)
}

func NewActionCard(cost int, name string) *ActionCard {
	return &ActionCard{
		card: newCard(name),
		cost: cost,
	}
}

func (ac *ActionCard) Cost() int {
	return ac.cost
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

type Action struct {
	Name   string
	Cost   Cards
	Data   map[string]interface{}
	Player *Player
}
