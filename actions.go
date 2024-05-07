package main

type ActionBasic string

func (a ActionBasic) Name() string {
	return string(a)
}

type Action struct {
	Name   string
	Cost   Cards
	Data   map[string]interface{}
	Player *Player
}

type Actions []*Action

func (a *Action) String() string {
	return a.Name
}
