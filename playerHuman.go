package main

import (
	"bufio"
	"fmt"
	"os"
)

func humanChoose[L ~[]E, E any](message string, options L) (selected E) {
	var stdin = bufio.NewReader(os.Stdin)
	var choice int
	Show()
	Show("PROMPT", message, ">")
	for i, o := range options {
		Show(">", i, "-", o)
	}
	for {
		fmt.Print("> ")
		input, _ := fmt.Fscanln(stdin, &choice)
		if input == 0 {
			selected = options[0]
			break
		} else if 0 <= choice && choice < len(options) {
			selected = options[choice]
			break
		}
	}
	Show()
	return
}

func chooseCharacter(cards Cards) (selected, rejected Card) {
	selected = humanChoose("Choose a card", cards)
	switch selected.Id() {
	case cards[0].Id():
		rejected = cards[1]
	case cards[1].Id():
		rejected = cards[0]
	}
	return
}

func humanChooseDirection(player *Player) (direction *Direction) {
	directions := player.Area.Directions()
	direction = humanChoose("Choose a direction", directions)
	return
}

func humanChooseAction(player *Player) (action *Action) {
	choices := player.AvailableActions()
	action = humanChoose("Choose an action", choices)
	return
}

func humanChooseIntruder(player *Player) (intruder *Intruder) {
	choices := player.Area.Intruders
	intruder = humanChoose("Choose an intruder", choices)
	return
}

func humanAction(player *Player) (action *Action) {
	if player.HandSize() < 1 {
		return
	}

	if action = humanChooseAction(player); action == nil {
		return
	}

	var name string
	var data map[string]interface{}
	switch action.Name {
	case basic_fire:
		name, data = basic_fire, map[string]interface{}{
			"intruder": humanChooseIntruder(player),
		}
	case basic_move:
		direction := humanChooseDirection(player)
		name, data = basic_move, map[string]interface{}{
			"corridor": direction.Corridor,
		}
	}

	action = &Action{
		Cost:   Cards{player.Hand[0]},
		Player: player,
		Name:   name,
		Data:   data,
	}

	return action
}

func Wait() {
	if wait {
		Show("Press Enter to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
