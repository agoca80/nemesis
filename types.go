package main

type Actor interface {
	Dies()
}

type ExplorationTokenType string

type Damage bool

type Fire bool

type Mode int

type Noise bool

type Step string

func (d Damage) String() string {
	if d {
		return "D"
	}
	return " "
}

func (f Fire) String() string {
	if f {
		return "F"
	}
	return " "
}

func (n Noise) String() string {
	if n {
		return "!"
	}
	return " "
}
