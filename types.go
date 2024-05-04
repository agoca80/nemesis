package main

import (
	"fmt"
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
