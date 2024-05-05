package main

import (
	"fmt"
	"sort"
)

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
		basic_sneak:    2,
	}

	return cost[a]
}

type Actor interface {
	Dies()
}

type ExplorationTokenType string

type Issue bool

type Mode int

type Numbers []int

func (n Numbers) String() string {
	sort.Ints(n)
	total := 0
	for _, d := range n {
		total = total*10 + d
	}
	return fmt.Sprintf("%02d", total)
}

type Step string

func (i Issue) String() string {
	if i {
		return "!"
	}
	return "."
}
