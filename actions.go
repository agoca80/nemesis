package main

type Action interface {
	Name() string
	Resolve()
}

type Actions []Action
