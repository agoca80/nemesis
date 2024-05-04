package main

import (
	"bufio"
	"fmt"
	"os"
)

type Human struct {
	*bufio.Reader
}

func newHuman() Controller {
	return &Human{
		Reader: bufio.NewReader(os.Stdin),
	}
}

func (h *Human) Choose(cards Cards) (selected, rejected Card) {
	if debug {
		return cards[0], cards[1]
	}

	var choice string
	for _, c := range cards {
		Show(c.Id(), c.Name())
	}
	for {
		Prompt("Choose a card")
		input, _ := fmt.Fscanln(h.Reader, &choice)
		switch {
		case input == 0:
			fallthrough
		case choice == cards[0].Id():
			return cards[0], cards[1]
		case choice == cards[1].Id():
			return cards[1], cards[0]
		}
	}
}

func (h *Human) askAction(options Actions) Action {
	option := map[int]Action{}
	for i, ca := range actions {
		option[i] = ca
		actionStr := fmt.Sprintf("%c - %s(%d)\n", 'a'+i, ca.Name(), ca.Cost())
		Show(actionStr)
	}

	for {
		Prompt("Choose an action")
		var choice string
		input, err := fmt.Fscanln(h.Reader, &choice)
		if input == 0 && err != nil {
			return nil // Human plyer passes
		}

		if input == 1 && len(choice) != 1 {
			continue
		}

		index := int(choice[0] - 'a')
		if _, ok := option[index]; ok {
			return option[index]
		}
	}
}

func (h *Human) askCost(action Action) (cost Cards) {
	if action.Cost() == 0 {
		return
	}

	for _, card := range action.Cost() {
		Show(card.Id(), card.Name())
	}

	for {
		Prompt("Choose a card")
		var choice string
		input, err := fmt.Fscanln(h.Reader, &choice)
		if input == 0 && err != nil {
			return
		}

		for _, card := range action.Cost() {
			if choice == card.Id() {
				return Cards{card}
			}
		}
	}
}

func (h *Human) NextAction(actions Actions) (action Action) {
	if debug {
		return
	}

	action = h.askAction(actions)
	if action == nil {
		return // Player passes
	}

	cost := h.askCost(action)

	return action
}
