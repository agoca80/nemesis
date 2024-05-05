package main

import (
	"fmt"
	"slices"
	"sort"
)

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

type Object struct {
	Area *Area
	Name string
}

type Objects []*Object

func (o *Object) String() string {
	return o.Name
}

type Symbols []string

func symbols(symbols ...string) Symbols {
	return symbols
}

func (s Symbols) Contains(symbol string) bool {
	return slices.Contains(s, symbol)
}
