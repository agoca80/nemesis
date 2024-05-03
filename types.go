package main

import (
	"fmt"
	"sort"
)

type Actor interface {
	Dies()
}

type ExplorationTokenType string

type Damage bool

type Fire bool

type Mode int

type Noise bool

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

func (d Damage) String() string {
	if d {
		return "X"
	}
	return " "
}

func (f Fire) String() string {
	if f {
		return "X"
	}
	return " "
}

func (n Noise) String() string {
	if n {
		return "!"
	}
	return "."
}
