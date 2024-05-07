package main

import (
	"bufio"
	"fmt"
	"os"
)

func Prompt(message string) {
	fmt.Print("PROMPT ", message, " > ")
}

func humanChoose(cards Cards) (selected, rejected Card) {
	var stdin = bufio.NewReader(os.Stdin)
	var choice string
	for _, c := range cards {
		Show(c.Id(), c.Name())
	}
	for {
		Prompt("Choose a card")
		input, _ := fmt.Fscanln(stdin, &choice)
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

func Wait() {
	if wait {
		Show("Press Enter to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
